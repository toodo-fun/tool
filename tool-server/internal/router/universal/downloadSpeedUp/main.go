package downloadSpeedUp

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/universal/downloadSpeedUp/service"
)

func handleGetDownloadSpeedUpInfo(c *gin.Context) {
	downloadURL := c.Query("url")
	logrus.Infof("downloadURL: %+v", downloadURL)
	if downloadURL == "" {
		res := service.DownloadSpeedUpInfo{Status: "Failed"}
		c.JSON(http.StatusOK, response.SuccessResponse(res))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(service.GetDownloadSpeedUpInfo(downloadURL)))
}

func InitRouter() {
	r := router.RegisterRouterGroup("universal/downloadSpeedUp")
	r.GET("", handleGetDownloadSpeedUpInfo)
}
