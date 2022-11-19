package global

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"tool-server/internal/utils/client/db"
	"tool-server/internal/utils/path"
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

func init() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)

	configFiles := []string{
		fmt.Sprintf("%s/config/application.yml", path.GetPlatformRoot()),
		fmt.Sprintf("%s/config/changelog.yml", path.GetPlatformRoot()),
		fmt.Sprintf("%s/config/navigation.yml", path.GetPlatformRoot()),
	}

	logrus.Infof("configFiles: %+v", configFiles)
	err := config.LoadFiles(configFiles...)
	if err != nil {
		panic(err)
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:       false,
		NoColors:       false,
		NoFieldsColors: false,
		ShowFullLevel:  false,
		TrimMessages:   false,
		CallerFirst:    true,
		FieldsOrder:    []string{"component", "category"},
	})
	logrus.SetLevel(logrus.Level(config.Int("server.log.level")))

	ChangeLog = Release{}
	config.BindStruct("release", &ChangeLog)

	NavigationMenu = Navigation{}
	config.BindStruct("navigation", &NavigationMenu)

	Prefix = config.String("server.prefix")

	DBClient = db.NewDbClient(config.String("database.source"), config.Int("database.maxOpenConn"), config.Int("database.maxIdleConn"), config.Int("database.connMaxLifeTime"))

	logrus.Info("数据库客户端创建成功")
	DBMaxThread = make(chan struct{}, config.Int("database.maxOpenConn"))
}
