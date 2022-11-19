package entity

import "time"

var (
	STATUS_UNINSTALL  = "uninstall"
	STATUS_INSTALLING = "installing"
	STATUS_INSTALLED  = "installed"
	STATUS_ERROR      = "error"
)

type ExtraTool struct {
	ID           string    `gorm:"size:36;primaryKey" json:"id"`
	Title        string    `gorm:"size:36;not null" json:"title"`
	Describe     string    `gorm:"size:1024" json:"describe"`
	Icon         string    `gorm:"size:1024;not null" json:"icon"`
	Status       string    `gorm:"size:36;not null" json:"status"`
	DownloadUrl  string    `gorm:"size:2048;not null" json:"downloadUrl"`
	DownloadInfo string    `gorm:"size:2048" json:"downloadInfo"`
	StartCmd     string    `gorm:"size:256;not null" json:"startCmd"`
	CreateAt     time.Time `gorm:"not null;autoCreateTime;type:timestamp" json:"createAt"`
	UpdateAt     time.Time `gorm:"not null;autoUpdateTime;type:timestamp" json:"updateAt"`
}
