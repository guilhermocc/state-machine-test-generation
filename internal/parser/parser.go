package parser

import (
	"encoding/csv"
	"io"
	"os"
)

func ParseStateMachineCsv(filePath string) (initialState string, events []string, transitions map[string][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	header, err := csvReader.Read()
	events = header[1:]
	transitions = make(map[string][]string)

	for {

		record, readErr := csvReader.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			err = readErr
			return
		}
		if initialState == "" {
			initialState = record[0]
		}
		transitions[record[0]] = record[1:]
	}

	return
}

func ParseEventsActionsCsv(filePath string) (eventsActions map[string]string, err error) {
	eventsActions = make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	csvReader.Comma = '|'
	csvReader.LazyQuotes = true

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		eventsActions[record[0]] = record[1]
	}

	return eventsActions, nil
}
