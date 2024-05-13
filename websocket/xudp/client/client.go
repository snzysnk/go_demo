package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8088,
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		readString, _ := reader.ReadString('\n')
		if strings.ToUpper(strings.Trim("readString", "\r\n")) == "Q" {
			fmt.Println("退出udp客户端")
			return
		}
		udp.Write([]byte(readString))
		buf := make([]byte, 1024)
		n, addr, _ := udp.ReadFromUDP(buf)
		fmt.Println(addr.String())
		fmt.Println(udp.RemoteAddr().String())
		fmt.Println("客户端收到:", string(buf[:n]))
	}
}
