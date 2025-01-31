package main

import (
	"net/http"
	"path/filepath"
)

func HandleMedia(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/media/"):]
	path := filepath.Join("media/", name)

	res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(res, req, path)
}
