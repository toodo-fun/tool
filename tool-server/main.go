package main

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/sirupsen/logrus"
	"tool-server/internal/router"
)

func init() {
	// test 代码
	//d := downloader.CreateDownloader("https://download.jetbrains.com/idea/ideaIU-2022.2.3.exe?test=123", "C:\\Users\\marui\\Downloads")
	//d.Start()
	//logrus.Infof("downloader: %+v", d)
}

func main() {
	r := router.CreateRoute()
	logrus.Infof("service run on http://127.0.0.1:%d%s", config.Int("server.port"), config.String("server.prefix"))
	r.Run(fmt.Sprintf(":%d", config.Int("server.port")))
}
