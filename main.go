package main

import (
	"fmt"
	"net/http"
)

func Log() {

}

func main() {
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/article/", HandleArticle)
	http.HandleFunc("/media/", HandleMedia)

	fmt.Println("Start listening on port 80")
	http.ListenAndServe(":80", nil)
}
