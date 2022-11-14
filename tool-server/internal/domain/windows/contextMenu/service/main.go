package service

import (
	"github.com/sirupsen/logrus"
	"strings"
	"time"
	"tool-server/internal/utils/executor"
)

type ContextMenuService struct {
}

func GetFoldStatus() (fold bool) {
	cmd := []string{"reg", "query", `HKCU\Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}`}
	e := executor.GetExecutor(true)
	_, fold = e.RunCmd(cmd)
	return fold
}

func Fold() (success bool) {
	cmd := []string{"reg", "delete", `HKCU\Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}`, "/f"}
	e := executor.GetExecutor(true)
	_, success = e.RunCmd(cmd)
	logrus.Debug(e.GetLog())
	if success {
		restartExplorer()
	}
	return success
}

func Unfold() (success bool) {
	cmd := []string{"reg", "add", `HKCU\Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}\InprocServer32`, "/f", "/ve"}
	e := executor.GetExecutor(true)
	_, success = e.RunCmd(cmd)
	logrus.Debug(strings.Join(e.GetLog(), "\n"))
	if success {
		restartExplorer()
	}
	return success
}

func restartExplorer() {
	var cmd []string
	var e executor.Executor
	cmd = []string{"taskkill", "/f", "/im", "explorer.exe"}
	e = executor.GetExecutor(true)
	e.RunCmd(cmd)
	logrus.Debug(strings.Join(e.GetLog(), "\n"))

	cmd = []string{"explorer.exe"}
	e = executor.GetExecutor(false)
	e.RunCmd(cmd)
	// todo 启动资源管理器的命令执行后不会结束，所以暂时硬写3秒后退出方法
	time.Sleep(3 * time.Second)
	logrus.Debug(strings.Join(e.GetLog(), "\n"))
}
