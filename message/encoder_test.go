package message

import (
	"testing"
	"reflect"
)

func TestEncodeDomain(t *testing.T) {
	domain := Domain { Literal, "google.com", 0 }
	expectedEncodedDomain := []byte {6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0}
	encodedDomain := domain.encode()

	if !reflect.DeepEqual(encodedDomain, expectedEncodedDomain) {
		t.Fatal("parse failed")
	}
}

func TestEncodeHeaderFlags(t *testing.T) {
	hf := HeaderFlags {}
	hf.rd = true

	t.Log(hf)
	t.Log(hf.encode())

	t.Fatal("parse failed")
}
