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

type Header struct {
	id           uint16
	header_flags HeaderFlags
	qd_count     uint16
	an_count     uint16
	ns_count     uint16
	ar_count     uint16
}

func NewHeader(id uint16, header_flags HeaderFlags, qd_count, an_count, ns_count, ar_count uint16) *Header {
	return &Header{id, header_flags, qd_count, an_count, ns_count, ar_count}
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

type Question struct {
	q_name  string
	q_type  uint16
	q_class uint16
}

type Record struct {
	name      string
	q_type    uint16
	q_class   uint16
	ttl       uint32
	rd_length uint16
	r_data    string
}

func NewRecord(name string, q_type, q_class uint16, ttl uint32, rd_length uint16, r_data string) Record {
	return Record{name, q_type, q_class, ttl, rd_length, r_data}
}

type Records []Record

func ParseMessage(req []byte) Message {
	pos := 0

	header, count := parseHeader(req, pos)
	pos += count

	question, count := parseQuestion(req, pos)
	pos += count

	answer, count := parseRecords(req, pos, int(header.an_count))
	pos += count

	authority, count := parseRecords(req, pos, int(header.ns_count))
	pos += count

	additional, count := parseRecords(req, pos, int(header.ar_count))
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
	q_name, count := parseDomainName(req, pos)
	pos += count

	q_type := parseUint16(req, pos)
	pos += 2

	q_class := parseUint16(req, pos)
	pos += 2

	ttl := parseUint32(req, pos)
	pos += 4

	rd_length := parseUint16(req, pos)
	pos += 2

	r_data_slice := make([]byte, rd_length)
	copy(r_data_slice, req[start:start+int(rd_length)])
	r_data := string(r_data_slice)

	pos += int(rd_length)

	record := Record{q_name, q_type, q_class, ttl, rd_length, r_data}

	return record, pos - start
}

func parseQuestion(req []byte, pos int) (Question, int) {
	q_name, count := parseDomainName(req, pos)
	pos += count

	q_type := parseUint16(req, pos)
	pos += 2

	q_class := parseUint16(req, pos)
	pos += 2

	question := Question{q_name, q_type, q_class}
	return question, count + 4
}

// todo: should handle both literal and pointer
func parseDomainName(req []byte, pos int) (string, int) {

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

	q_name := string(domain_slice)
	parse_count := pos - start
	return q_name, parse_count
}

func parseHeader(req []byte, pos int) (Header, int) {
	id := parseUint16(req, pos)
	pos += 2

	header_flags := parseHeaderFlags(req, pos)
	pos += 2

	qd_count := parseUint16(req, pos)
	pos += 2

	an_count := parseUint16(req, pos)
	pos += 2

	ns_count := parseUint16(req, pos)
	pos += 2

	ar_count := parseUint16(req, pos)

	header := Header{id, header_flags, qd_count, an_count, ns_count, ar_count}
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
