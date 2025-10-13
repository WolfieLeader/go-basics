package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func stringsExample() {
	text := "Hello Gopher 🚀"
	fmt.Printf("- Original: %q\n", text)
	fmt.Printf("- Uppercase: %q\n", strings.ToUpper(text))
	fmt.Printf("- Lowercase: %q\n", strings.ToLower(text))
	fmt.Printf("- Does it contain 'Gopher'? %t\n", strings.Contains(text, "Gopher"))
	fmt.Printf("- Index of 'Gopher': %d\n", strings.Index(text, "Gopher"))
	fmt.Printf("- Does it start with 'Hello'? %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("- Does it end with '🚀'? %t\n", strings.HasSuffix(text, "🚀"))
	fmt.Printf("- Replace O's with *'s: %q\n", strings.ReplaceAll(text, "o", "*"))
	fmt.Printf("- Split by space: %q\n", strings.Split(text, " "))
	fmt.Printf("- Join with hyphen: %q\n", strings.Join(strings.Split(text, " "), "-"))
}

func convertExample() {
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("- Error converting string to integer")
		return
	}
	fmt.Printf("- Converted '123' to integer: %d\n", num)

	floatNum, err := strconv.ParseFloat("123.456", 64)
	if err != nil {
		fmt.Println("- Error converting string to float")
		return
	}
	fmt.Printf("- Converted '123.456' to float: %f\n", floatNum)

	boolVal, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("- Error converting string to boolean")
		return
	}
	fmt.Printf("- Converted 'true' to boolean: %t\n", boolVal)

	fmt.Printf("- Converted integer %d back to string: %q\n", num, strconv.Itoa(num))
	fmt.Printf("- Converted float %f back to string: %q\n", floatNum, strconv.FormatFloat(floatNum, 'f', 2, 64))
	fmt.Printf("- Converted boolean %t back to string: %q\n", boolVal, strconv.FormatBool(boolVal))
}

func regexExample() {
	email1, email2, notEmail := "support@example.com", "sales@company.org", "not-an-email"
	text := fmt.Sprintf(`Contact us at %s or %s for more info. Do not use %s.`, email1, email2, notEmail)
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)

	fmt.Printf("- Text: %q\n", text)
	fmt.Printf("- Is %q a valid email? %t\n", email1, emailRegex.MatchString(email1))
	fmt.Printf("- Is %q a valid email? %t\n", email2, emailRegex.MatchString(email2))
	fmt.Printf("- Is %q a valid email? %t\n", notEmail, emailRegex.MatchString(notEmail))

	// -1 means find all matches not limited
	fmt.Printf("- Found emails: %q\n", emailRegex.FindAllString(text, -1))

	// -1 means find all matches not limited
	fmt.Printf("- Matched email positions: %v\n", emailRegex.FindAllStringIndex(text, -1))
	fmt.Printf("- Replaced emails with '[EMAIL]': %q\n", emailRegex.ReplaceAllString(text, "[EMAIL]"))
}

func bytesExample() {
	fmt.Println("\nByte Slice Example:")
	buf := []byte("Hello Gophers") // String is like a read-only byte slice
	fmt.Printf("- Original byte slice: %q\n", string(buf))
	fmt.Printf("- Uppercase byte slice: %q\n", string(bytes.ToUpper(buf)))
	fmt.Printf("- Lowercase byte slice: %q\n", string(bytes.ToLower(buf)))
	fmt.Printf("- Does it contain 'Gophers'? %t\n", bytes.Contains(buf, []byte("Gophers")))
	fmt.Printf("- Index of 'Gophers': %d\n", bytes.Index(buf, []byte("Gophers")))
	fmt.Printf("- Replace O's with *'s: %q\n", string(bytes.ReplaceAll(buf, []byte("o"), []byte("*"))))
	fmt.Printf("- Split by space (unicode): %q\n", bytes.Split(buf, []byte(" ")))
	fmt.Printf("- Join with hyphen: %q\n", string(bytes.Join(bytes.Split(buf, []byte(" ")), []byte("-"))))
	fmt.Printf("- Equal to 'Hello GopheRs'? %t\n", bytes.Equal(buf, []byte("Hello GopheRs")))
	fmt.Printf("- EqualFold to 'hello gophers'? %t\n", bytes.EqualFold(buf, []byte("hello gophers")))
}

func iterationExample() {
	str := "Go 🚀!"

	fmt.Println("- Iterating over string bytes and runes:")
	for i := range len(str) { // Iterating over bytes (not safe for multi-byte characters)
		fmt.Printf("- Byte %d: %c, Unicode: %U\n", i, str[i], str[i])
	}

	fmt.Println("\n- Iterating over string runes:")
	for i, r := range str { // Iterating over runes is safe and handles multi-byte characters
		fmt.Printf("- Rune %d: %c, Unicode: %U\n", i, r, r)
	}
}

func main() {
	fmt.Println("\nString Manipulation Examples:")
	stringsExample()

	fmt.Println("\nConversion Examples:")
	convertExample()

	fmt.Println("\nRegular Expression Examples:")
	regexExample()

	fmt.Println("\nByte and Rune Iteration Examples:")
	iterationExample()

	fmt.Println("\nByte Slice Manipulation Examples:")
	bytesExample()
}
