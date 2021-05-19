package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 将文件内容读取到内存中
	data, err := ioutil.ReadFile("defer.txt")
	checkErr(err)
	// 字节数组转字符串
	fmt.Println(string(data))

	//您通常会希望对方式和内容进行更多控制
	//读取文件的一部分。对于这些任务，请开始
	//通过“打开”文件以获得“ os.File”值。
	file, err2 := os.Open("defer.txt")
	checkErr(err2)

	//从文件的开头读取一些字节。
	//最多允许读取5个，但还要注意有多少个
	//实际上已被读取。
	b1 := make([]byte, 12)
	count, err3 := file.Read(b1)
	checkErr(err3)
	fmt.Printf("%d bytes: %s\n", count, string(b1))

	//您也可以“搜索”到文件中的已知位置
	//并从此处读取。
	o2, err4 := file.Seek(6, 0)
	checkErr(err)
	b2 := make([]byte, 2)
	n2, err4 := file.Read(b2)
	checkErr(err4)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//`io`软件包提供了一些功能，这些功能可能
	//有助于读取文件。例如，读取
	//像上面的那些一样可以更强大
	//通过`ReadAtLeast`实现。
	o3, err := file.Seek(6, 0)
	checkErr(err)


	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, b3, 2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
