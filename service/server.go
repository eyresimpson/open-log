package service

import (
	"alog/tools"
	"alog/tools/logger"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var version = "0.0.1 Feb 16, 2024"

// 获取版本
func GetVersion() {
	logger.Info("Open Log Version " + version)
	logger.Info("Creator: Noah Jones")
}

// 获取帮助
func GetHelp() {
	logger.Warn("Open Log Version " + version)
	logger.Info("log [-h|-v|-r|-f] file1 file2 file3 ...")
	logger.Warn("Amenity")
	logger.Info("    Enter the file you want to alogete directly")
	logger.Warn("Base")
	logger.Info("    -h	 	: show help")
	logger.Info("    -v		: show version")
	logger.Info("    -l		: Lists all alogeted files")

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
	LoggerShell(args)
}

func LoggerShell(args []string) {
	for _, exec := range args {
		println("Running", exec)
		go executeCommand(exec, "log.txt")
	}
}

// func executeCommand(prefix, command string) {
// 	cmd := exec.Command(command)

// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		fmt.Println("Error creating stdout pipe:", err)
// 		return
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		fmt.Println("Error starting command:", err)
// 		return
// 	}

// 	reader := bufio.NewReader(stdout)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			} else {
// 				fmt.Println("Error reading command output:", err)
// 				break
// 			}
// 		}
// 		fmt.Print(line)
// 	}

// 	err = cmd.Wait()
// 	if err != nil {
// 		fmt.Println("Error waiting for command to finish:", err)
// 		return
// 	} else {
// 		logger.Success("Command finished!")
// 	}
// }

func executeCommand(command string, outputFile string) error {
	var cmd *exec.Cmd
	if tools.IsWindows() {
		cmd = exec.Command("bash", "-c", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(out.String())
	if err != nil {
		return err
	}

	fmt.Printf("Command output has been written to %s\n", outputFile)

	return nil
}
