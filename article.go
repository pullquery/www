package main

import (
	"net/http"
	"path/filepath"
)

func HandleArticle(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/article/"):]
	path := filepath.Join("article/", name)

	if name != "" {
		HandleTemplate(res, req)
		return
	}

	res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(res, req, path)
}
