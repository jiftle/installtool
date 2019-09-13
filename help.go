package main

import (
	"fmt"
	"io/ioutil"
)

func showHelpFile() {
	var strHelpContent string
	buf, err := ioutil.ReadFile("files/help/readme.txt")
	if err != nil {
		//logger.Printf("[ERROR] read help file fail.")
		return
		// panic(err)
	}
	strHelpContent = string(buf)
	fmt.Println(strHelpContent)
}
