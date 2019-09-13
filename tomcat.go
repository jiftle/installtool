package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"install/function"
	"os"
	"time"
)

func func_uninstall_tomcat() (bool, string) {
	var errmsg string
	b, _ := func_tomcat_isStalled()
	if !b {
		errmsg = fmt.Sprintf("未安装")
		logger.Info("%s", errmsg)
		return false, errmsg
	}
	// 卸载tomcat自启动服务
	func_system_service_tomcat_uninstall()

	// kill tomcat进程
	function.Func_kill_proc("tomcat")

	// 删除tomcat
	strCmd := fmt.Sprintf("rm -rf %stomcat/", tomcat_install_dir)
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("del tomcat tar file fail")
		logger.Info("%s", errmsg)
		return false, errmsg
	}
	// 删除安装目录
	strCmd = fmt.Sprintf("rm -rf %s", tomcat_install_dir)
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("del install dir fail")
		logger.Info("%s", errmsg)
		return false, errmsg
	}
	// 从防火墙中移除tomcat端口
	func_port_remove_firewalld()

	return true, ""
}
func func_install_tomcat() (bool, string) {
	var errmsg string
	// 复制文件
	tomcat_tar := fmt.Sprintf("%s%s", strCurPath, "files/system-apps/apache-tomcat-8.5.32.tar.gz")
	strCmd := fmt.Sprintf("cp %v %v", tomcat_tar, tomcat_install_dir)
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("copy tomcat tar: %v fail", tomcat_tar)
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}
	// 解压文件
	strTomcatFile := fmt.Sprintf("%s%s", tomcat_install_dir, "apache-tomcat-8.5.32.tar.gz")
	//strCmd = fmt.Sprintf("tar -xzvf %s -C %s", strTomcatFile, tomcat_install_dir)
	strCmd = fmt.Sprintf("tar -xzf %s -C %s", strTomcatFile, tomcat_install_dir)
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("excat tomcat files: %v fail", strTomcatFile)
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}
	// 修改目录名
	strCmd = fmt.Sprintf("mv %sapache-tomcat-8.5.32/ %stomcat/", tomcat_install_dir, tomcat_install_dir)
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("modify dir name fail")
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}
	// 删除安装包
	strCmd = fmt.Sprintf("rm %s", strTomcatFile)
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("del tomcat tar file fail")
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}
	// 修改目录权限
	strCmd = fmt.Sprintf("chmod 777 -R %s%s", tomcat_install_dir, "tomcat")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("modify dir auth fail")
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}

	//------------- 加入系统服务
	func_system_service_tomcat_install()

	//------------- 端口加入到防火墙
	func_port_add_firewalld()

	return true, ""
}

func func_port_add_firewalld() {
	strCmd := fmt.Sprintf("firewall-cmd --zone=public --add-port=8080/tcp --permanent")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] port 8080 add firewalld fail")
	}
	strCmd = fmt.Sprintf("firewall-cmd --reload")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] firewall reload fail")
	}

}
func func_port_remove_firewalld() {
	strCmd := fmt.Sprintf("firewall-cmd --zone=public --remove-port=8080/tcp --permanent")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] port 8080 remove firewalld fail")
	}
	strCmd = fmt.Sprintf("firewall-cmd --reload")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Info("[ERROR] firewall reload fail")
	}
}

// tomcat 安装成系统服务
func func_system_service_tomcat_install() {
	// 复制文件,服务配置行为
	tomcat_service_file := fmt.Sprintf("%s%s", strCurPath, "files/system-apps/sdmtomcatd.service")
	strCmd := fmt.Sprintf("cp -f %s /usr/lib/systemd/system/", tomcat_service_file)
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] copy tomcat auto service file fail")
	}
	// 载入
	strCmd = fmt.Sprintf("systemctl daemon-reload")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] start tomcat fail")
	}
	// 启动
	strCmd = fmt.Sprintf("systemctl start sdmtomcatd")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] start tomcat fail")
	}
	// 加入自启动
	strCmd = fmt.Sprintf("systemctl enable sdmtomcatd")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] start tomcat fail")
	}
}

// 卸载tomcat系统服务
func func_system_service_tomcat_uninstall() {
	// 移除服务自启动
	strCmd := fmt.Sprintf("systemctl disable sdmtomcatd")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] remove auto start tomcat service fail")
	}
	// 停止服务
	strCmd = fmt.Sprintf("systemctl stop sdmtomcatd")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] stop tomcat service fail")
	}
	// 重新加载服务
	strCmd = fmt.Sprintf("systemctl daemon-reload")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] system service reload fail")
	}
	// 删除自动服务文件
	strCmd = fmt.Sprintf("rm -f /usr/lib/systemd/system/%s", "sdmtomcatd.service")
	_, err = func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] remove tomcat auto service file fail")
	}
}
func func_system_service_tomcat_restart() {
	// 重新启动
	strCmd := fmt.Sprintf("systemctl restart sdmtomcatd")
	_, err := func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("[ERROR] restart tomcat fail")
	}

	time.Sleep(5 * time.Second)
}

func func_tomcat_isStalled() (bool, string) {
	dir := "/usr/local/sdmwebplatform/tomcat/"

	// 判断目录是否存在，如果不存在就创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Warn("dir: %v isn't exist.\n", dir)
	} else {
		return true, "已安装"
	}
	return false, "未安装"
}
