package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"os"
)

func func_install_service() {
	//---------------------------------------
	// 步骤
	// 1. 修改配置
	// 2. 复制文件
	// 3. 启动服务
	// 4. 查看结果
	//---------------------------------------
	func_modify_conf()
	func_copy_file()
	func_system_service_tomcat_restart()
}
func func_uninstall_service() (bool, string) {
	b := func_check_service_isinstalled()
	if !b {
		return false, "未安装"
	}
	func_del_file()
	return true, ""
}
func func_modify_conf() (bool, string) {
	//---------------------------------------
	// 步骤
	// 从war中抽离配置文件
	// 从标准输入得到用户的配置参数
	// 写入配置参数到配置文件
	// 更新配置文件到war包
	// 复制war到指定目录
	//---------------------------------------

	// 从war中抽离配置文件
	var strCmd string
	var strServiceWar string
	var strConfigFile string
	var errmsg string

	strServiceWar = strCurPath + "files/apps/serviceweb.war"
	strConfigFile = fmt.Sprintf("%v%v", strCurPath, "WEB-INF/classes/")

	// 抽取配置文件
	strCmd = fmt.Sprintf("jar -xf %v WEB-INF/classes/jdbc.properties", strServiceWar)
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("extract config file fail, %s", err)
		logger.Error(errmsg)
		return false, errmsg
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
	var strConfigUser string
	var strConfigPwd string
	var strCnfFile string

	strCnfFile = fmt.Sprintf("%sjdbc.properties", strConfigFile)
	// 参数写入配置文件
	strConfigJdbc = fmt.Sprintf("sed -i \"s#^jdbc.url=.*#jdbc.url=jdbc:mysql://%s:%s/%s?useUnicode=true\\&characterEncoding=utf-8\\&autoReconnect=true#g\" %s", strIP, strPort, strDbName, strCnfFile)
	strCmd = strConfigJdbc
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Error(errmsg)
		return false, errmsg
	}

	strConfigUser = fmt.Sprintf("sed -i \"s#^jdbc.username=.*#jdbc.username=%s#g\" %s", strUser, strCnfFile)
	strCmd = strConfigUser
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Error(errmsg)
		return false, errmsg
	}

	strConfigPwd = fmt.Sprintf("sed -i \"s#^jdbc.password=.*#jdbc.password=%s#g\" %s", strPwd, strCnfFile)
	strCmd = strConfigPwd
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Error(errmsg)
		return false, errmsg
	}

	// 复制配置文件 -- 不知道使用哪个配置文件，所以两个文件都更新
	strCmd = fmt.Sprintf("cp %s %s", "WEB-INF/classes/jdbc.properties", "jdbc.properties")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Error(errmsg)
		return false, errmsg
	}
	// 更新配置文件 1
	strCmd = fmt.Sprintf("jar uf %v %s", strServiceWar, "WEB-INF/classes/jdbc.properties")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Info(errmsg)
		return false, errmsg
	}
	// 更新配置文件 2
	strCmd = fmt.Sprintf("jar uf %v %s", strServiceWar, "jdbc.properties")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Info(errmsg)
		return false, errmsg
	}
	// 删除临时配置文件 1
	strCmd = fmt.Sprintf("rm -rf %s", "WEB-INF/")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Info(errmsg)
		return false, errmsg
	}
	// 删除临时配置文件 2
	strCmd = fmt.Sprintf("rm %s", "jdbc.properties")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("exec command fail, %s", err)
		logger.Info(errmsg)
		return false, errmsg
	}
	return true, ""
}
func func_copy_file() {
	strWebAppFile := fmt.Sprintf("%s%s", strCurPath, "files/apps/serviceweb.war")
	// 复制文件到tomcat目录下
	strCmd := fmt.Sprintf("cp %s %s%s", strWebAppFile, tomcat_install_dir, "tomcat/webapps/")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] copy war file fail")
	}
	// 修改目录权限
	strCmd = fmt.Sprintf("chmod 777 -R %s%s", tomcat_install_dir, "tomcat")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] modify dir auth fail")
	}
}
func func_del_file() {
	// 删除文件
	strCmd := fmt.Sprintf("rm %s%s", tomcat_install_dir, "tomcat/webapps/serviceweb.war")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] del war file fail")
	}
}
func func_check_service_isinstalled() bool {
	dir := fmt.Sprintf("%s%s%s", tomcat_install_dir, "tomcat/webapps/", "serviceweb.war")

	// 判断目录是否存在，如果不存在就创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Info("dir: %v isn't exist.\n", dir)
		return false
	}
	return true
}
