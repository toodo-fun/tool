package executor

import (
	"github.com/google/uuid"
)

type AsyncExecutor struct {
	Log []string
}

func (a *AsyncExecutor) RunCmd(command []string) (instanceId string, success bool) {
	instanceId = uuid.New().String()
	a.Log = make([]string, 0)
	go RunCommand(instanceId, a, command...)
	return instanceId, true
}

func (a *AsyncExecutor) GetLog() []string {
	return a.Log
}

func (a *AsyncExecutor) AddLog(log string) {
	a.Log = append(a.Log, log)
}
