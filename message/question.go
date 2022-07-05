package message

type Question struct {
	domain string
	qType  uint16
	qClass uint16
}

func NewQuestion(domain string, qType uint16) Question {
	return Question{domain, qType, DEFAULT_Q_CLASS}
}

// todo: now
func (question *Question) encode() (result []byte) {
	result = append(result, encodeDomain(question.domain)...)
	result = append(result, encodeUint16(question.qType)...)
	result = append(result, encodeUint16(question.qClass)...)
	return
}

// todo: use generic? (this is similar to encodeRecords)
func encodeQuestions(questions []Question) (result []byte) {
	for _, question := range questions {
		result = append(result, question.encode()...)
	}
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
	question := Question{domain, qType, qClass}
	return question, count + 4
}

func parseQuestions(req []byte, pos, questionCount int) ([]Question, int) {
	questions := []Question{}
	start := pos

	for i := 0; i < questionCount; i++ {
		question, count := parseQuestion(req, pos)
		pos += count

		questions = append(questions, question)
	}
	return questions, pos - start
}
