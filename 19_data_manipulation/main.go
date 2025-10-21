package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/19_data_manipulation/encoding"
	"github.com/WolfieLeader/go-basics/19_data_manipulation/format"
)

func main() {
	fmt.Println("Base64 Simple Example:")
	encoding.Base64Example()
	fmt.Println()

	fmt.Println("Base64 Encoder Example:")
	encoding.Base64EncoderExample()
	fmt.Println()

	fmt.Println("Hex Simple Example:")
	encoding.HexExample()
	fmt.Println()

	fmt.Println("Hex Encoder Example:")
	encoding.HexEncoderExample()
	fmt.Println()

	fmt.Println("Base32 Simple Example:")
	encoding.Base32Example()
	fmt.Println()

	fmt.Println("Base32 Encoder Example:")
	encoding.Base32EncoderExample()
	fmt.Println()

	// JSON v2 examples will be added in 1.26
	fmt.Println("JSON V1 Write Example:")
	format.JsonV1WriteExample()
	fmt.Println()

	fmt.Println("JSON V1 Read Example:")
	format.JsonV1ReadExample()
	fmt.Println()

	fmt.Println("JSON V1 Encoder Write Example:")
	format.JsonV1EncoderExample()
	fmt.Println()

	fmt.Println("JSON V1 Decoder Read Example:")
	format.JsonV1DecoderExample()
	fmt.Println()

	fmt.Println("CSV Write Example:")
	format.CsvWriteExample()
	fmt.Println()

	fmt.Println("CSV Read Example:")
	format.CsvReadExample()
	fmt.Println()

	fmt.Println("XML Write Example:")
	format.XmlWriteExample()
	fmt.Println()

	fmt.Println("XML Read Example:")
	format.XmlReadExample()
	fmt.Println()
}
