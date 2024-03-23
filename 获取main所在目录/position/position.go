package position

import (
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

var (
	goRootPath       = runtime.GOROOT()
	matchPackageMain = regexp.MustCompile(`^package\s+main\s+`)
)

func GetMainDir() string {
	for i := 0; i < 1000; i++ {
		if pc, file, _, ok := runtime.Caller(i); ok {
			//根据go root path 忽略go自带的
			if goRootPath != "" && len(file) >= len(goRootPath) && goRootPath == file[0:len(goRootPath)] {
				continue
			}

			//忽略不是go的文件
			if path.Ext(file) != ".go" {
				continue
			}

			//还需要有main方法，才能认定为主包
			if fn := runtime.FuncForPC(pc); fn != nil {
				array := strings.Split(fn.Name(), ".")
				if array[0] != "main" {
					continue
				}
			}

			//是否已 package main 开头
			if matchPackageMain.Match(getByteContents(file)) {
				return path.Dir(file)
			}

		} else {
			break
		}
	}
	return ""
}

func getStringContents(filePath string) string {
	return string(getByteContents(filePath))
}

func getByteContents(filePath string) []byte {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}
	return fileContents
}
