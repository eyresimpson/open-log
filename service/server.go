package service

import (
	"open-log/tools"
)

var version = "0.0.1 Feb 16, 2024"

// 获取版本
func GetVersion() {
	tools.Info("Open Log Version " + version)
	tools.Info("Creator: Noah Jones")
}

// 获取帮助
func GetHelp() {
	tools.Warn("Open Log Version " + version)
	tools.Info("log [-h|-v|-r|-f] file1 file2 file3 ...")
	tools.Warn("Amenity")
	tools.Info("    Enter the file you want to open-logete directly")
	tools.Warn("Base")
	tools.Info("    -h	 	: show help")
	tools.Info("    -v		: show version")
	tools.Info("    -l		: Lists all open-logeted files")

}

// 功能检查
func Run(opt string, args []string) {

	// 判断参数
	// switch opt {
	// case "-r":
	// 	Recover(args)
	// case "-f":
	// 	ReallyRemove(args)
	// default:
	// 	Remove(args)
	// }
}
