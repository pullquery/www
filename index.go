package main

import (
	"net/http"
	"path/filepath"
)

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/"):]
	path := filepath.Join("index/", name)

	res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(res, req, path)
}
