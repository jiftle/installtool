package softmanager

import (
	"errors"
	"fmt"
	"install/function"
	"strings"

	logger "github.com/ccpaging/log4go"
)

// 查询已安装的软件，系统debian，软件包管理器dpkg
// 返回值 0，没有，非0，表示存在安装包的个数
func GetInstalledSoftByDpkg_debian(softName string) (int, string, error) {
	var errmsg string
	var count int

	strCmd := fmt.Sprintf("dpkg -l | grep %s", softName)
	_, output, err := function.Func_exec_cmd_output(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("dpkg query installed %s fail", softName)
		logger.Error("%s", errmsg)
		return 0, "", errors.New(errmsg)
	}

	lines := strings.Split(output, "\n")
	logger.Info("--> 已安装的软件包的数量=%d\n", len(lines)-1)

	return count, output, nil
}

// 查询已安装的软件，系统debian，软件包管理器apt-get
// 返回值 0，没有，非0，表示存在安装包的个数
func GetInstalledSoftByAptGet_debian(softName string) (int, string, error) {
	var errmsg string
	var count int

	strCmd := fmt.Sprintf("apt list --installed | grep %s", softName)
	_, output, err := function.Func_exec_cmd_output(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("dpkg query installed %s fail", softName)
		logger.Error("[ERROR] %s", errmsg)
		return 0, "", errors.New(errmsg)
	}

	lines := strings.Split(output, "\n")
	logger.Info("--> 已安装的软件包的数量=%d\n", len(lines)-1)

	return count, output, nil
}
