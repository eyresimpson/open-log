package main

import (
	"open-log/service"
	"open-log/tools"
	"open-log/tools/comm"
	"os"
	"strings"
)

// 主入口函数
func main() {
	if len(os.Args) < 2 {
		tools.Err("alt-open-log: Missing parameters/Incorrect number of parameters passed", nil)
		service.GetHelp()
		return
	}
	// 系统初始化
	comm.InitRootWorkSpace("open-log")

	if !strings.HasPrefix(os.Args[1], "-") {
		service.Run("", os.Args[1:])
		return
	}
	switch os.Args[1] {
	case "-v":
		service.GetVersion()
	case "-h":
		service.GetHelp()
	default:
		service.GetHelp()
	}

}
