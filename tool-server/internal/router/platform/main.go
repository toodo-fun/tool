package platform

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tool-server/internal/core/platform/update"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/platform/service"
)

func handleChangelog(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(service.GetChangelog()))
}

func handleNavigation(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(service.GetNavigation()))
}

func handleCheckReleaseUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(update.CheckReleaseUpdate()))
}

func handleUpdatePlatform(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeParamError))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(update.Update(url)))
}

func InitRouter() {
	r := router.RegisterRouterGroup("platform")
	r.GET("changelog", handleChangelog)
	r.GET("navigation", handleNavigation)
	r.GET("checkUpdate", handleCheckReleaseUpdate)
	r.GET("update", handleUpdatePlatform)
}
