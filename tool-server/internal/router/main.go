package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/router/windows/contextMenu"
	"tool-server/internal/router/windows/platform"
)

func handleNotFound(c *gin.Context) {
	logrus.Error(fmt.Sprintf("404 not found, %s", c.Request.URL.Path))
	c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeNotFound))
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, response.DefaultSuccessResponse())
}

func CreateRoute() *gin.Engine {
	r := gin.Default()

	//注册404路由
	r.NoRoute(handleNotFound)
	r.NoMethod(handleNotFound)

	router.RegisterEngine(r)

	contextMenu.InitRouter()
	platform.InitRouter()

	InitHealthRouter()

	return r
}

func InitHealthRouter() {
	r := router.RegisterRouterGroup("default")
	r.GET("health", handleHealthCheck)
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}
