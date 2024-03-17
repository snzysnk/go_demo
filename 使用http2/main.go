package main

import (
	"log"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":9001"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("go http2"))
	})
	log.Fatal(srv.ListenAndServeTLS("./server.crt", "./server.key"))
}
