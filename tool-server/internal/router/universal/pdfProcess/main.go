package pdfProcess

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tool-server/internal/core/response"
	"tool-server/internal/core/router"
	"tool-server/internal/domain/universal/pdfProcess/service"
)

func handleGetPDFInfo(c *gin.Context) {
	filePath := c.Query("filepath")
	pdfInfo, err := service.GetPDFInfo(filePath)
	if err == nil {
		pdfInfo.Status = "success"
	}
	c.JSON(http.StatusOK, response.SuccessResponse(pdfInfo))
}

func handleMergePDF(c *gin.Context) {
	type MergePDF struct {
		InFiles []string `json:"inFiles"`
		OutFile string   `json:"outFile"`
	}
	params := MergePDF{}
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeParamError))
		return
	}
	err = service.MergePDF(params.InFiles, params.OutFile)
	if err != nil {
		c.JSON(http.StatusOK, response.DefaultErrorResponse(response.CodeSystemError))
		return
	}
	c.JSON(http.StatusOK, response.DefaultSuccessResponse())
}

func InitRouter() {
	r := router.RegisterRouterGroup("pdf")
	r.GET("info", handleGetPDFInfo)
	r.POST("merge", handleMergePDF)
}
