package message

type Question struct {
	domain Domain
	qType  uint16
	qClass uint16
}

const DEFAULT_Q_TYPE = 1 // A
const DEFAULT_Q_CLASS = 1 // IN

func NewQuestion(domain string) Question {
	return Question { Domain {Literal, domain, 0}, DEFAULT_Q_TYPE, DEFAULT_Q_CLASS }
}

// todo: now
func (question *Question) encode() (result []byte) {
	result = append(result, question.domain.encode()...)
	result = append(result, encodeUint16(question.qType)...)
	result = append(result, encodeUint16(question.qClass)...)
	return
}

func parseQuestion(req []byte, pos int) (Question, int) {
	domain, count := parseDomain(req, pos)
	pos += count

	qType := parseUint16(req, pos)
	pos += 2

	qClass := parseUint16(req, pos)
	pos += 2

	// todo
	question := Question{Domain {Literal, domain, 0}, qType, qClass}
	return question, count + 4
}