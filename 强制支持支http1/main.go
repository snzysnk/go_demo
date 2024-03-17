package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":9001"

	// 禁用 HTTP/2，将 TLSNextProto 设置为空映射
	srv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("go http1"))
	})

	// 假设你已经有了一个有效的证书和私钥文件
	log.Fatal(srv.ListenAndServeTLS("./server.crt", "./server.key"))
}
