package contextMenu

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/windows/contextMenu/service"
)

func handleFold(c *gin.Context) {
	success := service.Fold()
	if success {
		c.JSON(http.StatusOK, response.DefaultSuccessResponse())
	} else {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeSystemError))
	}
}

func handleUnfold(c *gin.Context) {
	success := service.Unfold()
	if success {
		c.JSON(http.StatusOK, response.DefaultSuccessResponse())
	} else {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeSystemError))
	}
}

func InitRouter() {
	r := router.RegisterRouterGroup("windows/contextMenu")
	r.GET("fold", handleFold)
	r.GET("unfold", handleUnfold)
}
