package reading_tool

import (
	"encoding/csv"
	"os"
)

func ReadData(fileName string) ([][]string, error) {

	f, errOpenFile := os.Open(fileName)

	if errOpenFile != nil {
		return [][]string{}, errOpenFile
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true

	// Skip first line -> Header
	if _, errReadCSV := r.Read(); errReadCSV != nil {
		return [][]string{}, errReadCSV
	}

	records, errReadAll := r.ReadAll()
	if errReadAll != nil {
		return [][]string{}, errReadAll
	}

	return records, nil
}
