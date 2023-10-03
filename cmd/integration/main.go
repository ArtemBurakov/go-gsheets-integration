package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/turret-io/go-menu/menu"
	"gsheets-intergation/internal/sheets"
	"gsheets-intergation/internal/utils"
	"os"
)

var gsc *sheets.GoogleSheetsClient

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	credentialsFile := os.Getenv("CREDENTIAL_FILE_LOCATION")
	spreadsheetId := sheets.PromptForSpreadSheetID()

	gsc, err = sheets.NewGoogleSheetsClient(credentialsFile, spreadsheetId)
	if err != nil {
		fmt.Printf("failed to create Google Sheets client: %v\n", err)
		return
	}
	fmt.Print("Google Sheets client initialized successfully.\n\n")
}

func getSheetData(args ...string) error {
	sheetName := sheets.PromptForSheetName()
	rangeStr := sheets.PromptForRange()

	data, err := gsc.GetSheetData(sheetName + "!" + rangeStr)
	if err != nil {
		fmt.Printf("failed to get sheet data: %v\n", err)
	}

	formattedData := utils.FormatAsJSON(data)
	fmt.Println(formattedData)

	return nil
}

func addDataToSheet(args ...string) error {
	sheetName := sheets.PromptForSheetName()
	rangeStr := sheets.PromptForRange()

	data, err := sheets.PromptForSheetValue()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = gsc.AddData(sheetName+"!"+rangeStr, data)
	if err != nil {
		fmt.Printf("failed to add data to sheet: %v\n", err)
	}

	return nil
}

func updateExistingDataInSheet(args ...string) error {
	sheetName := sheets.PromptForSheetName()
	rangeStr := sheets.PromptForRange()

	data, err := sheets.PromptForSheetValue()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = gsc.UpdateData(sheetName+"!"+rangeStr, data)
	if err != nil {
		fmt.Printf("failed to update existing data: %v\n", err)
	}

	return nil
}

func deleteExistingDataInSheet(args ...string) error {
	sheetName := sheets.PromptForSheetName()
	rangeStr := sheets.PromptForRange()

	err := gsc.DeleteData(sheetName + "!" + rangeStr)
	if err != nil {
		fmt.Printf("failed to delete existing data: %v\n", err)
	}

	return nil
}

func createNewSheet(args ...string) error {
	newSheetName := sheets.PromptForSheetName()

	err := gsc.CreateSheet(newSheetName)
	if err != nil {
		fmt.Printf("failed to create a new sheet: %v\n", err)
	}

	return nil
}

func renameExistingSheet(args ...string) error {
	newSheetName := sheets.PromptForSheetName()
	sheetID := sheets.PromptForSheetID()

	err := gsc.RenameSheet(newSheetName, sheetID)
	if err != nil {
		fmt.Printf("failed to rename existing sheet: %v\n", err)
	}

	return nil
}

func deleteExistingSheet(args ...string) error {
	sheetId := sheets.PromptForSheetID()

	err := gsc.DeleteSheet(sheetId)
	if err != nil {
		fmt.Printf("failed to delete existing sheet: %v\n", err)
	}

	return nil
}

func main() {
	if gsc != nil {
		commandOptions := []menu.CommandOption{
			{Command: "getSheetData", Description: "Get sheet data and print to the console", Function: getSheetData},
			{Command: "addDataToSheet", Description: "Add data to sheet", Function: addDataToSheet},
			{Command: "updateExistingDataInSheet", Description: "Update existing data in sheet", Function: updateExistingDataInSheet},
			{Command: "deleteExistingDataInSheet", Description: "Delete existing data in sheet", Function: deleteExistingDataInSheet},
			{Command: "createNewSheet", Description: "Create a new sheet in spreadsheet", Function: createNewSheet},
			{Command: "renameExistingSheet", Description: "Rename existing sheet", Function: renameExistingSheet},
			{Command: "deleteExistingSheet", Description: "Delete existing sheet", Function: deleteExistingSheet},
		}
		menuOptions := menu.NewMenuOptions("\n> ", 0)
		integrationMenu := menu.NewMenu(commandOptions, menuOptions)

		integrationMenu.Start()
	}
}
