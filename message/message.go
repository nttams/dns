package message

import (
	// "fmt"
	"strconv"
)

const Q_TYPE_A = 1
const Q_TYPE_AAAA = 28

const Q_CLASS_IN = 1
const DEFAULT_Q_CLASS = Q_CLASS_IN

type Message struct {
	header      Header
	questions   []Question
	answers     []Record
	authorities []Record
	additionals []Record
}

func NewQuery(id uint16, domain string, qType uint16) (msg Message) {
	msg.header = newRequestHeader(id)
	msg.questions = []Question{NewQuestion(domain, qType)}
	msg.header.qdCount = 1

	return
}

func CreateResponseFromRequest(req Message) (res Message) {
	res.header.id = req.header.id

	res.header.headerFlags = req.header.headerFlags
	res.header.headerFlags.qr = true
	res.header.headerFlags.ra = true

	res.header.qdCount = req.header.qdCount
	res.questions = req.questions

	return
}

func (res *Message) SetAnswers(answers []Record) {
	res.answers = answers
	res.header.anCount = uint16(len(answers))
}

func (msg *Message) Encode() (result []byte) {
	result = append(result, msg.header.encode()...)
	result = append(result, encodeQuestions(msg.questions)...)
	result = append(result, encodeRecords(msg.answers)...)
	result = append(result, encodeRecords(msg.authorities)...)
	result = append(result, encodeRecords(msg.additionals)...)

	return
}

// todo: note this [0]
func (msg *Message) GetQuestionDomain() string {
	return msg.questions[0].domain
}

func (msg *Message) GetId() uint16 {
	return msg.header.id
}

func ParseMessage(req []byte) Message {
	pos := 0

	header, count := parseHeader(req, pos)
	pos += count

	question, count := parseQuestions(req, pos, int(header.qdCount))
	pos += count

	answer, count := parseRecords(req, pos, int(header.anCount))
	pos += count

	authority, count := parseRecords(req, pos, int(header.nsCount))
	pos += count

	additional, count := parseRecords(req, pos, int(header.arCount))
	pos += count

	return Message{header, question, answer, authority, additional}
}

func (msg *Message) GetAnswers() (result []string) {
	for _, answer := range msg.answers {
		ip := strconv.Itoa(int(answer.rData[0])) + "." +
			strconv.Itoa(int(answer.rData[1])) + "." +
			strconv.Itoa(int(answer.rData[2])) + "." +
			strconv.Itoa(int(answer.rData[3]))
		result = append(result, ip)
	}
	return
}

func (msg *Message) GetRawAnswers() []Record {
	return msg.answers
}
