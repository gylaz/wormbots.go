package main

import "net/http"

func serveWorld(w http.ResponseWriter, r *http.Request) {
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.HandleFunc("/world", serveWorld)
	mux.Handle("/", fs)

	http.ListenAndServe(":3000", mux)
}
