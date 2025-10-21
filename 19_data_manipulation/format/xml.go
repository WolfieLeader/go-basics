package format

import (
	"encoding/xml"
	"fmt"
	"log"
)

const CATALOG_FILE = "catalog.xml"

type Product struct {
	XMLName    xml.Name `xml:"product"`
	ID         int      `xml:"id"`
	Name       string   `xml:"name"`
	Price      float64  `xml:"price"`
	InStock    bool     `xml:"in_stock"`
	Categories []string `xml:"categories>category"` // nested list
}

type Catalog struct {
	XMLName  xml.Name  `xml:"catalog"`
	Products []Product `xml:"product"`
}

func XmlWriteExample() {
	catalog := Catalog{
		Products: []Product{
			{ID: 1, Name: "Laptop", Price: 1299.99, InStock: true, Categories: []string{"Computers", "Electronics"}},
			{ID: 2, Name: "Headphones", Price: 199.50, InStock: false, Categories: []string{"Audio", "Accessories"}},
			{ID: 3, Name: "Coffee Maker", Price: 89.99, InStock: true, Categories: []string{"Home Appliances", "Kitchen"}},
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

	var cat Catalog
	if err := xml.Unmarshal(data, &cat); err != nil {
		log.Fatalf("xml unmarshal error: %v", err)
	}

	fmt.Printf("- Decoded XML catalog with %d products:\n", len(cat.Products))
	for i, p := range cat.Products {
		fmt.Printf("  [%d] %+v\n", i+1, p)
	}
}
