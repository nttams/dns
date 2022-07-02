package message

import (
	// "fmt"
)

type Message struct {
	header     Header
	question   Question
	answer     Records
	authority  Records
	additional Records
}

type HeaderFlags struct {
	qr      bool
	op_code byte
	aa      bool
	tc      bool
	rd      bool
	ra      bool
	z       bool
	ad      bool
	cd      bool
	r_code  byte
}

type Header struct {
	id           uint16
	header_flags HeaderFlags
	qdCount      uint16
	anCount      uint16
	nsCount      uint16
	arCount      uint16
}

func NewHeader(id uint16, header_flags HeaderFlags, qdCount, anCount, nsCount, arCount uint16) *Header {
	return &Header{id, header_flags, qdCount, anCount, nsCount, arCount}
}

type domainType uint8

const (
	InvalidType domainType = iota
	Literal
	Pointer
)

type Domain struct {
	domainType    domainType
	domainLiteral string
	point         uint16
}

type Question struct {
	domain string
	qType  uint16
	qClass uint16
}

type Record struct {
	domain   string
	qType    uint16
	qClass   uint16
	ttl      uint32
	rdLength uint16
	rData    string
}

func (record *Record) GetDomain() string {
	return record.domain
}

func NewRecord(domain string, qType, qClass uint16, ttl uint32, rdLength uint16, rData string) Record {
	return Record{domain, qType, qClass, ttl, rdLength, rData}
}

type Records []Record

func (msg *Message) GetQuestionDomain() string {
	return msg.question.domain
}

func (msg *Message) GetId() uint16 {
	return msg.header.id
}

// encoding

func (hf *HeaderFlags) encode() []byte {

	var first byte = 0b0000_0000
	var second byte = 0b0000_0000

	if hf.qr {
		first = first | (1 << 7)
	}
	first = first | (hf.op_code << 3)
	if hf.aa {
		first = first | (1 << 2)
	}
	if hf.tc {
		first = first | (1 << 1)
	}
	if hf.rd {
		first = first | (1 << 0)
	}

	if hf.ra {
		second = second | (1 << 7)
	}
	if hf.z {
		second = second | (1 << 7)
	}
	if hf.ad {
		second = second | (1 << 7)
	}
	if hf.cd {
		second = second | (1 << 7)
	}
	second = second | (hf.r_code << 0)

	return []byte{first, second}
}

func (header *Header) encode() []byte {
	result := []byte{}

	result = append(result, byte((header.id&0b1111_1111_0000_0000)>>8))
	result = append(result, byte((header.id&0b0000_0000_1111_1111)>>0))

	result = append(result, header.header_flags.encode()...) //todo: learn this

	result = append(result, byte((header.qdCount&0b1111_1111_0000_0000)>>8))
	result = append(result, byte((header.qdCount&0b0000_0000_1111_1111)>>0))

	result = append(result, byte((header.anCount&0b1111_1111_0000_0000)>>8))
	result = append(result, byte((header.anCount&0b0000_0000_1111_1111)>>0))

	result = append(result, byte((header.nsCount&0b1111_1111_0000_0000)>>8))
	result = append(result, byte((header.nsCount&0b0000_0000_1111_1111)>>0))

	result = append(result, byte((header.arCount&0b1111_1111_0000_0000)>>8))
	result = append(result, byte((header.arCount&0b0000_0000_1111_1111)>>0))

	return result
}

func (question *Question) encode() []byte {
	result := []byte{}
	return result
}

// parsing

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

func parseRecords(req []byte, pos, record_count int) (Records, int) {
	records := []Record{}
	start := pos

	for i := 0; i < record_count; i++ {
		record, count := parseRecord(req, pos)
		pos += count

		records = append(records, record)
	}
	return records, pos - start
}

func parseRecord(req []byte, pos int) (Record, int) {
	start := pos
	domain, count := parsedomain(req, pos)
	pos += count

	qType := parseUint16(req, pos)
	pos += 2

	qClass := parseUint16(req, pos)
	pos += 2

	ttl := parseUint32(req, pos)
	pos += 4

	rdLength := parseUint16(req, pos)
	pos += 2

	rData_slice := make([]byte, rdLength)
	copy(rData_slice, req[pos:pos+int(rdLength)])
	rData := string(rData_slice)

	pos += int(rdLength)

	record := Record{domain, qType, qClass, ttl, rdLength, rData}

	return record, pos - start
}

func parseQuestion(req []byte, pos int) (Question, int) {
	domain, count := parsedomain(req, pos)
	pos += count

	qType := parseUint16(req, pos)
	pos += 2

	qClass := parseUint16(req, pos)
	pos += 2

	question := Question{domain, qType, qClass}
	return question, count + 4
}

// todo: should handle both literal and pointer
func parsedomain(req []byte, pos int) (string, int) {

	// todo: quick fix
	if req[pos] == 0 {
		root := []byte{0}
		return string(root), 1
	}

	start := pos
	for req[pos] != 0 {
		pos += int(req[pos] + 1)
	}
	pos += 1

	domain_slice := make([]byte, pos-start-2) // leave first and last .
	copy(domain_slice, req[start+1:pos-1])

	length := len(domain_slice)
	for i := 0; i < length; i++ {
		if domain_slice[i] < 32 {
			domain_slice[i] = 46 //dot
		}
	}

	domain := string(domain_slice)
	parse_count := pos - start
	return domain, parse_count
}

func parseHeader(req []byte, pos int) (Header, int) {
	id := parseUint16(req, pos)
	pos += 2

	header_flags := parseHeaderFlags(req, pos)
	pos += 2

	qdCount := parseUint16(req, pos)
	pos += 2

	anCount := parseUint16(req, pos)
	pos += 2

	nsCount := parseUint16(req, pos)
	pos += 2

	arCount := parseUint16(req, pos)

	header := Header{id, header_flags, qdCount, anCount, nsCount, arCount}
	return header, 12
}

func parseHeaderFlags(req []byte, pos int) HeaderFlags {
	qr := req[pos]&0b1000_0000 == 1
	op_code := req[pos] & 0b0111_1000
	aa := req[pos]&0b0000_0100 == 1
	tc := req[pos]&0b0000_0010 == 1
	rd := req[pos]&0b0000_0001 == 1

	ra := req[pos+1]&0b1000_0000 == 1
	z := req[pos+1]&0b0100_0000 == 1
	ad := req[pos+1]&0b0010_0000 == 1
	cd := req[pos+1]&0b0001_0000 == 1
	r_code := req[pos+1] & 0b0000_1111

	reqult := HeaderFlags{qr, op_code, aa, tc, rd, ra, z, ad, cd, r_code}

	return reqult
}

func parseUint16(req []byte, pos int) uint16 {
	return uint16(req[pos])<<8 | uint16(req[pos+1])
}

func parseUint32(req []byte, pos int) uint32 {
	return uint32(req[pos+0])<<24 |
		uint32(req[pos+1])<<16 |
		uint32(req[pos+2])<<8 |
		uint32(req[pos+3])
}
