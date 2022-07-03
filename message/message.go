package message

import (
	// "fmt"
	"strconv"
)

type Message struct {
	header     Header
	question   Question
	answer     Records
	authority  Records
	additional Records
}

// this will expose internal data
func (msg *Message) GetAnswers() (result []string) {
	for _, answer := range msg.answer {
		ip := strconv.Itoa(int(answer.rData[0])) + "." +
			  strconv.Itoa(int(answer.rData[1])) + "." +
			  strconv.Itoa(int(answer.rData[2])) + "." +
			  strconv.Itoa(int(answer.rData[3]))
		result = append(result, ip)
		// result = append(result, string(answer.rData))
	}
	return
}

func NewQuery(id uint16, domain string) (msg Message) {
	msg.header = newRequestHeader(id)
	msg.question = NewQuestion(domain)
	msg.header.qdCount = 1

	return
}

func (msg *Message) Encode() (result []byte) {
	result = append(result, msg.header.encode()...)
	result = append(result, msg.question.encode()...)
	result = append(result, msg.answer.encode()...)
	result = append(result, msg.authority.encode()...)
	result = append(result, msg.additional.encode()...)
	return
}

func (msg *Message) GetQuestionDomain() string {
	return msg.question.domain.domainLiteral
}

func (msg *Message) GetId() uint16 {
	return msg.header.id
}

func ParseMessage(req []byte) Message {
	pos := 0

	header, count := parseHeader(req, pos)
	pos += count

	question, count := parseQuestion(req, pos)
	pos += count

	answer, count := parseRecords(req, pos, int(header.anCount))
	pos += count

	authority, count := parseRecords(req, pos, int(header.nsCount))
	pos += count

	additional, count := parseRecords(req, pos, int(header.arCount))
	pos += count

	return Message{header, question, answer, authority, additional}
}

func encodeUint16(value uint16) (result []byte) {
	result = append(result, byte((value & 0b11111111_00000000) >> 8))
	result = append(result, byte((value & 0b00000000_11111111)))
	return
}

func encodeUint32(value uint32) (result []byte) {
	result = append(result, byte((value & 0b11111111_00000000_00000000_00000000) >> 24))
	result = append(result, byte((value & 0b00000000_11111111_00000000_00000000) >> 16))
	result = append(result, byte((value & 0b00000000_00000000_11111111_00000000) >> 8))
	result = append(result, byte((value & 0b00000000_00000000_00000000_11111111)))
	return
}

func parseUint16(req []byte, pos int) uint16 {
	return uint16(req[pos])<<8 | uint16(req[pos+1])
}

func parseUint32(req []byte, pos int) uint32 {
	return uint32(req[pos+0]) << 24 |
		   uint32(req[pos+1]) << 16 |
		   uint32(req[pos+2]) <<  8 |
		   uint32(req[pos+3])
}