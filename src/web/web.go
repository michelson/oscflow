package web

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Example() {
	//http.HandleFunc("/", handler)
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8181", nil)
}
