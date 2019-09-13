package function

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/liuyongshuai/goUtils"
)

type Osver struct {
	Id      string
	Name    string
	Version string
	Detail  string
}

func (o *Osver) ToString() string {
	var s string
	s = o.Detail
	return s
}

func GetJdkVer() string {
	var output []string = make([]string, 1024)
	var ret int

	goUtils.ExecCmd("java -version", &output, &ret)

	var strJdkVersion = output[0]

	return strJdkVersion
}

// 操作系统版本信息
func GetOsVer() Osver {

	var osver Osver

	f, err := os.Open("/etc/os-release")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		//fmt.Printf("--> %v", line)
		kv := strings.Split(line, "=")

		//fmt.Printf("--| %v: %v", kv[0], kv[1])
		if "NAME" == kv[0] {
			osver.Name = kv[1][1 : len(kv[1])-2]
		} else if "VERSION" == kv[0] {
			osver.Version = kv[1][1 : len(kv[1])-2]
		} else if "ID" == kv[0] {
			value := kv[1][0 : len(kv[1])-1]
			//	fmt.Printf("--|[%s]\n", value)
			if value[0:1] == "\"" {
				osver.Id = value[1 : len(value)-1]
			} else {
				osver.Id = value
			}
		} else {
		}
	}

	//	fmt.Printf("osver: \n %v\n", osver)
	//--------------- 取出详细的操作系统版本信息，eg. Centos 7.4
	if "deepin" == osver.Id {
		osver.Detail = get_deepin_ver()
	} else if "centos" == osver.Id {
		osver.Detail = get_centos_ver()
	} else {
		fmt.Printf("--> [x] 无效分支,(%s)\n", osver.Id)
	}

	return osver
}

func get_centos_ver() string {
	var version string

	buf, err := ioutil.ReadFile("/etc/redhat-release")
	if err != nil {
		panic(err)
	}

	version = fmt.Sprintf("%s", string(buf))
	//fmt.Printf("version: %s\n", version)
	return version
}
func get_deepin_ver() string {
	var name string
	var ver string
	var version string

	f, err := os.Open("/etc/os-release")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		kv := strings.Split(line, "=")

		if "NAME" == kv[0] {
			name = kv[1][1 : len(kv[1])-2]
		} else if "VERSION" == kv[0] {
			ver = kv[1][1 : len(kv[1])-2]
		} else {
		}
	}

	version = fmt.Sprintf("%s %s", name, ver)
	return version
}
func Func_kill_proc(procname string) {
	// 查询进程pid
	strCmd := fmt.Sprintf("ps -ef|grep \"%s\"|grep -v \"$0\"|grep -v \"grep\"|awk '{print $2}'", procname)
	_, pid, err := Func_exec_cmd_output(strCmd)
	if err != nil {
		//		logger.Fatal("[ERROR] lookup proc fail")
	}
	if pid == "" {
		return
	}
	pids := strings.Split(pid, "\n")

	for _, spid := range pids {
		if spid == "" {
			break
		}
		// kill 进程
		strCmd = fmt.Sprintf("kill -9 %s", spid)
		_, err = Func_exec_cmd(strCmd)
		if err != nil {
			//			logger.Fatal("[ERROR] kill proc fail")
		}
	}

}

// 获取ip地址
func Func_get_local_ip() []string {
	var sliceIp []string
	sliceIp = make([]string, 5)
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}

	var count int = 0
	for i, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//		fmt.Printf("--- %v ----> loop 地址\n", i)
				//				fmt.Println(ipnet.IP.String())
				sliceIp[i-1] = ipnet.IP.String()
				count++
			}

		}
	}

	//fmt.Printf("--> count=%d\n", count)
	// 计算实际的长度
	ret_slice := sliceIp[0:count]
	return ret_slice
}
