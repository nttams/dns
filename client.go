package dns

import (
	"bufio"
	"math/rand"
	"net"

	"dns/message"
)

type Client struct {
	addr string
}

func NewClient(addr string) Client {
	return Client{
		addr: addr,
	}
}

func (c *Client) Query(domain string) ([]string, error) {
	result, err := c.query(domain, message.Q_TYPE_A)
	if err != nil {
		return nil, err
	}
	return message.ConvertRecordsToStrings(result), nil
}

func (c *Client) query(domain string, qType uint16) ([]message.Record, error) {
	query := message.NewQuery(generateUniqueId(), domain, qType)
	conn, err := net.Dial("udp", c.addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.Write(query.Encode())
	buf := make([]byte, 2048)
	count, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		return nil, err
	}
	response := message.ParseMessage(buf[:count])
	return response.GetRawAnswers(), nil
}

// todo: make it unique. For now, it's just random
// client needs to use something like map to store ids from multiple queries
func generateUniqueId() uint16 {
	return uint16(rand.Intn(65536))
}
