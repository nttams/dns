package zone_handler

import (
	"fmt"
	msg "message"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func ReadZone() {
	fmt.Println("hi")

	readFile("./zone")
}

func readFile(file_path string) {
	file, err := os.Open(file_path)

	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string {}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// fmt.Println(sanner.Text())
	}

	for _, line := range lines {
		fmt.Println(strings.Fields(line))
	}

	fmt.Println(lines)
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