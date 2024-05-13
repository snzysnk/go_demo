package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8088,
	})
	defer udp.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)
		n, addr, err := udp.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("接收消息出错:", err)
		}
		fmt.Printf("从地址:%s,收到:%s", addr, buf[:n])
		udp.WriteToUDP([]byte("我收到了"), addr)
	}
}
