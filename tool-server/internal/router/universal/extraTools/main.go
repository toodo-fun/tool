package extraTools

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	service3 "tool-server/internal/domain/universal/downloadSpeedUp/service"
	service2 "tool-server/internal/domain/universal/downloader/service"
	"tool-server/internal/domain/universal/extraTool/entity"
	"tool-server/internal/domain/universal/extraTool/service"
	"tool-server/internal/utils/client/downloader"
	"tool-server/internal/utils/executor"
	"tool-server/internal/utils/path"
	"tool-server/internal/utils/zip"
)

func handleGetExtraTools(c *gin.Context) {
	extraTools, err := service.ExtraToolService{}.List(entity.ExtraTool{})
	if err != nil {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeSystemError))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(extraTools))
}

func handleInstallExtraTool(c *gin.Context) {
	extraTool := entity.ExtraTool{}
	err := c.Bind(&extraTool)
	if err != nil {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeParamError))
		return
	}

	extraToolPath := fmt.Sprintf("%s/extraTools", path.GetPlatformRoot())
	if !path.IsExist(extraToolPath) {
		os.MkdirAll(extraToolPath, 0655)
	}

	extraTool.Status = entity.STATUS_INSTALLING
	service.ExtraToolService{}.Update(extraTool)

	speedUpInfo := service3.GetDownloadSpeedUpInfo(extraTool.DownloadUrl)
	if speedUpInfo.Status != "Failed" {
		extraTool.DownloadUrl = speedUpInfo.DlURL
		logrus.Infof("加速下载解析成功，使用加速的地址进行下载: %+v", extraTool.DownloadUrl)
	}

	instanceId := service2.CreateDownloadTask(extraTool.DownloadUrl, extraToolPath)
	c.JSON(http.StatusOK, response.DefaultSuccessResponse())
	go func() {
		for {
			downloadInstance, _ := service2.GetDownloadTaskStatus(instanceId)
			// 更新表
			downloadInfo := map[string]any{
				"total":      downloadInstance.SizeTotle,
				"downloaded": downloadInstance.SizeDownloaded,
				"filename":   downloadInstance.Filename,
			}
			downloadInfoJson, _ := json.Marshal(downloadInfo)
			extraTool.DownloadInfo = string(downloadInfoJson)

			// 终止条件
			if downloadInstance.Status == downloader.STATUS_FAILED {
				extraTool.Status = entity.STATUS_ERROR
				break
			}
			if downloadInstance.Status == downloader.STATUS_SUCCESS {
				downloadInfo["downloaded"] = downloadInstance.SizeTotle
				service.ExtraToolService{}.Update(extraTool)
				// 开始解压
				logrus.Infof("开始解压")
				err = zip.Unzip(extraToolPath+string(filepath.Separator)+downloadInstance.Filename, extraToolPath+string(filepath.Separator)+extraTool.Title)
				if err != nil {
					logrus.Errorf("解压失败: %+v", err)
					extraTool.Status = entity.STATUS_ERROR
				} else {
					logrus.Infof("解压完成")
					extraTool.Status = entity.STATUS_INSTALLED
				}
				os.Remove(extraToolPath + string(filepath.Separator) + downloadInstance.Filename)
				service.ExtraToolService{}.Update(extraTool)
				break
			}
			service.ExtraToolService{}.Update(extraTool)
			time.Sleep(1 * time.Second)
		}
	}()
}

func handleRunExtraTool(c *gin.Context) {
	cmd := c.Query("cmd")
	getExecutor := executor.GetExecutor(false)
	getExecutor.RunCmd([]string{fmt.Sprintf("%s%s", path.GetPlatformRoot(), cmd)})
}

func InitRouter() {
	r := router.RegisterRouterGroup("extraTool")
	r.GET("list", handleGetExtraTools)
	r.POST("install", handleInstallExtraTool)
	r.GET("start", handleRunExtraTool)
}
