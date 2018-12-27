// simple app to listen on a port and display content
package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Lateef Koshemani"))
}

func main() {
	// handle web requests
	http.HandleFunc("/", hello)
	// startup webserver
	http.ListenAndServe(":8000", nil)
}
