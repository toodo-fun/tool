package update

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"tool-server/internal/domain/universal/downloadSpeedUp/service"
	service2 "tool-server/internal/domain/universal/downloader/service"
	"tool-server/internal/global"
	"tool-server/internal/utils/client/downloader"
	"tool-server/internal/utils/client/http"
	"tool-server/internal/utils/path"
)

const (
	releaseGithub = "https://api.github.com/repos/toodo-fun/tool/releases/latest"
)

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
		Url  string `json:"browser_download_url"`
	} `json:"assets"`
	ReleaseDate time.Time `json:"published_at"`
	Body        string    `json:"body"`
	Update      bool      `json:"update"`
}

func CheckReleaseUpdate() (releaseInfo ReleaseInfo) {
	releaseInfo.Update = false
	res, err := http.Get(releaseGithub, nil)
	if err != nil {
		logrus.Errorf("查询更新信息失败: %+v", err)
		return releaseInfo
	}
	err = json.Unmarshal(res, &releaseInfo)
	if err != nil {
		logrus.Errorf("更新信息解析失败: %+v", err)
		return releaseInfo
	}
	logrus.Info(releaseInfo)
	logrus.Infof("current: %s", global.ChangeLog[0].Version)
	if releaseInfo.TagName != global.ChangeLog[0].Version {
		releaseInfo.Update = true
		return releaseInfo
	}
	return releaseInfo
}

func Update(url string) (filename string) {
	speedUpInfo := service.GetDownloadSpeedUpInfo(url)
	if speedUpInfo.Status != "Failed" {
		url = speedUpInfo.DlURL
		logrus.Infof("加速下载解析成功，使用加速的地址进行下载: %+v", url)
	}
	instanceId := service2.CreateDownloadTask(url, fmt.Sprintf("%s", path.GetPlatformRoot()))
	for {
		downloadInstance, _ := service2.GetDownloadTaskStatus(instanceId)
		if downloadInstance.Status == downloader.STATUS_FAILED {
			return ""
		}
		if downloadInstance.Status == downloader.STATUS_SUCCESS {
			return downloadInstance.Filename
		}
	}
}
