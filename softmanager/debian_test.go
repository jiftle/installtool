package softmanager

import (
	"fmt"
	"testing"
)

// 检查JDK是否安装
func TestGetJdkIsInstallled(t *testing.T) {
	b := GetJdkIsInstallled()
	//	if !b {
	//		//		t.Error("jdk没有安装")
	//	} else {
	//		t.Info("jdk已安装")
	//	}
	fmt.Printf("check jdk : %v\n", b)
}
