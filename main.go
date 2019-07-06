package main

import (
	"github.com/hikaru7719/bookshelf-bot/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handle)
	http.ListenAndServe(":8080", nil)
}
