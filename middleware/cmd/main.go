package main

import (
	"net/http"

	"goprojects/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	logger := middleware.CreateLogger("section4")
	http.Handle("/", middleware.Time(logger, hello))
	http.ListenAndServe(":8080", nil)
}
