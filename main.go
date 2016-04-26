package main

import (
	"fmt"
	"net/http"
	"time"
)

func streamData(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")

	for {
		fmt.Fprintf(res, "data: "+time.Now().Format(time.UnixDate)+"\n\n")
		res.(http.Flusher).Flush()
		time.Sleep(time.Second)
	}
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	mux.HandleFunc("/world", streamData)

	http.ListenAndServe(":3000", mux)
}
