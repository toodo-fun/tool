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
	Instance = map[string]Executor{}
)

type Executor interface {
	RunCmd(command []string) (instanceId string, success bool)
	GetLog() []string
	AddLog(string)
}

func GetExecutor(sync bool) Executor {
	if sync {
		return &SyncExecutor{}
	} else {
		return &AsyncExecutor{}
	}
}

func RunCommand(instanceId string, executor Executor, command ...string) (success bool) {
	Instance[instanceId] = executor
	// 支持中文编码
	enc := mahonia.NewDecoder("gbk")

	// 判断运行系统
	platform := runtime.GOOS
	var cmd *exec.Cmd
	if platform == "windows" {
		command = append([]string{"/c"}, command...)
		cmd = exec.Command("cmd", command...)
	} else {
		command = append([]string{"-c"}, command...)
		cmd = exec.Command("bash", command...)
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
			executor.AddLog(enc.ConvertString(strings.Replace(readString, "\n", "", -1)))
		}
		readerStdErr := bufio.NewReader(stderr)
		for {
			readString, err := readerStdErr.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			executor.AddLog(enc.ConvertString(strings.Replace(readString, "\n", "", -1)))
		}
	}()

	err := cmd.Start()
	executor.AddLog(fmt.Sprintf("PID: %d; command: %s", cmd.Process.Pid, command))

	wg.Wait()
	err = cmd.Wait()

	if err != nil {
		return false
	}

	return true
}
