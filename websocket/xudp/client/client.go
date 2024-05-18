package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// TODO 可以水文章
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
		if strings.ToUpper(strings.Trim(readString, "\r\n")) == "Q" {
			fmt.Println("退出udp客户端")
			return
		}
		udp.Write([]byte(readString))
		buf := make([]byte, 1024)
		n, addr, _ := udp.ReadFromUDP(buf)
		fmt.Printf("从服务端:%s,收到:%s\r\n", addr.String(), string(buf[:n]))
	}
}
