package main

import (
	"github.com/stealthrocket/net/wasip1"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/stealthrocket/net/http"
)

var templateDirs = []string{"views"}
var templates *template.Template

func getTemplates() (templates *template.Template, err error) {
	var allFiles []string
	for _, dir := range templateDirs {
		files2, _ := os.ReadDir(dir)
		for _, file := range files2 {
			filename := file.Name()
			if strings.HasSuffix(filename, ".gohtml") {
				filePath := filepath.Join(dir, filename)
				allFiles = append(allFiles, filePath)
			}
		}
	}

	templates, err = template.New("").ParseFiles(allFiles...)
	return
}

func rootHandler(wr http.ResponseWriter, req *http.Request) {
	title := "home"

	data := map[string]interface{}{
		"title": title,
	}

	err := templates.ExecuteTemplate(wr, "index", data)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	l, err := wasip1.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", rootHandler)

	server := &http.Server{}
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}

func init() {
	var err error
	templates, err = getTemplates()
	if err != nil {
		panic(err)
	}
}
