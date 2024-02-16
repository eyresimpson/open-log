package service

import (
	"open-log/tools"
	"os"
)

// 获取版本
func GetVersion() {
	tools.Info("TLog Version 0.1.0 Feb 15, 2024")
	tools.Info("Creator: Noah Jones")
}

// 获取帮助
func GetHelp() {
	tools.Warn("open-log Version 0.1.0 Feb 15, 2024")
	tools.Info("open-log [-h|-v|-r|-f] file1 file2 file3 ...")
	tools.Warn("Amenity")
	tools.Info("    Enter the file you want to open-logete directly")
	tools.Warn("Base")
	tools.Info("    -h	 	: show help")
	tools.Info("    -v		: show version")
	tools.Info("    -f	 	: Forced open-logetion")
	// 列出所有已删除
	tools.Info("    -l		: Lists all open-logeted files")
	// 恢复指定文件
	tools.Info("    -r		: Recover specified file, use id")
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

// 初始化运行空间
// 用于初始化暂存目录和记述文件
func InitRootWorkSpace(dumpsterConf string) {
	// 判断配置文件是否存在
	if !tools.IsFileExist(dumpsterConf) {
		// 获取回收站目录
		dumpster := tools.GetCurrentUserRootDirectory() + "/.dumpster"
		os.MkdirAll(dumpster, 0755)
		os.Create(dumpsterConf)
		tools.WriteFile("[]", dumpsterConf)
	}

}
