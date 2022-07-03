package client

import (
	// "fmt"
	"net"
	"time"
	"bufio"
	"message"
	"math/rand"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

// todo: add NetworkController to handle TCP/UDP
func Query(domain string) []string {
	query := message.NewQuery(generateUniqueId(), domain)
	encodedRequest := query.Encode()

	conn, _ := net.Dial("udp", "8.8.8.8:53")

	conn.Write(encodedRequest)

	buffer := make([]byte, 1024)
	count, _ := bufio.NewReader(conn).Read(buffer)
	encodedResponse := buffer[:count]
	response := message.ParseMessage(encodedResponse)

	return response.GetAnswers()
}

// todo: make it unique. For now, it's just random
// client needs to use something like map to store ids from multiple queries
func generateUniqueId() uint16{
	return uint16(rand.Intn(65536))
}