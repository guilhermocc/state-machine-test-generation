package parser

import (
	"encoding/csv"
	"io"
	"os"
)

func ParseStateMachineCsv(filePath string) (events []string, transitions map[string][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	header, err := csvReader.Read()
	events = header[1:]
	transitions = make(map[string][]string)

	for {

		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		transitions[record[0]] = record[1:]
	}

	return events, transitions, nil
}
