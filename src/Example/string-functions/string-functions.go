package main

// https://www.yiibai.com/go/golang-string-functions.html#article-start
// 字符串函数

import (
	print "fmt"
	s "strings"
)

var p = print.Println

func main() {

	// string函数
	str := "hEllO world,hello go!"
	_, err := p("Contains:", s.Contains(str, "hello"))
	if err != nil {
		return
	}
	_, err1 := p("Count:", s.Count(str, "h"))
	if err1 != nil {
		return
	}
	_, err2 := p("ContainsAny:", s.ContainsAny(str, "h"))
	if err2 != nil {
		return
	}
	_, err3 := p("ContainsRune:", s.ContainsRune(str, 100))
	if err3 != nil {
		return
	}
	_, err4 := p("Compare:", s.Compare("hello:", "world"))
	if err4 != nil {
		return
	}
	_, err5 := p("Fields:", s.Fields(str))
	if err5 != nil {
		return
	}
	_, err6 := p("HasPrefix:", s.HasPrefix(str, "h"))
	if err6 != nil {
		return
	}
	_, err7 := p("HasSuffix:", s.HasSuffix(str, "!"))
	if err7 != nil {
		return
	}
	_, err8 := p("Index:", s.Index(str, "d"))
	if err8 != nil {
		return
	}
	_, err9 := p("IndexAny:", s.IndexAny(str, "d"))
	if err9 != nil {
		return
	}
	_, err10 := p("Join:", s.Join([]string{"ab", "cd"}, "|"))
	if err10 != nil {
		return
	}
	_, err11 := p("LastIndex:", s.LastIndex(str, "o"))
	if err11 != nil {
		return
	}
	_, err12 := p("ToLower:", s.ToLower(str))
	if err12 != nil {
		return
	}
	_, err13 := p("ToUpper:", s.ToUpper(str))
	if err13 != nil {
		return
	}
	_, err14 := p("ReplaceAll:", s.ReplaceAll(str, "l:", "&"))
	if err14 != nil {
		return
	}
	_, err15 := p("ReplaceAll:", s.Replace(str, "l:", "/:", 1))
	if err15 != nil {
		return
	}
}
