package main

import (
	"github.com/quic-go/quic-go/http3"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("go http3"))
	})
	log.Fatal(http3.ListenAndServeQUIC("localhost:9501", "./server.crt", "./server.key", nil))
}
