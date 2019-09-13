package function

import (
	"os/exec"
	"regexp"
	"strings"

	logger "github.com/ccpaging/log4go"
)

func init() {
	logger_init()
}

func logger_init() {
	logger.LoadConfiguration("conf/log4go.xml")
}
func Func_exec_cmd(strCmd string) (bool, error) {
	cmd := exec.Command("/bin/bash", "-c", strCmd)
	output, err := cmd.Output()
	if err != nil {
		logger.Error("exec cmd: \"%s\", execute fail. %v\n", strCmd, err)
		return false, err
	}
	logger.Info("exec cmd： \"%v\", execute success. \nexec result:\n---------------------------\n%s\n------------------------------\n", strCmd, output)
	return true, nil
}

// 执行命令，需要返回输出结果
func Exec_cmd_output(strCmd string) (bool, string, error) {
	cmd := exec.Command(strCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("exec cmd：\"%s\", execut fail.\nerror: \n\"%v\"\nout: \n\"%s\"", strCmd, err, output)
		return false, string(output), err
	}
	logger.Info("exec cmd： \"%v\", execute success. \nexec result:\n---------------------------\n%s\n------------------------------\n", strCmd, output)
	return true, string(output), nil
}

// 执行命令，需要返回输出结果
func Func_exec_cmd_output(strCmd string) (bool, string, error) {
	cmd := exec.Command("/bin/bash", "-c", strCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("exec cmd：\"%s\", execute fail.\n:error: \"%v\"\nout: \"%s\"\n", strCmd, err, output)
		return false, "", err
	}
	logger.Info("exec cmd： \"%v\", execute success. \nexec result:\n---------------------------\n%s\n------------------------------\n", strCmd, output)
	return true, string(output), nil
}

//returnVar：0成功、1失败
func ExecCmd(command string, output *[]string, returnVar *int) string {
	r, _ := regexp.Compile(`[ ]+`)
	parts := r.Split(command, -1)
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	cmd := exec.Command(parts[0], args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		*returnVar = 1
		return ""
	} else {
		*returnVar = 0
	}
	*output = strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	if l := len(*output); l > 0 {
		return (*output)[l-1]
	}
	return ""
}
