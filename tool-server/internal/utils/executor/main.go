package executor

type Executor interface {
	RunCmd(command string) bool
	GetLog() []string
}

func GetExecutor(sync bool) Executor {
	if sync {
		return SyncExecutor{}
	} else {
		return SyncExecutor{}
	}
}
