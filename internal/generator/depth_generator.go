package generator

import "github.com/guilhermocc/test-case-generator/internal/parser"

func GenerateTestCases(inputFilePath string) {
	events, transitions, err := parser.ParseStateMachineCsv(inputFilePath)
	if err != nil {
		panic(err)
	}

	testCases := generateTestCases(events, transitions)
	writeTestCases(testCases)
}

func generateTestCases(events []string, transitions map[string][]string) []string {
	return []string{}
}

func writeTestCases(cases interface{}) {

}
