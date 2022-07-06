package server

import (
	"client"
	msg "message"
	nc "network_controller"
)

func Listen() {
	cli := client.NewClient()

	udpCtrl := nc.NewUdpCtrl()
	udpCtrl.Listen(":15353")

	for {
		encodedRequest, remoteAddr := udpCtrl.Read()

		request := msg.ParseMessage(encodedRequest)

		answers := cli.QueryARecordApi(request.GetQuestionDomain())

		response := msg.CreateResponseFromRequest(request)
		response.SetAnswers(answers)

		udpCtrl.Write(response.Encode(), remoteAddr)
	}
}
