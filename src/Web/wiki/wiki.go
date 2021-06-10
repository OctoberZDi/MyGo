package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	/*	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
		p1.save()
		p2, _ := loadPage("test")
		fmt.Println(string(p2.Body))*/

	// http.HandleFunc("/", handler)

	// http.HandleFunc() 注册一个处理器函数handler和对应的pattern
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// http.ListenAndServer 监听TCP地址，并且会使用handler参数调用server函数处理接收到的链接。
	// handler参数一般会设置为nil，此时会使用DefaultServerMux
	// 请求地址：http://localhost:8080/view/test
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 闭包
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler fn
		m := validPath.FindStringSubmatch(request.URL.Path)
		if m == nil {
			http.NotFound(writer, request)
			return
		}
		fn(writer, request, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// loads the page (or, if it doesn't exist, create an empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// The function template.ParseFiles will read the contents of edit.html
	// and return a *template.Template.
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("./src/Web/wiki/edit.html", "./src/Web/wiki/view.html"))

//var templates = template.Must(template.ParseFiles("./web/wiki/edit.html", "./web/wiki/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// tmpl模板文件名，后面需要添加.html
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//validPath表达式来验证路径并提取页面标题的函数：
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}

	//The title is the second subexpression.
	return m[2], nil
}
