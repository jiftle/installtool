package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"install/function"
	"os"
)

func func_create_cron_auto_clean_job() (bool, string) {
	file := "clean_log.sh"
	dir := fmt.Sprintf("%scron/", tomcat_install_dir)
	cronfile := "/var/spool/cron/root"
	srcfile := fmt.Sprintf("%sfiles/cron/%s", strCurPath, file)
	destfile := fmt.Sprintf("%s%s", dir, file)
	var errmsg string

	// 判断目录是否存在，如果不存在就创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errmsg = fmt.Sprintf("cron dir: %s isn't exist.\n", dir)
		logger.Warn(errmsg)

		// 创建目录
		strCmd := fmt.Sprintf("mkdir -p %s", dir)
		_, err := function.Func_exec_cmd(strCmd)
		if err != nil {
			errmsg = fmt.Sprintf("create cron dir: %s fail", dir)
			logger.Error("[ERROR] %s", errmsg)
			//	panic(err)
			return false, errmsg
		}
	}

	// 复制文件到目录
	strCmd := fmt.Sprintf("cp -f %s %s", srcfile, destfile)
	_, err := function.Func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("copy file: %s --> %s fail", srcfile, destfile)
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}

	// 加入执行权限
	strCmd = fmt.Sprintf("chmod +x %s", destfile)
	_, err = function.Func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("copy file: %s --> %s fail", srcfile, destfile)
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}

	// -----------------------------------------------------
	// 判断计划任务配置文件是否存在
	if _, err := os.Stat(cronfile); os.IsNotExist(err) {
		logger.Info("cron file: %s isn't exist.\n", cronfile)

		// 创建目录
		strCmd := fmt.Sprintf("touch %s", cronfile)
		_, err := function.Func_exec_cmd(strCmd)
		if err != nil {
			errmsg = fmt.Sprintf("create file: %s fail", cronfile)
			logger.Error("[ERROR] %s", errmsg)
			return false, errmsg
		}
	}

	// 加入计划任务
	strCmd = fmt.Sprintf("echo \"0 1 * * * %s\" > %s", destfile, cronfile)
	_, err = function.Func_exec_cmd(strCmd)
	if err != nil {
		errmsg = fmt.Sprintf("append cron task fail")
		logger.Error("[ERROR] %s", errmsg)
		return false, errmsg
	}

	return true, ""
}
