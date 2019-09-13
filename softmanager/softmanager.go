package softmanager

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"install/function"
)

func init() {
	logger.LoadConfiguration("conf/log4go.xml")
}

// 检查JDK是否安装
func GetJdkIsInstalled() bool {
	var n int
	var out string
	var err error
	softname := "jdk"
	osver := function.GetOsVer()

	if osver.Id == "centos" {
		n, out, err = GetInstalledSoftByRpm(softname)
	} else {
		n, out, err = GetInstalledSoftByDpkg_debian(softname)
	}
	if err == nil {
		logger.Error("获取JDK信息失败,%v", err)
	}
	if n > 0 {
		logger.Info("JDK存在多个:\n%s", out)
		return true
	} else {
		return false
	}
}

// 检查Jar是否安装
func GetJarIsInstallled() bool {
	strCmd := fmt.Sprintf("jar")
	_, output, err := function.Exec_cmd_output(strCmd)
	if err != nil {
		if len(output) == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}
