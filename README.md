# go-gsheets-implementation

This Go application provides a command-line interface (CLI) to interact with Google Sheets. It allows you to perform various operations on Google Sheets, such as retrieving data, adding data, updating existing data, and more.

## Prerequisites

Before using this application, make sure you have the following prerequisites in place:

1. **Google Cloud Platform (GCP) Account**: You need to have a GCP account. If you don't have one, you can sign up [here](https://cloud.google.com/).

2. **Create a GCP Project**: Create a new GCP project where you'll manage your credentials and enable the Google Sheets API. You can follow the [official guide](https://cloud.google.com/resource-manager/docs/creating-managing-projects) to create a new project.

3. **Create Service Account and Credentials**: You need to create a service account and download its credentials in JSON format. This service account will be used to access Google Sheets. Follow the instructions provided [here](https://developers.google.com/identity/protocols/oauth2/service-account) to create a service account and obtain credentials.

4. **Share Google Sheet**: Share the Google Sheets you want to work with the service account email address obtained in the previous step. Give the necessary permissions (e.g., edit or view) to the service account.

## Installation and running cmd

1. Clone this repository:

```sh
git clone https://github.com/ArtemBurakov/go-gsheets-integration.git
cd go-gsheets-implementation
```

2. Create a .env file in the project's root directory with the following content:

```sh
CREDENTIAL_FILE_LOCATION=./credentials.json
```
>Replace ./credentials.json with the actual path to your downloaded service account credentials file.

3. Run cmd:

```sh
cd cmd/integration
go run main.go
```

## Usage

The application provides various commands to interact with Google Sheets. Run the CLI with the following commands:

* getSheetData: Get sheet data and print it to the console.
* addDataToSheet: Add data to a sheet.
* updateExistingDataInSheet: Update existing data in a sheet.
* deleteExistingDataInSheet: Delete existing data in a sheet.
* createNewSheet: Create a new sheet in the spreadsheet.
* renameExistingSheet: Rename an existing sheet.
* deleteExistingSheet: Delete an existing sheet.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.