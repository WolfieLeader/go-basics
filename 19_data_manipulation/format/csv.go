package format

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
)

const BOOKS_FILE = "books.csv"

type BookCSV struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func CsvWriteExample() {
	records := []BookCSV{
		{Title: "Go in Action", Author: "William Kennedy", Year: 2015, Price: 39.99},
		{Title: "The Go Programming Language", Author: "Donovan & Kernighan", Year: 2015, Price: 44.95},
		{Title: "Concurrency in Go", Author: "Katherine Cox-Buday", Year: 2017, Price: 49.99},
	}

	file := writerToFile(BOOKS_FILE)
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	if err := csvWriter.Write([]string{"Title", "Author", "Year", "Price"}); err != nil {
		log.Fatalf("CSV Write header error: %v", err)
	}

	for _, record := range records {
		row := []string{
			record.Title,
			record.Author,
			strconv.Itoa(record.Year),
			strconv.FormatFloat(record.Price, 'f', 2, 64),
		}
		if err := csvWriter.Write(row); err != nil {
			log.Fatalf("CSV Write record error: %v", err)
		}
	}

	fmt.Printf("- Wrote %d records to CSV file %s\n", len(records), BOOKS_FILE)
}

func CsvReadExample() {
	file := readerFromFile(BOOKS_FILE)
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	header, err := csvReader.Read() // Read header
	if err != nil {
		log.Fatalf("CSV Read header error: %v", err)
	}

	fmt.Printf("- CSV Header: %v\n", header)

	var books []BookCSV
	for {
		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("CSV Read record error: %v", err)
		}

		year, _ := strconv.Atoi(record[2])
		price, _ := strconv.ParseFloat(record[3], 64)
		book := BookCSV{Title: record[0], Author: record[1], Year: year, Price: price}
		books = append(books, book)
		fmt.Printf("- Record %d: %+v\n", len(books), book)
	}
	fmt.Println()

	fmt.Printf("- Read %d records from CSV file %s\n", len(books), BOOKS_FILE)
}
