package file

import "os"

// Exists 文件是否存在
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// IsDir 是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Pwd 获取当前路径
func Pwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}
