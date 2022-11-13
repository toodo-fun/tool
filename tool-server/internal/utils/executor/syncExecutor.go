package executor

import (
	"github.com/google/uuid"
)

type SyncExecutor struct {
	Log []string
}

func (s *SyncExecutor) RunCmd(command []string) (instanceId string, success bool) {
	instanceId = uuid.New().String()
	s.Log = make([]string, 0)
	success = RunCommand(instanceId, s, command...)
	return instanceId, success
}

func (s *SyncExecutor) GetLog() []string {
	return s.Log
}

func (s *SyncExecutor) AddLog(log string) {
	s.Log = append(s.Log, log)
}
