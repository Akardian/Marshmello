package handler

import "net/http"

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<p>Hello from Go!</p>"))
}
