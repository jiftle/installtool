package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"install/function"
	"os"
	"time"
)

// 定义全局变量
var (
	strCurPath         = fmt.Sprintf("%s/", function.GetCurPath())
	tomcat_install_dir = "/usr/local/webplatform/"
)

func main() {
	defer logger.Close()

	// 环境检查
	b := func_sub_env_check()
	if !b {
		return
	}

	// 创建安装目录和计划任务
	func_install_pre()

	// 引导界面
	showInstallGui()
}

func init() {
	logger_init()
}

func logger_init() {
	logger.LoadConfiguration("conf/log4go.xml")
}

// -------------- 环境检查 ---------------
func func_sub_env_check() bool {
	fmt.Println("--> 开始进行环境检查")

	// 环境检测
	check_env()

	// 检查是否是root用户
	b := func_checkAuth()
	if !b {
		return false
	}

	return true
}

// 安装前准备
func func_install_pre() {

	// 创建安装目录
	func_create_install_dir()

	//创建计划任务
	func_create_cron_auto_clean_job()
}

// 安装选择界面
func showInstallGui() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!! 欢迎使用道统先生安装程序！!!!!!!!!!!!!!!!!!!!!!!!!!\n")

	var option int
	for {
		fmt.Println()
		fmt.Println("============ 请选择要**安装/卸载**的应用程序：============")
		fmt.Println("[0]	    查看帮助文档")
		fmt.Println("------------------------------------")
		fmt.Println("[1]	    安装++ 道统先生管理平台")
		fmt.Println("------------------------------------")
		fmt.Println("[-1]	    卸载-- 道统先生管理平台")
		fmt.Println("------------------------------------")
		fmt.Println("[2]	    安装++ 道统先生前置服务")
		fmt.Println("------------------------------------")
		fmt.Println("[-2]	    卸载-- 道统先生前置服务")
		fmt.Println("------------------------------------")
		fmt.Println("[8001]     检查安装环境")
		fmt.Println("------------------------------------")
		fmt.Println("[9001]     安装++ jdk")
		fmt.Println("------------------------------------")
		fmt.Println("[-9001]    卸载-- jdk")
		fmt.Println("------------------------------------")
		fmt.Println("[9002]     安装++ tomcat")
		fmt.Println("------------------------------------")
		fmt.Println("[-9002]    卸载-- tomcat")
		fmt.Println("------------------------------------")
		fmt.Println("[100]	    退出安装导引程序")
		fmt.Println("==========================================================")

		if _, err := fmt.Scanf("%d", &option); err != nil {
			fmt.Printf("--> [出现错误] %s，输入项不合法!\n", err)
			continue
		}

		fmt.Printf("<-- 你选择了[%d]选项，等待...\n", option)
		switch option {
		case 0:
			showHelpFile()
		case 1:
			func_sub_install_plat()
		case -1:
			func_sub_uninstall_plat()
		case 2:
			func_sub_install_service()
		case -2:
			func_sub_uninstall_service()
		case 8001:
			func_sub_env_check()
		case 9001:
			func_sub_install_jdk()
		case -9001:
			func_sub_uninstall_jdk()
		case 9002:
			func_sub_install_tomcat()
		case -9002:
			func_sub_uninstall_tomcat()
		case 100:
			func_sub_exit()
		default:
			fmt.Println("输入的选项不正确，请重新输入选择项!")
		}
	}
}

// -------------- jdk --------------------
func func_sub_install_jdk() {
	fmt.Println("--> 开始安装JDK，请耐心等待 ...")
	b, err := func_install_jdk()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Printf("--> [x] 安装JDK失败，原因：%s\n", err)
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Println("--> JDK安装完成")
}
func func_sub_uninstall_jdk() {
	fmt.Println("--> 开始卸载JDK，请耐心等待 ...")
	b, err := func_uninstall_jdk()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Printf("--> [x] 卸载JDK失败，原因：%s\n", err)
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Println("--> JDK卸载完成")
}

func func_install_step1() bool {
	b, _ := func_check_jdk_passmuster()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Println("--> [x] jdk没有安装，请先安装jdk")
		return false
	}
	b, _ = func_check_tomcat_passmuster()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Println("--> [x] tomcat没有安装，请先安装tomcat")
		return false
	}
	return true
}

// -------------- 道统先生管理平台 -----------------
func func_sub_install_plat() {
	b := func_install_step1()
	if !b {
		time.Sleep(2 * time.Second)
		return
	}

	b = func_check_plat_isinstalled()
	if b {
		time.Sleep(2 * time.Second)
		fmt.Println("--> [x] 管理平台已安装，如需重新安装请先卸载!")
		return
	}

	fmt.Println("--> 开始安装道统先生管理平台，请耐心等待 ...")
	func_install_plat()
	time.Sleep(3 * time.Second)
	fmt.Println("--> 道统先生管理平台安装完成")

	// ----------------- 显示登录页面地址 ---------------
	fmt.Println()
	fmt.Println("------------------------------------")
	fmt.Printf("管理平台登录地址（复制保存到本地，备登录平台使用）\n")
	n := func_get_webapp_url("platform/loginController.do?login")
	if n > 1 {
		fmt.Printf("提示： (存在多个访问地址，请选择实际需要使用的URL地址)\n")
	}
	fmt.Println("------------------------------------")
	fmt.Println()
}
func func_sub_uninstall_plat() {
	fmt.Println("--> 开始卸载道统先生管理平台，请耐心等待 ...")
	b, err := func_uninstall_plat()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Println("--> [x] 卸载管理平台失败，原因：" + err)
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println("--> 道统先生管理平台卸载完成")
}

// -------------- 道统先生前置服务 -----------------
func func_sub_install_service() {
	b := func_install_step1()
	if !b {
		time.Sleep(2 * time.Second)
		return
	}

	b = func_check_service_isinstalled()
	if b {
		time.Sleep(2 * time.Second)
		fmt.Println("--> [x] 前置服务已安装，如需重新安装请先卸载!")
		return
	}
	fmt.Println("--> 开始安装道统先生前置服务，请耐心等待 ...")
	func_install_service()
	time.Sleep(5 * time.Second)
	fmt.Println("--> 道统先生前置服务安装完成")

	// ----------------- 显示登录页面地址 ---------------
	fmt.Println()
	fmt.Println("------------------------------------")
	fmt.Printf("道统先生前置服务地址（复制保存到本地，配置到客户端参数）\n")
	n := func_get_webapp_url("serviceweb/")
	if n > 1 {
		fmt.Printf("提示： (存在多个访问地址，请选择实际需要使用的URL地址)\n")
	}
	fmt.Println("------------------------------------")
	fmt.Println()
}
func func_sub_uninstall_service() {
	fmt.Println("--> 开始卸载道统先生前置服务，请耐心等待 ...")
	b, err := func_uninstall_service()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Printf("--> [x] 卸载道统先生前置服务失败，原因：%s\n", err)
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println("--> 道统先生前置服务卸载完成")
}

// -------------- Tomcat -----------------
func func_sub_install_tomcat() {
	fmt.Println("--> 开始安装Tomcat，请耐心等待 ...")
	b, err := func_install_tomcat()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Printf("--> [x] 安装Tomcat失败，原因：%s\n", err)
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println("--> Tomcat安装完成")
}
func func_sub_uninstall_tomcat() {
	fmt.Println("--> 开始卸载Tomcat，请耐心等待 ...")
	b, err := func_uninstall_tomcat()
	if !b {
		time.Sleep(2 * time.Second)
		fmt.Printf("--> [x] 卸载Tomcat失败，原因：%s\n", err)
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println("--> Tomcat卸载完成")
}

// ------------- 退出安装工具 ------------
func func_sub_exit() {
	time.Sleep(1 * time.Second)
	fmt.Println("--> 安装程序退出")
	os.Exit(1)
}
