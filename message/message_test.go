package message

import (
	"testing"
)

// todo: add another request
var REQUEST_0 = []byte{
	0x00, 0x05, // id
	0x01, 0x20, // header flags
	0x00, 0x01, // qd count
	0x00, 0x00, // an count
	0x00, 0x00, // ns count
	0x00, 0x01, // ar count
	0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00,
	0x00, 0x01, // qType
	0x00, 0x01, // qClass
	// OPT record
	0x00,       // root domain, for OPT record
	0x00, 0x29, // TYPE
	0x10, 0x00, // CLASS (UDP payload size in OPT record)
	0x00, 0x00, 0x00, 0x00, // TTL (extended rcode and flags)
	0x00, 0x0c, // RD_LENGTH
	0x00, 0x0a, 0x00, 0x08, 0x54, 0xe6, 0x02, 0x6b, 0x32, 0xc0, 0x4b, 0x93,
}

func TestParseHeader(t *testing.T) {
	expectedHeaderFlags := HeaderFlags { false, 0, false, false, true, false, false, false, false, 0 }
	expectedHeader := Header {5, expectedHeaderFlags, 1, 0, 0, 1}
	expectedCount := 12

	pos := 0
	header, count := parseHeader(REQUEST_0, pos)

	if header != expectedHeader {
		t.Fatalf("parse content failed")
	}
	if count != expectedCount {
		t.Fatalf("parse count failed")
	}
}

func TestParseQuestion(t *testing.T) {
	expectedQuestion := Question {"google.com", 1, 1}
	expectedCount := 16

	pos := 12
	question, count := parseQuestion(REQUEST_0, pos)

	if question != expectedQuestion {
		t.Fatalf("parse content failed")
	}

	if count != expectedCount {
		t.Fatalf("parse content failed")
	}
}

func TestParseRecords(t *testing.T) {
	expectedRecord := Record {
		string(byte(0)), 41, 4096, 0, 12,
		string([]byte {0x00, 0x0a, 0x00, 0x08, 0x54, 0xe6, 0x02, 0x6b, 0x32, 0xc0, 0x4b, 0x93} ),
	}
	expectedRecords := Records { expectedRecord }
	expectedCount := 23

	pos := 28
	records, count := parseRecords(REQUEST_0, pos, 1)
	if len(records) != len(expectedRecords) {
		t.Fatalf("parse content failed")
	}

	for i, _ := range records {
		if records[i] != expectedRecords[i] {
			t.Fatalf("parse content failed")
		}
	}

	if count != expectedCount {
		t.Fatalf("parse content failed")
	}
}
