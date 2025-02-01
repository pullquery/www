package main

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func HandleTemplate(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/article/"):]
	path := filepath.Join("article/", name)

	article, err := executeTemplate(path)

	if err != nil {
		http.Error(
			res,
			http.StatusText(http.StatusNotFound),
			http.StatusNotFound,
		)

		return
	}

	res.Write(article)
}

func executeTemplate(path string) ([]byte, error) {
	html := template.Must(template.ParseFiles("template.html"))

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// handle LF or CRLF by system type
	splited := strings.SplitN(string(content), "\r\n\r\n", 2)

	structed := struct {
		Head string
		Body string
	}{
		Head: splited[0],
		Body: splited[1],
	}

	var buffer bytes.Buffer
	html.Execute(&buffer, structed)

	return buffer.Bytes(), nil
}
