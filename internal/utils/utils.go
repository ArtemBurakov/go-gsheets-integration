package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	ErrInvalidInput   = "invalid input. Please enter a valid value."
	ErrEmptyInput     = "input cannot be empty. Please provide a value."
	ErrInvalidSheetID = "invalid sheet ID. Please enter a valid integer."
)

func FormatAsJSON(data [][]interface{}) string {
	var formattedData []string

	for _, row := range data {
		var formattedRow []string
		for _, cell := range row {
			formattedRow = append(formattedRow, fmt.Sprintf(`"%v"`, cell))
		}
		formattedData = append(formattedData, "["+strings.Join(formattedRow, ", ")+"]")
	}

	return "[" + strings.Join(formattedData, ", ") + "]"
}

func ParseUserInput(inputData string) ([][]interface{}, error) {
	if strings.TrimSpace(inputData) == "" {
		return nil, fmt.Errorf(ErrEmptyInput)
	}

	var data [][]interface{}
	err := json.Unmarshal([]byte(inputData), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user input: %v", err)
	}

	return data, nil
}
