package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9502")
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := listen.Accept()
		if err != nil {
			con.Close()
			fmt.Print("connect failed")
		}

		go do(con)
	}
}

func do(con net.Conn) {
	defer con.Close()
	for {
		buf := make([]byte, 1024)
		_, err := con.Read(buf)

		fmt.Printf("客户端说:%s", buf)
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("服务端说:%s", readString)
		if _, err = con.Write([]byte(readString)); err != nil {
			fmt.Println(err)
			break
		}
	}
}
