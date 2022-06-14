package zone_handler

import (
	"fmt"
	"errors"
	msg "message"
	"os"
	"bufio"
	"strings"
	"strconv"
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

func ReadZone() {
	// readWholeFile("./zone")
	fmt.Println(readWholeFile("./zone"))
}

func readWholeFile(file_path string) ([]msg.Record, error) {
	file, err := os.Open(file_path)

	if err != nil {
		return nil, errors.New("can not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string {}

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
	q_type, _ := strconv.Atoi(fields[3])
	q_class, _ := strconv.Atoi(fields[2])
	ttl, _ := strconv.Atoi(fields[1])
	r_data := fields[4]
	rd_length := len(r_data)

	return msg.NewRecord(
		name,
		uint16(q_type),
		uint16(q_class),
		uint32(ttl),
		uint16(rd_length),
		r_data,
	)
}