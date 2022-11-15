package service

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"tool-server/internal/utils/client/http"
)

type DownloadSpeedUpInfo struct {
	DlURL    string  `json:"dlurl"`
	Filename string  `json:"filename"`
	MD5      string  `json:"md5"`
	Size     float32 `json:"size"`
	Sizeunit string  `json:"sizeunit"`
	Status   string  `json:"status"`
}

func GetDownloadSpeedUpInfo(downloadURL string) (result DownloadSpeedUpInfo) {
	baseURL := "https://proxy.chromiumer.com/dl?url="
	logrus.Info(baseURL)
	res, err := http.Get(baseURL+downloadURL, nil)
	if err != nil {
		logrus.Errorf("下载加速解析失败: %+v", err)
		result.Status = "Failed"
	} else {
		err := json.Unmarshal(res, &result)
		if err != nil {
			logrus.Errorf("下载加速解析失败: %+v", err)
			result.Status = "Failed"
		}
	}
	return result
}
