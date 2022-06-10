package main

import (
	"fmt"
	"net"
)

func main() {
	// testSlices()
	listen()
}

func testSlices() {
	arr := [5]int{1, 2, 3, 4, 5}

	sli := arr[0:3]

	arr[0] = 20

	fmt.Println(arr)
	fmt.Println(sli)
}

func listen() {
	addr, _ := net.ResolveUDPAddr("udp", "192.168.1.6:15353")

	conn, _ := net.ListenUDP("udp", addr)

	// fmt.Println("err: ", err)

	// defer conn.Close()

	req := make([]byte, 512)
	len_read, remote_addr, _ := conn.ReadFromUDP(req)


	// res := make([]byte, 124)
	res := []byte{0x3b, 0x0c, 0x81, 0x80, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0x01, 0x00, 0x01, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x64, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x71, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x8b, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x8a, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x65, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x17, 0x00, 0x04, 0x4a, 0x7d, 0x44, 0x66}


	res[0] = req[0]
	res[1] = req[1]

	conn.WriteToUDP(res, remote_addr)

	fmt.Println(addr, conn)
	fmt.Println("read: ", len_read)
	fmt.Println("remote: ", remote_addr)
	fmt.Println("request: ", req[:len_read])
	fmt.Println("response: ", res)
}

func test_client() {
	fmt.Println("dns server - start")

	req := make([]byte, 28)

	req[0] = 0x3b
	req[1] = 0x0c
	req[2] = 0x01
	req[3] = 0x20
	req[4] = 0x00
	req[5] = 0x01
	req[6] = 0x00
	req[7] = 0x00
	req[8] = 0x00
	req[9] = 0x00
	req[10] = 0x00
	req[11] = 0x00
	req[12] = 0x06
	req[13] = 0x67
	req[14] = 0x6f
	req[15] = 0x6f
	req[16] = 0x67
	req[17] = 0x6c
	req[18] = 0x65
	req[19] = 0x03
	req[20] = 0x63
	req[21] = 0x6f
	req[22] = 0x6d
	req[23] = 0x00
	req[24] = 0x00
	req[25] = 0x01
	req[26] = 0x00
	req[27] = 0x01
	fmt.Println("req: ", req)

	ip := net.IP{8, 8, 8, 8}
	address := net.UDPAddr{ip, 53, ""}

	conn, err := net.DialUDP("udp", nil, &address)

	fmt.Println("connection: ", conn)
	fmt.Println("error: ", err)

	// send
	_, err = conn.Write(req);

	// read
	res := make([]byte, 512)
	read, err := conn.Read(res)

	fmt.Println("read: ", read)
	fmt.Println("received: ", res)

	fmt.Println("dns server - end")
}