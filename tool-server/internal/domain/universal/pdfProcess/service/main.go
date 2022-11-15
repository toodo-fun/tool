package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
	"os"
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

func MergePDF(inFile []string, outFile string) (err error) {
	return api.MergeCreateFile(inFile, outFile, nil)
}
