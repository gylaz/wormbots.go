package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gylaz/wormbots.go/farm"
)

func streamData(f *farm.Farm) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-type", "text/event-stream")
		res.Header().Set("Cache-Control", "no-cache")
		res.Header().Set("Connection", "keep-alive")

		for {
			fmt.Fprintf(res, "data: %v\n\n", data(f))
			res.(http.Flusher).Flush()
			time.Sleep(time.Second)
			f.Tick()
		}
	}
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	f := farm.New()

	mux.Handle("/", fs)
	mux.HandleFunc("/world", streamData(&f))

	http.ListenAndServe(":3000", mux)
}

func data(f *farm.Farm) string {
	data, _ := json.Marshal(f.Worms)

	return string(data)
}
