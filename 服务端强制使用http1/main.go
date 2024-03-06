package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:         ":9501",
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)), //只要不是nil,代表禁止alpn自动升级协议至http2
	}
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("http1.1"))
	})

	log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}
