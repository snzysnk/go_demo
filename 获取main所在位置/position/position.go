package position

import (
	"fmt"
	"path"
	"runtime"
)

var goRootPath = runtime.GOROOT()

func GetMainPath() string {
	for i := 0; i < 1000; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			fmt.Printf("pc:%v,file:%v,line:%v,ok:%v\n", pc, file, line, ok)
			//忽略go自带的
			if goRootPath != "" && len(file) >= len(goRootPath) && goRootPath == file[0:len(goRootPath)] {
				continue
			}

			//忽略不是go的文件
			if path.Ext(file) != ".go" {
				continue
			}

		} else {
			break
		}
	}
	return ""
}
