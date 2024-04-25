package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listen, err := net.Dial("tcp", "127.0.0.1:9502")
	defer listen.Close()
	if err != nil {
		log.Fatal(err)
	}

	do(listen)
}

func do(con net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("客户端说:%s", readString)

		_, err = con.Write([]byte(readString))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 1024)
		_, err = con.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("服务端说:%s", buf)

	}
}
