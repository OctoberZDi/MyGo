package main

import (
	"crypto/sha1"
	"fmt"
)

//https://www.yiibai.com/go/golang-sha1-hashes.html#article-start
// SHA1散列经常用于计算二进制或文本块的短标识。
// crypto/*包中实现了几个哈希函数
func main() {

	s1 := "sha1 this string"
	sha1 := sha1.New()
	sha1.Write([]byte(s1))

	bs := sha1.Sum(nil)
	fmt.Println(s1)
	fmt.Printf("%x\n", bs)
}
