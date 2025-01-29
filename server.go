package main

import "main/message"

type DnsServer struct {
	DnsClient *Client
	UdpCtrl   *UdpCtrl
}

func NewDnsServer() *DnsServer {
	c := NewClient()
	uc := NewUdpCtrl()

	return &DnsServer{
		DnsClient: &c,
		UdpCtrl:   &uc,
	}
}
func (s *DnsServer) Listen() {
	s.UdpCtrl.Listen(":15353")

	for {
		encodedRequest, remoteAddr := s.UdpCtrl.Read()

		request := message.ParseMessage(encodedRequest)

		answers := s.DnsClient.QueryARecordApi(request.GetQuestionDomain())

		response := message.CreateResponseFromRequest(request)
		response.SetAnswers(answers)

		s.UdpCtrl.Write(response.Encode(), remoteAddr)
	}
}
