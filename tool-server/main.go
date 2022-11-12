package main

import (
	"fmt"
	"strings"
	"tool-server/internal/utils/executor"
)

func main() {
	a := executor.GetExecutor(true)
	b := a.RunCmd("i")
	c := a.GetLog()
	fmt.Printf("b: %+v", b)
	fmt.Printf("c: %+v", strings.Join(c, "\n"))
}
