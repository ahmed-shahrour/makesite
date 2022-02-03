package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	fileContents, _ := ioutil.ReadFile("first-post.txt")

	page := Page{
		TextFilePath: "./",
		TextFileName: "first-post",
		HTMLPagePath: "first-post.html",
		Content:      string(fileContents),
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile, _ := os.Create("first-post.html")
	t.Execute(newFile, page)
}
