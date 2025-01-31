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

	fmt.Println("Start listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
