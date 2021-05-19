package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写文件
func main() {
	//首先，这是转储字符串的方法（或仅转储
	//字节）放入文件中。
	b1 := []byte("hello world \nhello go!")
	error := ioutil.WriteFile("data1.txt", b1, 0644)
	checkError(error)

	file, errorA := os.Create("data2.txt")
	// 打开文件后，推迟关闭 defer
	defer file.Close()
	checkError(errorA)

	// 使用字节切片写入文件
	data1 := []byte{115, 111, 109, 101, 10}
	// write不断写入，不会覆盖
	n1, error1 := file.Write(data1)
	fmt.Printf("wrote %d bytes\n", n1)
	checkError(error1)

	// 使用字节写入
	n2, error2 := file.Write([]byte("I love china ,i love world\n"))
	checkError(error2)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, error3 := file.WriteString("直接用字符串写入：hello world,hello 222\n")
	checkError(error3)
	fmt.Printf("wrote %d bytes\n", n3)

	// 发出“同步”以将写入刷新到稳定存储。
	file.Sync()

	write := bufio.NewWriter(file)
	n4, error4 := write.WriteString("使用NewWriter写入文件")
	checkError(error4)
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用`Flush`来确保所有缓冲操作都具有应用于writer
	// 如果不适用flush的话，无法写入文件 44行
	write.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
