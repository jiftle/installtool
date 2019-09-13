package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"os"
	"time"
)

var (
	plat_war = "platform.war"
)

func func_install_plat() {
	// 修改配置文件
	func_modify_conf_plat()
	func_copy_file_plat()
	func_system_service_tomcat_restart()
}
func func_uninstall_plat() (bool, string) {
	b := func_check_plat_isinstalled()
	if !b {
		return b, "未安装"
	}
	func_del_file_plat()
	return true, ""
}
func func_modify_conf_plat() {
	// 从war中抽离配置文件
	var strCmd string
	var strServiceWar string
	var strConfigFile string

	strServiceWar = strCurPath + "files/apps/" + plat_war
	strConfigFile = fmt.Sprintf("%v%v", strCurPath, "WEB-INF/classes/")

	// 抽取配置文件
	strCmd = fmt.Sprintf("jar -xf %v WEB-INF/classes/dbconfig.properties", strServiceWar)
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal("抽取配置文件失败，%s", err)
	}

	// ------------ 从标准输入得到用户的配置参数
	var strIP string
	var strPort string
	var strPwd string
	var strUser string
	var strDbName string

	for {
		// ------------ 从标准输入得到用户的配置参数
		fmt.Printf("--------------配置参数----------\n")
		fmt.Printf("|| 请输入数据库IP地址：\n")
		for {
			strIP = func_scan_input()
			b := func_ip_isvalid(strIP)
			if !b {
				fmt.Println("--> [x] ip地址不合法，请再次输入：")
				continue
			}
			break
		}
		fmt.Printf("|| 请输入数据库端口（默认3306）：\n")
		strPort = func_scan_input()

		fmt.Printf("|| 请输入数据库名称：\n")
		strDbName = func_scan_input()

		fmt.Printf("|| 请输入用户名：\n")
		strUser = func_scan_input()

		fmt.Printf("|| 请输入密码：\n")
		strPwd = func_scan_input()

		//--- 输出参数
		fmt.Printf("--> [----配置参数信息----]\nip: %v\nport: %v\ndbname: %v\nuser: %v\npwd: %v\n-----------------\n", strIP, strPort, strDbName, strUser, strPwd)

		fmt.Println("--> || 请确认配置参数是否正确无误[y/n],确认输入y,重新配置输入n:")
		input := func_scan_input()
		if input == "y" {
			break
		}
		continue
	}

	fmt.Println("--> 正在写入配置")
	// 拼凑写入参数
	var strConfigJdbc string
	var strCnfFile string

	strCnfFile = fmt.Sprintf("%sdbconfig.properties", strConfigFile)
	// 参数写入配置文件
	strConfigJdbc = fmt.Sprintf("echo \"hibernate.dialect=org.hibernate.dialect.MySQLDialect\" > %s", strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
	strConfigJdbc = fmt.Sprintf("echo \"validationQuery.sql=SELECT 1\" >> %s", strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
	strConfigJdbc = fmt.Sprintf("echo \"jdbc.url.jeecg=jdbc:mysql://%s:%s/%s?useUnicode=true&characterEncoding=UTF-8\" >> %s", strIP, strPort, strDbName, strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
	strConfigJdbc = fmt.Sprintf("echo \"jdbc.username.jeecg=%s\" >> %s", strUser, strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
	strConfigJdbc = fmt.Sprintf("echo \"jdbc.password.jeecg=%s\" >> %s", strPwd, strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}

	strConfigJdbc = fmt.Sprintf("echo \"jdbc.dbType=mysql\" >> %s", strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
	// 更新配置文件 1
	strCmd = fmt.Sprintf("jar uf %v %s", strServiceWar, "WEB-INF/classes/dbconfig.properties")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
	// 删除临时配置文件 1
	strCmd = fmt.Sprintf("rm -rf %s", "WEB-INF/")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal(err)
	}
}
func func_copy_file_plat() {
	strWebAppFile := fmt.Sprintf("%s%s%s", strCurPath, "files/apps/", plat_war)
	// 复制文件到tomcat目录下
	strCmd := fmt.Sprintf("cp %s %s%s", strWebAppFile, tomcat_install_dir, "tomcat/webapps/")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal("[ERROR] copy war file fail")
	}
	time.Sleep(5 * time.Second)
	// 修改目录权限
	strCmd = fmt.Sprintf("chmod 777 -R %s%s", tomcat_install_dir, "tomcat")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal("[ERROR] modify dir auth fail")
	}
}
func func_del_file_plat() {
	// 删除文件
	strCmd := fmt.Sprintf("rm %s%s%s", tomcat_install_dir, "tomcat/webapps/", plat_war)
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Fatal("[ERROR] del war file fail")
	}
}

func func_check_plat_isinstalled() bool {
	dir := fmt.Sprintf("%s%s%s", tomcat_install_dir, "tomcat/webapps/", plat_war)

	// 检查平台是否安装
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Info("file: %v isn't exist.\n", dir)
		return false
	}
	return true
}
