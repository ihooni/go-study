package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello route"))
}

func main() {
	http.HandleFunc("/", home)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
