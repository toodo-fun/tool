package mapper

import (
	"github.com/sirupsen/logrus"
	"tool-server/internal/core/mapper"
	"tool-server/internal/domain/universal/extraTool/entity"
	"tool-server/internal/global"
)

type ExtraToolMapper struct {
	mapper.BaseMapper[entity.ExtraTool]
}

func init() {
	err := global.DBClient.AutoMigrate(&entity.ExtraTool{})
	if err != nil {
		logrus.Errorf("DB ExtraTools表迁移失败")
	}
}
