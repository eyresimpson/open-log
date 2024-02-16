package comm

import (
	"open-log/tools"
	"os"
	"path"
)

// 初始化运行空间
// 用于初始化主目录和记述文件
func InitRootWorkSpace(confName string) {
	confName = tools.GetCurrentUserRootDirectory() + "/.aine/" + confName + "/root.json"
	// 判断配置文件是否存在
	if !tools.IsFileExist(confName) {
		os.MkdirAll(path.Dir(confName), 0755)
		os.Create(confName)
		tools.WriteFile("[]", confName)
	}

}
