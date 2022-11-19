package service

import (
	"github.com/google/uuid"
	"tool-server/internal/utils/client/downloader"
)

var (
	instances = make(map[string]*downloader.Downloader, 0)
)

func CreateDownloadTask(url string, downloadDir string) (instanceId string) {
	instanceId = uuid.New().String()
	downloadInstance := downloader.CreateDownloader(url, downloadDir)
	instances[instanceId] = downloadInstance
	go downloadInstance.Start()
	return instanceId
}

func GetDownloadTaskStatus(instanceId string) (instance *downloader.Downloader, ok bool) {
	if instance, ok = instances[instanceId]; ok {
		return instance, ok
	} else {
		return &downloader.Downloader{}, ok
	}
}

func GetDownloadTaskList() map[string]*downloader.Downloader {
	return instances
}

func GetRunningDownloadTaskCount() (count int) {
	count = 0
	for _, v := range instances {
		if v.Status != downloader.STATUS_FAILED && v.Status != downloader.STATUS_SUCCESS {
			count += 1
		}
	}
	return count
}
