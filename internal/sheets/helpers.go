package sheets

import (
	"bufio"
	"fmt"
	"gsheets-intergation/internal/utils"
	"os"
	"strconv"
	"strings"
)

func PromptForSheetID() int64 {
	var sheetIDStr string

	for {
		fmt.Print("Enter the Google Sheets sheet ID: ")
		_, err := fmt.Scanln(&sheetIDStr)
		if err != nil {
			fmt.Println(utils.ErrInvalidInput)
			continue
		}

		sheetIDStr = strings.TrimSpace(sheetIDStr)
		if sheetIDStr == "" {
			fmt.Println(utils.ErrEmptyInput)
			continue
		}

		sheetID, err := strconv.ParseInt(sheetIDStr, 10, 64)
		if err != nil {
			fmt.Println(utils.ErrInvalidSheetID)
		} else {
			return sheetID
		}
	}
}

func PromptForSpreadSheetID() string {
	var sheetID string

	for {
		fmt.Print("Enter the Google Sheets spreadsheet ID: ")
		_, err := fmt.Scanln(&sheetID)
		if err != nil {
			fmt.Println(utils.ErrInvalidInput)
			continue
		}

		sheetID = strings.TrimSpace(sheetID)
		if sheetID == "" {
			fmt.Println(utils.ErrEmptyInput)
			continue
		}

		return sheetID
	}
}

func PromptForSheetName() string {
	var sheetName string

	for {
		fmt.Print("Enter the sheet name: ")
		_, err := fmt.Scanln(&sheetName)
		if err != nil {
			fmt.Println(utils.ErrInvalidInput)
			continue
		}

		sheetName = strings.TrimSpace(sheetName)
		if sheetName != "" {
			return sheetName
		} else {
			fmt.Println(utils.ErrEmptyInput)
		}
	}
}

func PromptForRange() string {
	var inputRange string

	for {
		fmt.Print("Enter the data range (e.g., 'A1:C3'): ")
		_, err := fmt.Scanln(&inputRange)
		if err != nil {
			fmt.Println(utils.ErrInvalidInput)
			continue
		}

		inputRange = strings.TrimSpace(inputRange)
		if inputRange != "" {
			return inputRange
		} else {
			fmt.Println(utils.ErrEmptyInput)
		}
	}
}

func PromptForSheetValue() ([][]interface{}, error) {
	fmt.Print("Enter data in JSON format: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputData := scanner.Text()

	if strings.TrimSpace(inputData) == "" {
		fmt.Println(utils.ErrEmptyInput)
		return nil, fmt.Errorf(utils.ErrEmptyInput)
	}

	data, err := utils.ParseUserInput(inputData)
	if err != nil {
		fmt.Printf("Failed to parse user input: %v\n", err)
		return nil, err
	}

	return data, nil
}
