package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func AddRootCA(certPool *x509.CertPool) {
	caCertPath := path.Join("./ca.crt")
	caCertRaw, err := os.ReadFile(caCertPath)
	if err != nil {
		panic(err)
	}
	if ok := certPool.AppendCertsFromPEM(caCertRaw); !ok {
		panic("Could not add root ceritificate to pool.")
	}
}

func main() {
	var (
		addr = "https://www.open1.com:9001"
		pool *x509.CertPool
		err  error
	)
	if pool, err = x509.SystemCertPool(); err != nil {
		log.Fatal(err)
	}

	AddRootCA(pool)
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
		},
		QuicConfig: &quic.Config{},
	}
	defer roundTripper.Close()
	h3Client := &http.Client{
		Transport: roundTripper,
	}

	rsp, err := h3Client.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	all, err := io.ReadAll(rsp.Body)
	fmt.Printf("%s\n", all)
}
