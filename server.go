package main

import (
	"github.com/fachriza/galaxy/src/trading"
	"net/http"
)

func main() {
	http.Handle("/convert", http.HandlerFunc(trading.PostTrading))
	static := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", static))
	http.ListenAndServe(":8089", nil)
}
