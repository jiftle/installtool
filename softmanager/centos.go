package softmanager

import (
	//	"bytes"
	"errors"
	"fmt"
	"install/function"
	"strings"

	logger "github.com/ccpaging/log4go"
)

// 查询已安装的软件
// 返回值 0，没有，非0，表示存在安装包的个数
func GetInstalledSoftByRpm(softName string) (int, string, error) {
	var errmsg string
	var count int

	strCmd := fmt.Sprintf("rpm -qa | grep %s", softName)
	_, output, err := function.Func_exec_cmd_output(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("rpm query installed %s fail", softName)
		logger.Error("%s", errmsg)
		return 0, "", errors.New(errmsg)
	}

	lines := strings.Split(output, "\n")
	logger.Info("--> 已安装的软件包的数量=%d\n", len(lines)-1)

	return count, output, nil
}

// 查询已安装的软件
// 返回值 0，没有，非0，表示存在安装包的个数
func GetInstalledSoftByYum(softName string) (int, string, error) {
	var errmsg string
	var count int

	strCmd := fmt.Sprintf("yum list installed |grep %s", softName)
	_, output, err := function.Func_exec_cmd_output(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("yum list installed %s fail", softName)
		logger.Error("[ERROR] %s", errmsg)
		return 0, "", errors.New(errmsg)
	}

	lines := strings.Split(output, "\n")
	logger.Info("--> 已安装的软件包的数量=%d\n", len(lines)-1)

	return count, output, nil
}
