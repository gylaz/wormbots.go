package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gylaz/wormbots.go/farm"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	f := farm.New()

	go func() {
		for {
			time.Sleep(time.Second / 15)

			f.Tick()
		}
	}()

	mux.Handle("/", fs)
	mux.HandleFunc("/world", streamData(&f))

	http.ListenAndServe(":4000", mux)
}

func streamData(f *farm.Farm) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-type", "text/event-stream")
		res.Header().Set("Cache-Control", "no-cache")
		res.Header().Set("Connection", "keep-alive")

		for {
			worms, _ := json.Marshal(f.Worms)
			fmt.Fprintf(res, "data: %v\n\n", string(worms))
			res.(http.Flusher).Flush()
			time.Sleep(time.Second / 15)
		}
	}
}
