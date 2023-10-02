package sheets

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsClient struct {
	Client        *sheets.Service
	SpreadsheetId string
}

func NewGoogleSheetsClient(credentialsFile, spreadsheetId string) (*GoogleSheetsClient, error) {
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &GoogleSheetsClient{
		Client:        client,
		SpreadsheetId: spreadsheetId,
	}, nil
}

func (gsc *GoogleSheetsClient) GetSheetData(range_ string) ([][]interface{}, error) {
	ctx := context.Background()

	resp, err := gsc.Client.Spreadsheets.Values.Get(gsc.SpreadsheetId, range_).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	if len(resp.Values) == 0 {
		return nil, nil
	}

	return resp.Values, nil
}

func (gsc *GoogleSheetsClient) AddData(range_ string, data [][]interface{}) error {
	ctx := context.Background()

	valueRange := &sheets.ValueRange{
		Values: data,
	}

	_, err := gsc.Client.Spreadsheets.Values.Append(gsc.SpreadsheetId, range_, valueRange).
		ValueInputOption("USER_ENTERED").
		InsertDataOption("INSERT_ROWS").
		Context(ctx).
		Do()
	if err != nil {
		return err
	}

	return nil
}

func (gsc *GoogleSheetsClient) CreateSheet(newSheetName string) error {
	ctx := context.Background()

	addSheetRequest := sheets.AddSheetRequest{
		Properties: &sheets.SheetProperties{
			Title: newSheetName,
		},
	}

	batchUpdateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &addSheetRequest,
			},
		},
	}

	_, err := gsc.Client.Spreadsheets.BatchUpdate(gsc.SpreadsheetId, &batchUpdateRequest).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (gsc *GoogleSheetsClient) UpdateData(range_ string, data [][]interface{}) error {
	ctx := context.Background()

	valueRange := &sheets.ValueRange{
		Values: data,
	}

	_, err := gsc.Client.Spreadsheets.Values.Update(gsc.SpreadsheetId, range_, valueRange).
		ValueInputOption("USER_ENTERED").
		Context(ctx).
		Do()
	if err != nil {
		return err
	}

	return nil
}

func (gsc *GoogleSheetsClient) RenameSheet(newSheetName string, sheetID int64) error {
	ctx := context.Background()

	requests := []*sheets.Request{
		{
			UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
				Properties: &sheets.SheetProperties{
					SheetId: sheetID,
					Title:   newSheetName,
				},
				Fields: "title",
			},
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}

	_, err := gsc.Client.Spreadsheets.BatchUpdate(gsc.SpreadsheetId, batchUpdateRequest).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (gsc *GoogleSheetsClient) DeleteData(range_ string) error {
	ctx := context.Background()
	clearValuesRequest := &sheets.ClearValuesRequest{}

	_, err := gsc.Client.Spreadsheets.Values.Clear(gsc.SpreadsheetId, range_, clearValuesRequest).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (gsc *GoogleSheetsClient) DeleteSheet(sheetId int64) error {
	ctx := context.Background()
	deleteSheetRequest := &sheets.DeleteSheetRequest{
		SheetId: sheetId,
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteSheet: deleteSheetRequest,
			},
		},
	}

	_, err := gsc.Client.Spreadsheets.BatchUpdate(gsc.SpreadsheetId, batchUpdateRequest).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}
