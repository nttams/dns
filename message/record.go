package message

type Record struct {
	domain   string
	qType    uint16
	qClass   uint16
	ttl      uint32
	rdLength uint16
	rData    []byte
}

func NewRecord(domain string, qType, qClass uint16, ttl uint32, rdLength uint16, rData []byte) Record {
	return Record{domain, qType, qClass, ttl, rdLength, rData}
}

func (record *Record) encode() (result []byte) {
	// result = append(result, []byte(record.domain)...)
	result = append(result, encodeDomain(record.domain)...)
	result = append(result, encodeUint16(record.qType)...)
	result = append(result, encodeUint16(record.qClass)...)
	result = append(result, encodeUint32(record.ttl)...)
	result = append(result, encodeUint16(record.rdLength)...)
	result = append(result, []byte(record.rData)...)
	return
}

func encodeRecords(records []Record) (result []byte) {
	for _, record := range records {
		result = append(result, record.encode()...)
	}
	return
}

func parseRecord(req []byte, pos int) (Record, int) {
	start := pos
	domain, count := parseDomain(req, pos)
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
	rData := rData_slice

	pos += int(rdLength)

	record := Record{domain, qType, qClass, ttl, rdLength, rData}

	return record, pos - start
}

func parseRecords(req []byte, pos, recordCount int) ([]Record, int) {
	records := []Record{}
	start := pos

	for i := 0; i < recordCount; i++ {
		record, count := parseRecord(req, pos)
		pos += count

		records = append(records, record)
	}
	return records, pos - start
}

func (record *Record) GetDomain() string {
	return record.domain
}
