package service

import (
	"github.com/sirupsen/logrus"
	"time"
	"tool-server/internal/domain/universal/extraTool/entity"
	"tool-server/internal/domain/universal/extraTool/mapper"
)

type ExtraToolService struct {
	mapper.ExtraToolMapper
}

func init() {
	extraTool := entity.ExtraTool{
		ID:           "defaultTool",
		Title:        "TrafficMonitor",
		Describe:     "Traffic Monitor是一款用于Windows平台的网速监控悬浮窗软件，可以显示当前网速、CPU及内存利用率，支持嵌入到任务栏显示，支持更换皮肤、历史流量统计等功能。",
		Icon:         "https://pngimg.com/uploads/github/github_PNG40.png",
		Status:       "uninstall",
		DownloadUrl:  "https://github.com/zhongyang219/TrafficMonitor/releases/download/V1.84.1/TrafficMonitor_V1.84.1_x64.zip",
		DownloadInfo: "",
		StartCmd:     "/extraTools/TrafficMonitor/TrafficMonitor.exe",
		CreateAt:     time.Time{},
		UpdateAt:     time.Time{},
	}
	err := ExtraToolService{}.Insert(extraTool)
	if err != nil {
		logrus.Infof("默认工具已存在")
	}

	// 状态归位
	extraTools, err := ExtraToolService{}.List(entity.ExtraTool{})
	for _, item := range extraTools {
		if item.Status == entity.STATUS_INSTALLING {
			item.Status = entity.STATUS_UNINSTALL
			ExtraToolService{}.Update(item)
		}
	}

}
