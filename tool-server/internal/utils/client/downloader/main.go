package downloader

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	_map "tool-server/internal/utils/common/map"
	path2 "tool-server/internal/utils/path"
)

var (
	STATUS_SUCCESS = "success"
	STATUS_FAILED  = "failed"
	STATUS_RUNNING = "running"
	STATUS_MERGE   = "merge"

	MODEL_MULTIPLE = "multiple"
	MODEL_SINGLE   = "single"
)

type Downloader struct {
	URL            string `json:"url"`
	headResponse   *http.Response
	Status         string `json:"status"`
	Filename       string `json:"filename"`
	SizeTotle      int64  `json:"sizeTotle"`
	SizeDownloaded int64  `json:"sizeDownloaded"`
	DownloadModel  string `json:"downloadModel"`
	Thread         int    `json:"thread"`
	DownloadDir    string `json:"downloadDir"`
	lock           *sync.RWMutex
}

func CreateDownloader(url string, downloaderDir string) *Downloader {
	return &Downloader{
		URL:           url,
		Thread:        8,
		DownloadModel: MODEL_MULTIPLE,
		DownloadDir:   downloaderDir,
		lock:          new(sync.RWMutex),
	}
}

func (d *Downloader) getHeadResponse() {
	res, err := http.Head(d.URL)
	if err != nil {
		logrus.Errorf("Get Head response failed: %+v", err)
		d.Status = STATUS_FAILED
		return
	}
	logrus.Infof("%+v", res.Header)
	d.headResponse = res
}

func (d *Downloader) checkUrlValid() {
	if d.headResponse == nil {
		d.Status = STATUS_FAILED
		return
	}
	if d.headResponse.StatusCode != http.StatusOK {
		d.Status = STATUS_FAILED
		logrus.Errorf("check url vaild unpass")
	}
}

func (d *Downloader) checkSupportMultiple() {
	if d.headResponse == nil {
		d.Status = STATUS_FAILED
		return
	}
	if d.headResponse.Header.Get("Accept-Ranges") != "bytes" {
		d.Status = STATUS_FAILED
		logrus.Errorf("check support multiple unpass")
	}
}

func (d *Downloader) getFileInfo() {
	if d.headResponse == nil {
		d.Status = STATUS_FAILED
		return
	}
	size, err := strconv.ParseInt(d.headResponse.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		logrus.Warnf("get file size failed, use single model: %+v", err)
		d.DownloadModel = MODEL_SINGLE
	}
	d.SizeTotle = size
	d.Filename = path.Base(strings.Split(d.URL, "?")[0])

	// 如果文件小于1MB，则使用单线程下载
	if size/int64(d.Thread) < 1024*1024 {
		d.DownloadModel = MODEL_SINGLE
	}

	logrus.Infof("文件名称: %s, 文件大小：%d, 下载模式: %s", d.Filename, size, d.DownloadModel)
}

func (d *Downloader) download() {
	file := d.DownloadDir + string(filepath.Separator) + d.Filename
	c := 1
	for {
		if path2.IsExist(file) {
			file = d.DownloadDir + string(filepath.Separator) + strconv.Itoa(c) + "-" + d.Filename
			c += 1
		} else {
			if c != 1 {
				d.Filename = strconv.Itoa(c) + "-" + d.Filename
			}
			break
		}
	}
	if d.DownloadModel == MODEL_SINGLE {
		tmpFilename := uuid.New().String() + ".download"
		d.downloader(0, d.SizeTotle, tmpFilename)
		if d.Status != STATUS_FAILED {
			d.Status = STATUS_MERGE
			os.Rename(d.DownloadDir+string(filepath.Separator)+tmpFilename, file)
		}
	} else {
		// 文件分块
		bufferSize := int64(math.Ceil(float64(d.SizeTotle / int64(d.Thread))))
		var bufferList []int64
		for i := 0; i <= d.Thread; i++ {
			tmpSize := int64(i) * bufferSize
			if i == d.Thread {
				tmpSize = d.SizeTotle
			}
			bufferList = append(bufferList, tmpSize)
		}
		bufferFileList := _map.OrderedMap[[]int64]{}
		for index, _ := range bufferList {
			if index != len(bufferList)-1 {
				var start int64 = 0
				if index == 0 {
					start = 0
				} else {
					start = bufferList[index] + 1
				}
				bufferFileList.Set(uuid.New().String()+".download", []int64{start, bufferList[index+1]})
			}
		}

		// 开始分块下载
		wg := sync.WaitGroup{}
		for _, item := range bufferFileList.Keys {
			rangeList := bufferFileList.Get(item)
			maxTheard := make(chan struct{}, d.Thread)
			maxTheard <- struct{}{}
			wg.Add(1)
			item := item
			go func() {
				defer wg.Done()
				d.downloader(rangeList[0], rangeList[1], item)
				<-maxTheard
			}()
		}
		wg.Wait()

		// 开始合并
		logrus.Infof("开始合并文件")
		if d.Status != STATUS_FAILED {
			d.Status = STATUS_MERGE
			finalFile, err := os.Create(file)
			if err != nil {
				d.Status = STATUS_FAILED
				return
			}
			defer finalFile.Close()
			for _, item := range bufferFileList.Keys {
				source, err := os.Open(d.DownloadDir + string(filepath.Separator) + item)
				if err != nil {
					d.Status = STATUS_FAILED
					return
				}
				io.Copy(finalFile, source)
				source.Close()
				os.Remove(d.DownloadDir + string(filepath.Separator) + item)
			}
		}
	}
}

func (d *Downloader) downloader(rangeStart int64, rangeEnd int64, filename string) {
	buf := make([]byte, 32*1024)
	req, _ := http.NewRequest("GET", d.URL, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("download get error: %+v", err)
		d.Status = STATUS_FAILED
		return
	}
	tmpFile, err := os.Create(d.DownloadDir + string(filepath.Separator) + filename)
	if err != nil {
		logrus.Errorf("download create tmp file error: %+v", err)
		d.Status = STATUS_FAILED
		return
	}
	defer tmpFile.Close()
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, ew := tmpFile.Write(buf[0:nr])
			d.lock.Lock()
			d.SizeDownloaded += int64(nw)
			d.lock.Unlock()
			//logrus.Infof("进度：%d/%d", d.SizeDownloaded, d.SizeTotle)
			if ew != nil {
				logrus.Errorf("buff error : %+v, 即将重试", ew)
				tmpFile.Close()
				os.Remove(d.DownloadDir + string(filepath.Separator) + filename)
				d.downloader(rangeStart, rangeEnd, filename)
				break
			}
			if nr != nw {
				logrus.Errorf("buff error: %+v, 即将重试", io.ErrShortBuffer)
				tmpFile.Close()
				os.Remove(d.DownloadDir + string(filepath.Separator) + filename)
				d.downloader(rangeStart, rangeEnd, filename)
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				if er != nil {
					logrus.Errorf("buff error: %+v, 即将重试", er)
					tmpFile.Close()
					os.Remove(d.DownloadDir + string(filepath.Separator) + filename)
					d.downloader(rangeStart, rangeEnd, filename)
					break
				}
			}
			break
		}
	}
}

func (d *Downloader) Start() {
	d.Status = STATUS_RUNNING

	runStages := _map.OrderedMap[func()]{}
	runStages.Set("getHeadResponse", d.getHeadResponse)
	runStages.Set("checkUrlValid", d.checkUrlValid)
	runStages.Set("checkSupportMultiple", d.checkSupportMultiple)
	runStages.Set("getFileInfo", d.getFileInfo)
	runStages.Set("download", d.download)

	for _, name := range runStages.Keys {
		logrus.Infof("开始执行 %s", name)
		runner := runStages.Get(name)
		runner()
		if d.Status == STATUS_FAILED {
			logrus.Errorf("%s 执行失败", name)
			break
		}
		logrus.Infof("%s 执行成功", name)
	}
	if d.Status != STATUS_FAILED {
		d.Status = STATUS_SUCCESS
	}
}
