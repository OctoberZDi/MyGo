package main

import (
	"fmt"
	"net/url"
)

// Go URL解析实例
// URL提供了一种统一的方法来定位资源。
func main() {

	urlStr := "postgres://user:pass@host.com:5432/path?name=zhangdi&age=20#f"
	u, error := url.Parse(urlStr)
	if error != nil {
		panic(error)
	}

	fmt.Println("url=", u)
	host := u.Host
	fmt.Println("host=", host)
	path := u.Path
	fmt.Println("path=", path)
	rawPath := u.RawPath
	fmt.Println("rawPath=", rawPath)
	schema := u.Scheme
	fmt.Println("schema=", schema)
	forceQuery := u.ForceQuery
	fmt.Println("forceQuery=", forceQuery)
	fragment := u.Fragment
	fmt.Println("fragment=", fragment)
	user := u.User
	fmt.Println("user=", user)
	fmt.Println("user.UserName=", user.Username())
	password, _ := user.Password()
	fmt.Println("user.Password=", password)

	// values 是个map[string][]string->key:string;value:[]string
	values, error := url.ParseQuery(u.RawQuery)
	if error != nil {
		panic(error)
	}
	for s, strings := range values {
		fmt.Println("index=", s, "value=", strings)
	}

}
