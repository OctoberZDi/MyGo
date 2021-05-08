package main

import "fmt"

//for 初始化语句和后置语句是可选的
func main() {

	// 省略初始化语句和后置语句
	sum := 1
	for sum <= 100 {
		sum += sum
	}
	fmt.Println(sum)
	// 正常的for
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	str1 := "go is a beautiful language!"
	for i, val := range str1 {
		//fmt.Println()
		fmt.Printf("Character on position %d is: %c \n", i, val)
	}

	str1 = "Go is a beautiful language!"
	fmt.Printf("The length of str1 is: %d\\n", len(str1))
	for pos, char := range str1 {
		fmt.Printf("Character on position %d is: %c \\n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\\n", char, pos)
	}
	fmt.Println()
	/*	fmt.Println("index int(rune) rune    char bytes")
		for index, rune := range str2 {
			fmt.Printf("%-2d      %d      %U '%c' % X\\n", index, rune, rune, rune, []byte(string(rune)))
		}*/
}

func g1() {

}
