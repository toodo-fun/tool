package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/sirupsen/logrus"
	"tool-server/internal/global"
	"tool-server/internal/router"
	"tool-server/internal/utils/client/db"
	"tool-server/internal/utils/path"
)

func init() {
	devMode := false

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)

	var configFiles []string
	if devMode {
		configFiles = []string{
			"config/application.yml",
			"config/changelog.yml",
			"config/navigation.yml",
		}
	} else {
		configFiles = []string{
			fmt.Sprintf("%s/config/application.yml", path.GetPlatformRoot()),
			fmt.Sprintf("%s/config/changelog.yml", path.GetPlatformRoot()),
			fmt.Sprintf("%s/config/navigation.yml", path.GetPlatformRoot()),
		}
	}

	err := config.LoadFiles(configFiles...)
	if err != nil {
		panic(err)
	}

	changelog := global.Release{}
	config.BindStruct("release", &changelog)
	global.ChangeLog = changelog

	navigation := global.Navigation{}
	config.BindStruct("navigation", &navigation)
	global.NavigationMenu = navigation

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

	global.Prefix = config.String("server.prefix")

	dBClient := db.NewDbClient(config.String("database.source"), config.Int("database.maxOpenConn"), config.Int("database.maxIdleConn"), config.Int("database.connMaxLifeTime"))
	global.DBClient = dBClient
	global.DBMaxThread = make(chan struct{}, config.Int("database.maxOpenConn"))
}

func main() {
	r := router.CreateRoute()
	logrus.Infof("service run on http://127.0.0.1:%d%s", config.Int("server.port"), config.String("server.prefix"))
	r.Run(fmt.Sprintf(":%d", config.Int("server.port")))
}
