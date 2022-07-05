package zone_handler

import (
	// "fmt"
	"bufio"
	"errors"
	msg "message"
	"os"
	"strconv"
	"strings"
)

func FindDomain(domain string) []msg.Record {

	result := []msg.Record{}

	allRecords, _ := readWholeFile("./zone")

	for _, record := range allRecords {
		if record.GetDomain() == domain {
			result = append(result, record)
		}
	}

	return result
}

func readWholeFile(filePath string) ([]msg.Record, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("can not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	result := []msg.Record{}
	for _, line := range lines {
		result = append(result, parseLineToRecord(line))
	}

	return result, nil
}

func parseLineToRecord(line string) msg.Record {
	fields := strings.Fields(line)

	name := fields[0]
	qType, _ := strconv.Atoi(fields[3])
	qClass, _ := strconv.Atoi(fields[2])
	ttl, _ := strconv.Atoi(fields[1])
	rData := fields[4]
	rdLength := len(rData)

	return msg.NewRecord(
		name,
		uint16(qType),
		uint16(qClass),
		uint32(ttl),
		uint16(rdLength),
		rData,
	)
}
