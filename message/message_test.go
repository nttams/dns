package message

import (
	"testing"
)

var REQ = []byte{
	0x00, 0x05, // id
	0x01, 0x20, // header flags
	0x00, 0x01, // qd count
	0x00, 0x00, // an count
	0x00, 0x00, // ns count
	0x00, 0x01, // ar count
	0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00, //q_name
	0x00, 0x01, // qType
	0x00, 0x01, // qClass
	// OPT record
	0x00,       // root domain, for OPT record
	0x00, 0x29, // TYPE
	0x04, 0xd0, // CLASS (UDP payload size in OPT record)
	0x00, 0x00, 0x00, 0x00, // TTL (extended rcode and flags)
	0x00, 0x0c, // RD_LENGTH
	0x00, 0x0a, 0x00, 0x08, 0x54, 0xe6, 0x02, 0x6b, 0x32, 0xc0, 0x4b, 0x93,
}

func TestParseHeader(t *testing.T) {
	pos := 0
	header, count := parseHeader(REQ, pos)

	t.Logf("%+v\n", header)
	t.Log(count)

	t.Fatalf("failed")
}

func TestParseRecords(t *testing.T) {
	pos := 0
	records, count := parseRecords(REQ, pos, 1)

	t.Logf("%+v\n", records)
	t.Log(count)

	t.Fatalf("failed")
}

func TestParseQuestion(t *testing.T) {
	pos := 12
	question, count := parseQuestion(REQ, pos)

	t.Logf("%+v\n", question)
	t.Log(count)

	t.Fatalf("failed")
}
