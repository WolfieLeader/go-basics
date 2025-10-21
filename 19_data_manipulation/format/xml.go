package format

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

const CATALOG_FILE = "catalog.xml"

type Course struct {
	XMLName     xml.Name   `xml:"course"`
	Id          string     `xml:"id,attr"`           // attribute that identifies the course
	IsAvailable bool       `xml:"is_available,attr"` // attribute that shows availability
	IsOnline    bool       `xml:"is_online"`
	Buyers      int        `xml:"buyers"`
	Title       string     `xml:"title"`
	Description string     `xml:"description,omitempty"` // optional field
	Price       float64    `xml:"price"`
	Categories  []string   `xml:"categories>category"` // nested list
	CreatedAt   time.Time  `xml:"created_at"`
	UpdatedAt   *time.Time `xml:"updated_at"`
}

type Catalog struct {
	XMLName   xml.Name `xml:"catalog"`
	FirstName string   `xml:"author>first_name"` // nested struct fields
	LastName  string   `xml:"author>last_name"`  // nested struct fields
	Courses   []Course `xml:"course"`
}

func XmlWriteExample() {
	date := time.Date(2023, time.April, 5, 0, 0, 0, 0, time.UTC)
	catalog := Catalog{
		FirstName: "Alice",
		LastName:  "Smith",
		Courses: []Course{
			{
				Id:         "course-101",
				Buyers:     150,
				Title:      "Introduction to Go",
				Price:      49.99,
				Categories: []string{"Web", "Go", "Backend Development"},
				CreatedAt:  time.Date(2023, time.March, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt:  &date,
			},
			{
				Id:          "course-252",
				IsAvailable: false,
				IsOnline:    false,
				Buyers:      85,
				Title:       "Advanced Go Concurrency",
				Price:       79.99,
				Categories:  []string{"Programming", "Go", "Concurrency"},
				CreatedAt:   time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC),
			},
			{
				Id:          "course-087",
				IsAvailable: true,
				IsOnline:    true,
				Title:       "Embedded Systems with Go",
				Description: "Learn how to build embedded systems using Go programming language.",
				Price:       499.99,
				Categories:  []string{"Go", "Embedded Systems"},
				CreatedAt:   time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	compact, err := xml.Marshal(catalog)
	if err != nil {
		log.Fatalf("xml marshal error: %v", err)
	}
	fmt.Printf("- XML (compact):\n%s\n\n", compact)

	pretty, err := xml.MarshalIndent(catalog, "", "  ")
	if err != nil {
		log.Fatalf("xml marshal error: %v", err)
	}
	fmt.Printf("- XML (pretty):\n%s\n\n", pretty)
	out := append([]byte(xml.Header), pretty...)

	file := writerToFile(CATALOG_FILE) // from your helpers
	defer file.Close()

	if _, err := file.Write(out); err != nil {
		log.Fatalf("xml write file error: %v", err)
	}
	fmt.Printf("- Wrote XML file %s (%d bytes)\n", CATALOG_FILE, len(out))
}

func XmlReadExample() {
	data := readBytesFromFile(CATALOG_FILE)

	var catalog Catalog
	if err := xml.Unmarshal(data, &catalog); err != nil {
		log.Fatalf("xml unmarshal error: %v", err)
	}

	for index, course := range catalog.Courses {
		fmt.Printf("- Course %d: %+v\n", index+1, course)
	}
}
