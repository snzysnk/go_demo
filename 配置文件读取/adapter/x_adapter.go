package adapter

import (
	"go_demo/安全的map或切片/mp"
	"go_demo/安全的map或切片/sl"
)

type XAdapterFile struct {
	searchName  string        //搜索文件名
	searchPaths sl.IXStrSlice //搜索路径
	jsonMap     mp.IXMap      //数据
}
