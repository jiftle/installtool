package function

import (
	"log"
	"os"
	"path/filepath"
)

func GetCurPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//logger.Printf("-----> current dir is: \"%s\"", dir)
	return dir
}
