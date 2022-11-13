package platform

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/platform/service"
)

func handleChangelog(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(service.GetChangelog()))
}

func InitRouter() {
	r := router.RegisterRouterGroup("platform")
	r.GET("changelog", handleChangelog)
}
