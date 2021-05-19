package main

import (
	"fmt"
	"sort"
)

func main() {

	// 字符串排序
	strs := []string{"b", "a", "d", "c", "e", "f"}
	fmt.Println("Unsorted strings:", strs)
	// 字符串是否排序
	fmt.Println("Sorted:", sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)

	// int类型排序
	ints := []int{1, 3, 412, 2, 4, 5, 122}
	fmt.Println("Unsorted ints:", ints)
	// int类型是否排序
	fmt.Println("Sorted:", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	//
	issorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", issorted)

}
