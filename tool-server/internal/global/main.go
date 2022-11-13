package global

import (
	"gorm.io/gorm"
)

var (
	DBClient    *gorm.DB
	DBMaxThread chan struct{}
	Prefix      string
	ChangeLog   Release
)

type Release []struct {
	Version   string   `yaml:"version" json:"version"`
	Changelog []string `yaml:"changelog" json:"changelog"`
}
