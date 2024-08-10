package api

type User struct {
	// 	名字
	// Required: true
	Name string `json:"name"`
	// 年龄
	Age int `json:"age"`
	// 喜好
	Like string `json:"like"`
}
