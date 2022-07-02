package main

import (
	"fmt"
	msg "message"
	"net"
	zone "zone_handler"
)

func main() {
	listen()
}

func listen() {
	addr, _ := net.ResolveUDPAddr("udp", "192.168.1.99:15353")
	conn, _ := net.ListenUDP("udp", addr)

	req := make([]byte, 512)
	for {

		_, remoteAddr, _ := conn.ReadFromUDP(req)

		result := msg.ParseMessage(req)
		domain := result.GetQuestionDomain()

		resultRecords := zone.FindDomain(domain)

		// fmt.Println(domain)
		fmt.Println("result: ", resultRecords)

		// ////////////////////////
		// res := make([]byte, 124)
		res := []byte{0x3b, 0x0c, 0x81, 0x80, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0x01, 0x00, 0x01, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x64, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x71, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x8b, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x8a, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x65, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x66}

		res[0] = req[0]
		res[1] = req[1]

		conn.WriteToUDP(res, remoteAddr)
	}
}
