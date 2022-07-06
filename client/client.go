package client

import (
	"time"
	"math/rand"
	msg "message"
	nc "network_controller"
)

type Client struct {

}

func NewClient() Client {
	rand.Seed(time.Now().UnixNano())
	return Client {}
}

func (client *Client) Query(domain string) []string {
	return msg.ConvertRecordsToStrings(query(domain, msg.Q_TYPE_A))
}

func (client *Client) QueryARecordApi(domain string) []msg.Record {
	return query(domain, msg.Q_TYPE_A)
}

func (client *Client) QueryAAAARecordApi(domain string, qType uint16) []msg.Record {
	return query(domain, qType)
}

func query(domain string, qType uint16) []msg.Record {
	udpCtrl := nc.NewUdpCtrl()
	query := msg.NewQuery(generateUniqueId(), domain, qType)

	encodedRequest := query.Encode()

	udpCtrl.Send(encodedRequest, "8.8.8.8:53")
	encodedResponse := udpCtrl.Receive()

	response := msg.ParseMessage(encodedResponse)

	return response.GetRawAnswers()
}

// todo: make it unique. For now, it's just random
// client needs to use something like map to store ids from multiple queries
func generateUniqueId() uint16 {
	return uint16(rand.Intn(65536))
}
