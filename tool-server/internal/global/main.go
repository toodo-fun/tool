package global

import (
	"gorm.io/gorm"
)

var (
	DBClient       *gorm.DB
	DBMaxThread    chan struct{}
	Prefix         string
	ChangeLog      Release
	NavigationMenu Navigation
)

type Release []struct {
	Version   string   `yaml:"version" json:"version"`
	Changelog []string `yaml:"changelog" json:"changelog"`
}

type Navigation []struct {
	Title    string     `yaml:"title" json:"title"`
	Id       string     `yaml:"id" json:"id"`
	Describe string     `yaml:"describe" json:"describe"`
	Icon     string     `yaml:"icon" json:"icon"`
	Href     string     `yaml:"href" json:"href"`
	Children Navigation `yaml:"children" json:"children"`
}
