package message

import (
	"fmt"
	"strconv"
)

func ConvertRecordsToStrings(records []Record) (result []string) {
	for _, record := range records {
		switch record.qType {
		case Q_TYPE_A:
			result = append(result, ipv4ToString(record.rData))
		case Q_TYPE_AAAA:
			result = append(result, ipv6ToString(record.rData))
		default:
			panic("not recognize qType")
		}
	}
	return
}

func ipv4ToString(ip []byte) (result string) {
	for i := 0; i < 4; i++ {
		result += strconv.Itoa(int(ip[i])) + "."
	}
	return result[:len(result)-1]
}

func ipv6ToString(ip []byte) (result string) {
	for i := 0; i < 16; i += 2 {
		result += fmt.Sprintf("%x", int(ip[i])<<8+int(ip[i+1])) + ":"
	}
	return result[:len(result)-1]
}

func encodeUint16(value uint16) (result []byte) {
	result = append(result, byte((value&0b11111111_00000000)>>8))
	result = append(result, byte((value & 0b00000000_11111111)))
	return
}

func encodeUint32(value uint32) (result []byte) {
	result = append(result, byte((value&0b11111111_00000000_00000000_00000000)>>24))
	result = append(result, byte((value&0b00000000_11111111_00000000_00000000)>>16))
	result = append(result, byte((value&0b00000000_00000000_11111111_00000000)>>8))
	result = append(result, byte((value & 0b00000000_00000000_00000000_11111111)))
	return
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
