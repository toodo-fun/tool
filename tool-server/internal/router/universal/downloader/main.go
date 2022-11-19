package downloader

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/universal/downloader/service"
)

func handleCreateDownloadTask(c *gin.Context) {
	type DownloadParams struct {
		Url         string `json:"url"`
		DownloadDir string `json:"downloadDir"`
	}
	params := DownloadParams{}
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeParamError))
		return
	}
	instanceId := service.CreateDownloadTask(params.Url, params.DownloadDir)
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{
		"instanceId": instanceId,
	}))
}

func handleGetDownloadTaskStatus(c *gin.Context) {
	instanceId := c.Query("instanceId")
	if instanceId == "" {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeParamError))
		return
	}
	if instance, ok := service.GetDownloadTaskStatus(instanceId); ok {
		c.JSON(http.StatusOK, response.SuccessResponse(instance))
	} else {
		c.JSON(http.StatusOK, response.ErrorResponse(response.CodeParamError, "流程不存在"))
	}
}

func handleGetDownloadTaskList(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(service.GetDownloadTaskList()))
}

func handleGetRunningDownloadTaskCount(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]int{
		"count": service.GetRunningDownloadTaskCount(),
	}))
}

func InitRouter() {
	r := router.RegisterRouterGroup("downloader")
	r.POST("create", handleCreateDownloadTask)
	r.GET("status", handleGetDownloadTaskStatus)
	r.GET("tasks", handleGetDownloadTaskList)
	r.GET("count", handleGetRunningDownloadTaskCount)
}
