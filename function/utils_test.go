package function

import (
	"fmt"
	"testing"
)

func TestGetOsVer(t *testing.T) {
	osver := getOsVer()
	s := fmt.Sprintf("---------------\n%v\n", osver)
	t.Log(s)
}
