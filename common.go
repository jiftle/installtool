package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"install/function"
	"net"
	"os"
	"os/exec"
)

func func_create_install_dir() {
	dir := "/usr/local/webplatform"

	// 判断目录是否存在，如果不存在就创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Warn("dir: %v isn't exist.", dir)
	} else {
		return
	}

	// 创建目录
	strCmd := fmt.Sprintf("mkdir -p %v", dir)
	_, err := function.Func_exec_cmd(strCmd)
	if err != nil {
		logger.Error("create dir: %v fail", dir)
		panic(err)
	}
	logger.Info("create dir: %v success.", dir)
}

// 执行命令
func func_exec_cmd(strCmd string) (bool, error) {
	cmd := exec.Command("/bin/bash", "-c", strCmd)
	output, err := cmd.Output()
	if err != nil {
		logger.Error("exec cmd: \"%s\", execute fail. %v\n", strCmd, err)
		return false, err
	}
	logger.Info("exec cmd： \"%v\", execute success. \nexec result:\n---------------------------\n%s\n------------------------------\n", strCmd, output)
	return true, nil
}

// 执行命令，需要返回输出结果
func func_exec_cmd_output(strCmd string) (bool, string, error) {
	cmd := exec.Command("/bin/bash", "-c", strCmd)
	output, err := cmd.Output()
	if err != nil {
		logger.Error("exec cmd：\"%s\", execute fail. error: \"%v\"\n", strCmd, err)
		return false, "", err
	}
	logger.Info("exec cmd： \"%v\", execute success. \nexec result:\n---------------------------\n%s\n------------------------------\n", strCmd, output)
	return true, string(output), nil
}

func func_get_webapp_root_url() []string {
	var sliceWebUrl []string
	ips := function.Func_get_local_ip()
	sliceWebUrl = ips
	for i, ip := range ips {
		sliceWebUrl[i] = fmt.Sprintf("http://%s:%d/", ip, 8080)
		//fmt.Println(sliceWebUrl[i])
	}
	return sliceWebUrl
}
func func_get_webapp_url(url string) int {
	webUrls := func_get_webapp_root_url()
	for _, webUrl := range webUrls {
		surl := fmt.Sprintf("URL: %s%s", webUrl, url)
		fmt.Println(surl)
	}
	return len(webUrls)
}

func func_scan_input() string {
	var strInput string
	for {
		fmt.Scanln(&strInput)
		if strInput == "" {
			fmt.Println("--> [x] 输入参数为空，请再次输入:")
		} else {
			break
		}
	}
	return strInput
}
func func_ip_isvalid(ip string) bool {
	address := net.ParseIP(ip)
	if address != nil {
		return true
	} else {
		return false
	}
}
