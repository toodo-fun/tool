package global

import "gorm.io/gorm"

var (
	DBClient    *gorm.DB
	DBMaxThread chan struct{}
	Prefix      string
)
