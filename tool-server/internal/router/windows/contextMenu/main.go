package contextMenu

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/windows/contextMenu/service"
)

func handleFold(c *gin.Context) {
	res := make(map[string]any, 0)
	res["success"] = false
	if !checkSystem() {
		res["message"] = "仅支持Windows11系统使用该功能"
		c.JSON(http.StatusOK, response.SuccessResponse(res))
		return
	}
	res["success"] = service.Fold()
	c.JSON(http.StatusOK, response.SuccessResponse(res))
}

func handleUnfold(c *gin.Context) {
	res := make(map[string]any, 0)
	res["success"] = false
	if !checkSystem() {
		res["message"] = "仅支持Windows11系统使用该功能"
		c.JSON(http.StatusOK, response.SuccessResponse(res))
		return
	}
	res["success"] = service.Unfold()
	c.JSON(http.StatusOK, response.SuccessResponse(res))
}

func handleGetFoldStatus(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(service.GetFoldStatus()))
}

func checkSystem() (access bool) {
	if runtime.GOOS != "windows" {
		return false
	}
	return true
}

func InitRouter() {
	r := router.RegisterRouterGroup("windows/contextMenu")
	r.GET("fold", handleFold)
	r.GET("unfold", handleUnfold)
	r.GET("status", handleGetFoldStatus)
}
