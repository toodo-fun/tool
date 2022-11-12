package executor

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var (
	Log = make([]string, 0)
)

type SyncExecutor struct {
}

func (s SyncExecutor) RunCmd(command string) bool {
	// 支持中文编码
	enc := mahonia.NewDecoder("gbk")

	// 判断运行系统
	platform := runtime.GOOS
	cmd := exec.Command("cmd", "/c", command)
	if platform != "windows" {
		cmd = exec.Command("bash", "-c", command)
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// 日志处理
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		readerStdOut := bufio.NewReader(stdout)
		for {
			readString, err := readerStdOut.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			Log = append(Log, enc.ConvertString(strings.Replace(readString, "\n", "", -1)))
		}
		readerStdErr := bufio.NewReader(stderr)
		for {
			readString, err := readerStdErr.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			Log = append(Log, enc.ConvertString(strings.Replace(readString, "\n", "", -1)))
		}
	}()

	err := cmd.Start()
	Log = append(Log, fmt.Sprintf("PID: %d; command: %s", cmd.Process.Pid, command))

	wg.Wait()
	err = cmd.Wait()

	if err != nil {
		return false
	}

	return true
}

func (s SyncExecutor) GetLog() []string {
	return Log
}
