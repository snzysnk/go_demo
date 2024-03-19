package adapter

import (
	"read_file_config/mp"
	"read_file_config/sl"
)

type XAdapterFile struct {
	searchName  string        //搜索文件名
	searchPaths sl.IXStrSlice //搜索路径
	jsonMap     mp.IXMap      //数据
}
