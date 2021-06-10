package main

import "time"

// Login "error": "Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"
// 类似于出现这种错误，是因为binding:"required"导致的，可以改为 binding:"-" 以此来跳过认证。
type Login struct {
	User     string `xml:"user" json:"user" form:"user" binding:"required"`
	Password string `json:"password" xml:"password" yaml:"password" form:"password" binding:"required"`
}

// Person Person
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type Person2 struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

// URI 测试uri绑定
type URI struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

//TestHeader 测试header绑定
type TestHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

type myForm struct {
	Colors []string `form:"colors[]"`
}
