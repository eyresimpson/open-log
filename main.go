package main

import (
	"alog/service"
	"alog/tools/comm"
	"alog/tools/logger"
	"os"
	"strings"
)

// 主入口函数
func main() {
	if len(os.Args) < 2 {
		logger.Err("alt-alog: Missing parameters/Incorrect number of parameters passed", nil)
		service.GetHelp()
		return
	}
	// 系统初始化
	comm.InitRootWorkSpace("alog")

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
