package main

import (
	"fmt"
	"install/logger"
	"install/softmanager"
)

// 安装jdk
func func_install_jdk() (bool, string) {
	// 检查jdk是否安装
	bJdkExisted := softmanager.GetJdkIsInstalled()
	if !bJdkExisted {
		logger.Error("检测到jdk未安装")
	}

	// 安装jdk
	strCmd := fmt.Sprintf("yum install %s%sjdk-8u131-linux-x64.rpm -y", strCurPath, "files/system-apps/")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Errorf("[ERROR] install jdk fail")
		return false, "install jdk fail"
	}
	return true, ""
}

// 卸载jdk
func func_uninstall_jdk() (bool, string) {
	//--------- 检查不了
	//	// 检查jdk是否安装
	//	bJdkExisted := softmanager.GetJdkIsInstalled()
	//	if !bJdkExisted {
	//		logger.Error("检测到jdk未安装")
	//		return false, "JDK未安装"
	//	}

	// 卸载jdk
	strCmd := fmt.Sprintf("yum remove jdk -y")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Infof("[ERROR] remove jdk fail")
		return false, "remove jdk fail"
	}

	// 卸载openjdk
	strCmd = fmt.Sprintf("yum remove *openjdk* -y")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Infof("[ERROR] remove jdk fail")
		return false, "remove jdk fail"
	}
	// 卸载之前的jdk,jre
	strCmd = fmt.Sprintf("yum remove java -y")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Infof("[ERROR] remove jre fail")
		return false, "remove jdk fail"
	}

	return true, ""
}

// 是否安装过jdk
func get_isinstalled_jdk(osid string) (bool, string) {
	var errmsg string
	if "centos" == osid {
		strCmd := fmt.Sprintf("rpm -qa|grep java")
		_, err := func_exec_cmd(strCmd)
		if err != nil {
			errmsg = "rpm query installed jdk fail"
			logger.Infof("[ERROR] %s", errmsg)
			return false, errmsg
		}

		// --------- rpm 方式 查询软件是否安装 -------------
		strCmd = fmt.Sprintf("yum list installed|grep java")
		_, err = func_exec_cmd(strCmd)
		if err != nil {
			logger.Infof("[ERROR] yum read jdk fail")
			return false, "yum read jdk fail"
		}
	} else if "deepin" == osid {
		strCmd := fmt.Sprintf("rpm -qa|grep java")
		_, err := func_exec_cmd(strCmd)
		if err != nil {
			errmsg = "rpm query installed jdk fail"
			logger.Infof("[ERROR] %s", errmsg)
			return false, errmsg
		}
	} else {
	}

	return true, ""

}
