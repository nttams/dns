package server

import (
	"client"
	"github.com/k0kubun/pp"
	msg "message"
	"net"
)

func Listen() {
	client.Init()

	// addr, err := net.ResolveUDPAddr("udp", "192.168.1.4:15353")
	addr, err := net.ResolveUDPAddr("udp", ":15353")
	if err != nil {
		panic("cannot connect")
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic("cannot connect")
	}

	encodedRequest := make([]byte, 512)
	for {
		_, remoteAddr, _ := conn.ReadFromUDP(encodedRequest)

		request := msg.ParseMessage(encodedRequest)

		answers := client.QueryARecordApi(request.GetQuestionDomain())
		pp.Println(answers)

		response := msg.CreateResponseFromRequest(request)
		response.SetAnswers(answers)

		conn.WriteToUDP(response.Encode(), remoteAddr)
	}
}
