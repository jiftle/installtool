package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"github.com/liuyongshuai/goUtils"
	"install/function"
	"install/softmanager"
	"os"
	"os/user"
	"strings"
)

// 检查项目
// 1. 操作系统 2. JDK 3. Tomcat
func check_env() {
	osver := function.GetOsVer()
	var strOsVer string = osver.ToString()
	var strJdkVer string = function.GetJdkVer()

	//------------ 执行环境检查 ----------------
	//	fmt.Println("--> 开始执行环境检查")
	var table goUtils.TerminalTable
	headers := []string{"序号", "检测项", "输出值", "期望值", "是否合格"}
	table.SetHeader(headers)

	//插入记录
	// 操作系统
	bPassOs, strPassmuster := func_check_os_passmuster()
	row := []string{"1", "操作系统版本", strOsVer, "CentOS Linux 7", strPassmuster}
	table.AddRow(row)
	//插入记录
	// JDK
	bPassJdk, strPassmuster := func_check_jdk_passmuster()
	row = []string{"2", "JDK版本", strJdkVer, "JDK 1.8", strPassmuster}
	table.AddRow(row)
	//插入记录
	// Tomcat
	bPassTomcat, strPassmuster := func_check_tomcat_passmuster()
	row = []string{"3", "Tomcat", "/usr/local/webplatform/tomcat/", "目录存在", strPassmuster}
	table.AddRow(row)

	//输出表格
	fmt.Println(table.Render())

	//------------- 针对不合格的输出建议 ---------------
	if !bPassOs {
		fmt.Printf("*%s: %s\n", "操作系统建议", "更换操作系统或找个Centos7的系统")
	}
	if !bPassJdk {
		fmt.Printf("*%s: %s\n", "JDK建议     ", "卸载干净机器上存在的JDK，使用安装程序附带的JDK。查看系统中存在的JDK指令： rpm -qa|grep jdk")
	}
	if !bPassTomcat {
		fmt.Printf("*%s: %s\n", "Tomcat建议  ", "使用安装程序安装Tomcat")
	}
	fmt.Println()
}
func func_checkAuth() bool {
	//--------- 当前用户 -------------
	u, err := user.Current()
	if err != nil {
		logger.Error("get current user info fail. ")
		return false
	}

	// 不是root用户
	if u.Uid != "0" {
		fmt.Println("--> [X] must use root user operate. 请使用root用户安装!")
		logger.Info("[X] must use root user operate.")
		return false
	}

	return true
}

// 检查操操作系统
func func_check_os_passmuster() (bool, string) {
	osver := function.GetOsVer()
	var strOsVer string = osver.ToString()

	b := strings.Contains(strOsVer, "CentOS") && strings.Contains(strOsVer, "7")
	if b {
		return true, "合格"
	}
	return false, "不合格"
}

// 检查jdk是否存在
func func_check_jdk_passmuster() (bool, string) {
	var b bool
	var bJdkExisted bool
	var bJarExisted bool
	var bMoreJdk bool

	bJdkExisted = false
	bJarExisted = false

	// 检查jdk
	var strJdkVer string = function.GetJdkVer()
	bJdkExisted = strings.Contains(strJdkVer, "1.8") && (strings.Contains(strJdkVer, "jdk") || strings.Contains(strJdkVer, "java"))

	// 检查jar指令是否存在
	bJarExisted = softmanager.GetJarIsInstallled()
	if !bJarExisted {
		logger.Error("检测到jar命令未安装")
	}

	//检查是否存在多个jdk
	bMoreJdk = softmanager.GetJdkIsInstalled()
	if bMoreJdk {
		logger.Error("存在多个JDK")
	}

	//b = bJdkExisted && bJarExisted && bMoreJdk
	b = bJdkExisted && bJarExisted
	if b {
		return true, "合格"
	}
	return false, "不合格"
}

// 检查tomcat是否存在
func func_check_tomcat_passmuster() (bool, string) {
	dir := "/usr/local/webplatform/tomcat/"

	// 判断目录是否存在，如果不存在就创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Warn("dir: %v isn't exist.", dir)
	} else {
		return true, "合格"
	}
	return false, "不合格"
}
