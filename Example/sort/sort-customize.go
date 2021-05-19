package main

import "sort"

// ByLength https://www.yiibai.com/go/golang-sorting-by-functions.html#article-start
// 为了按Go中的自定义函数排序，我们需要一个
// 对应的类型。在这里，我们创建了一个“ ByLength”
// 类型，它只是内置`[] string`的别名
type ByLength []string

func main() {
	fruits := []string{"apple", "banana", "peach", "kiwi"}
	sort.Sort(ByLength(fruits))
}
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
