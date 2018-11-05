package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>Go server is <strong>up</strong> and <strong>running</strong>..!!</div>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`<div>
		<h3>Welcome to Go Programming</h3>
		</div>`)
}
