package message

import (
	// "fmt"
	"strings"
)

type domainType uint8

const (
	InvalidType domainType = iota
	Literal
	Pointer
)

type Domain struct {
	domainType    domainType
	domainLiteral string
	pointer       uint16
}

// todo: only deal with liternal domain now
func (domain *Domain) encode() (result []byte) {
	if domain.domainType == Literal {
		labels := strings.Split(domain.domainLiteral, ".")
		for _, label := range labels {
			result = append(result, byte(len(label)))
			result = append(result, label...)
		}
		result = append(result, 0)
		return
	} else {
		panic("not implemented Pointer yet")
	}
}

// todo: should handle both literal and pointer
// rename req to msg
func parseDomain(req []byte, pos int) (string, int) {
	// RFC6891, 6.1.2.  Wire Format
	// OPT record
	if req[pos] == 0 {
		return string([]byte{ 0 }), 1
	}

	// RFC1035, 4.1.4. Message compression
	// pointer case
	if req[pos] & 0b1100_0000 != 0 {
		pointerHigh := uint16(req[pos] & 0b00111111) << 8
		pointerLow := uint16(req[pos + 1])
		pointedPos := pointerHigh | pointerLow
		domainLiteral, _ := parseDomain(req, int(pointedPos))
		return domainLiteral, 2
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