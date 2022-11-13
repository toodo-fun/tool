package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tool-server/internal/global"
)

var (
	engine *gin.Engine
)

func RegisterEngine(e *gin.Engine) {
	engine = e
}

func GetEngine() *gin.Engine {
	return engine
}

func RegisterRouterGroup(group string) *gin.RouterGroup {
	routerGroup := fmt.Sprintf("%s/%s", global.Prefix, group)
	logrus.Infof("路由组注册成功：%s", routerGroup)
	return engine.Group(routerGroup)
}
