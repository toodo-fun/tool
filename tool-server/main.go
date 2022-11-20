package main

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/sirupsen/logrus"
	"tool-server/internal/core/platform/update"
	"tool-server/internal/router"
)

func init() {
	releaseInfo := update.CheckReleaseUpdate()
	logrus.Infof("%+v", releaseInfo)
}

func main() {
	r := router.CreateRoute()
	logrus.Infof("service run on http://127.0.0.1:%d%s", config.Int("server.port"), config.String("server.prefix"))
	r.Run(fmt.Sprintf(":%d", config.Int("server.port")))
}
