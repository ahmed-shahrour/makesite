package main

import (
	"flag"
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
	fileFlag := flag.String("file", "first-post.txt", "Enter the name of the file to be converted")
	flag.Parse()

	fileContents, _ := ioutil.ReadFile(*fileFlag)
	fileName := (*fileFlag)[:len(*fileFlag)-4]

	page := Page{
		TextFilePath: "./",
		TextFileName: fileName,
		HTMLPagePath: fileName + ".html",
		Content:      string(fileContents),
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile, _ := os.Create(page.HTMLPagePath)
	t.Execute(newFile, page)
}
