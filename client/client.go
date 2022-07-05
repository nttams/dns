package client

import (
	"bufio"
	"math/rand"
	msg "message"
	"net"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

// todo: add NetworkController to handle TCP/UDP
func Query(domain string) []string {
	// return msg.ConvertRecordsToStrings(query(domain, msg.Q_TYPE_A))
	return msg.ConvertRecordsToStrings(query(domain, msg.Q_TYPE_AAAA))
}

func QueryARecordApi(domain string) []msg.Record {
	return query(domain, msg.Q_TYPE_A)
}

func QueryAAAARecordApi(domain string, qType uint16) []msg.Record {
	return query(domain, qType)
}

func query(domain string, qType uint16) []msg.Record {
	query := msg.NewQuery(generateUniqueId(), domain, qType)
	encodedRequest := query.Encode()

	conn, _ := net.Dial("udp", "8.8.8.8:53")

	conn.Write(encodedRequest)

	buffer := make([]byte, 1024)
	count, _ := bufio.NewReader(conn).Read(buffer)
	encodedResponse := buffer[:count]
	response := msg.ParseMessage(encodedResponse)

	return response.GetRawAnswers()
}

// todo: make it unique. For now, it's just random
// client needs to use something like map to store ids from multiple queries
func generateUniqueId() uint16 {
	return uint16(rand.Intn(65536))
}
