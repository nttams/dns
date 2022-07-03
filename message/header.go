package message

type Header struct {
	id           uint16
	headerFlags  HeaderFlags
	qdCount      uint16
	anCount      uint16
	nsCount      uint16
	arCount      uint16
}

type HeaderFlags struct {
	qr      bool
	opCode  byte
	aa      bool
	tc      bool
	rd      bool
	ra      bool
	z       bool
	ad      bool
	cd      bool
	rCode   byte
}

func newRequestHeaderFlags() (headerFlags HeaderFlags) {
	headerFlags.rd = true
	return
}

func newRequestHeader(id uint16) (header Header) {
	header.id = id
	header.headerFlags = newRequestHeaderFlags()
	return
}

func NewHeader(id uint16, headerFlags HeaderFlags, qdCount, anCount, nsCount, arCount uint16) *Header {
	return &Header{id, headerFlags, qdCount, anCount, nsCount, arCount}
}

func (hf *HeaderFlags) encode() []byte {

	var first byte = 0b0000_0000
	var second byte = 0b0000_0000

	if hf.qr {
		first = first | (1 << 7)
	}
	first = first | (hf.opCode << 3)
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
		second = second | (1 << 6)
	}
	if hf.ad {
		second = second | (1 << 5)
	}
	if hf.cd {
		second = second | (1 << 4)
	}
	second = second | hf.rCode

	return []byte{first, second}
}

func (header *Header) encode() []byte {
	result := []byte{}

	result = append(result, byte((header.id&0b1111_1111_0000_0000)>>8))
	result = append(result, byte((header.id&0b0000_0000_1111_1111)>>0))

	result = append(result, header.headerFlags.encode()...) //todo: learn this

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

func parseHeader(req []byte, pos int) (Header, int) {
	id := parseUint16(req, pos)
	pos += 2

	headerFlags := parseHeaderFlags(req, pos)
	pos += 2

	qdCount := parseUint16(req, pos)
	pos += 2

	anCount := parseUint16(req, pos)
	pos += 2

	nsCount := parseUint16(req, pos)
	pos += 2

	arCount := parseUint16(req, pos)

	header := Header{id, headerFlags, qdCount, anCount, nsCount, arCount}
	return header, 12
}

func parseHeaderFlags(req []byte, pos int) HeaderFlags {
	qr := req[pos]&0b1000_0000 == 1
	opCode := req[pos] & 0b0111_1000
	aa := req[pos]&0b0000_0100 == 1
	tc := req[pos]&0b0000_0010 == 1
	rd := req[pos]&0b0000_0001 == 1

	ra := req[pos+1]&0b1000_0000 == 1
	z := req[pos+1]&0b0100_0000 == 1
	ad := req[pos+1]&0b0010_0000 == 1
	cd := req[pos+1]&0b0001_0000 == 1
	rCode := req[pos+1] & 0b0000_1111

	reqult := HeaderFlags{qr, opCode, aa, tc, rd, ra, z, ad, cd, rCode}

	return reqult
}