package main

import (
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	// Create a new Excel file
	file := xlsx.NewFile()

	// Add a new sheet named "Cover"
	sheet, err := file.AddSheet("Cover")
	if err != nil {
		log.Fatalf("Error creating sheet: %v", err)
	}

	// Define headers and data
	headers := []string{"City Name", "Site Name", "Center Name", "Circle", "State", "Circle Code", "Site Code"}
	data := []string{"Bangalore", "IMS", "bangalore", "ka", "Karnataka", "ka", "bglr"}

	// Add headers column-wise
	for _, header := range headers {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = header
	}

	// Add data column-wise
	for _, value := range data {
		row := sheet.Rows[0] // Get the first row
		cell := row.AddCell()
		cell.Value = value
	}

	// Save the Excel file
	err = file.Save("Cover.xlsx")
	if err != nil {
		log.Fatalf("Error saving file: %v", err)
	}

	log.Println("Excel file saved successfully.")
}
