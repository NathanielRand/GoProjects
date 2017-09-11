package main

import (
	"net/http"

	"goprojects/cms"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/post/", cms.ServePost)

	http.ListenAndServe(":8080", nil)
}
