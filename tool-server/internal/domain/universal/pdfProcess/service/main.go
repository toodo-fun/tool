package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
	"os"
	"strings"
	"tool-server/internal/utils/path"
)

type PDFInfo struct {
	Page     int    `json:"page"`
	Filename string `json:"filename"`
	MD5      string `json:"md5"`
	Status   string `json:"status"`
}

func GetPDFInfo(inFile string) (result PDFInfo, err error) {
	page, err := api.PageCountFile(inFile)
	pFile, err := os.Open(inFile)
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	return PDFInfo{
		Page:     page,
		Filename: inFile,
		MD5:      hex.EncodeToString(md5h.Sum(nil)),
	}, err

}

func SplitPDF(inFile string, outDir string, span int) (err error) {
	return api.SplitFile(inFile, outDir, span, nil)
}

func MergePDF(inFile []string, outFile string) (oFile string, err error) {
	if path.IsExist(outFile) {
		outFile = strings.Replace(outFile, "result.pdf", fmt.Sprintf("%s.pdf", uuid.New().String()), -1)
	}
	return outFile, api.MergeCreateFile(inFile, outFile, nil)
}
