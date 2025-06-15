package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func stringsExample() {
	fmt.Println("\nString Manipulation Examples:")
	text := "Hello Gopher 🚀"
	fmt.Printf("- Original: %s\n", text)
	fmt.Printf("- Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("- Lowercase: %s\n", strings.ToLower(text))
	fmt.Printf("- Does it contain 'Gopher'? %t\n", strings.Contains(text, "Gopher"))
	fmt.Printf("- Index of 'Gopher': %d\n", strings.Index(text, "Gopher"))
	fmt.Printf("- Does it start with 'Hello'? %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("- Does it end with '🚀 '? %t\n", strings.HasSuffix(text, "🚀"))
	fmt.Printf("- Replace O's with 0's: %s\n", strings.ReplaceAll(text, "o", "0"))
	fmt.Printf("- Split by space: %v\n", strings.Split(text, " "))
	fmt.Printf("- Join with hyphen: %s\n", strings.Join(strings.Split(text, " "), "-"))
}

func convertExample() {
	fmt.Println("\nConversion Examples:")
	intStr := "123"
	intNum, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Printf("- Error converting '%s'", intStr)
		return
	}
	fmt.Printf("- Converted '%s' to integer: %d\n", intStr, intNum)

	floatStr := "123.456"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("- Error converting '%s'", floatStr)
		return
	}
	fmt.Printf("- Converted '%s' to float: %f\n", floatStr, floatNum)

	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Printf("- Error converting '%s'", boolStr)
		return
	}
	fmt.Printf("- Converted '%s' to boolean: %t\n", boolStr, boolVal)

	fmt.Printf("- Formatted integer: %s\n", strconv.Itoa(intNum))
	fmt.Printf("- Formatted float: %s\n", strconv.FormatFloat(floatNum, 'f', 2, 64))
}

func regexExample() {
	fmt.Println("\nRegular Expression Examples:")
	email1, email2, notEmail := "support@example.com", "sales@company.org", "not-an-email"
	text := fmt.Sprintf(`Contact us at %s or %s for more info. Do not use %s.`, email1, email2, notEmail)
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)

	fmt.Printf("- Text: %s\n", text)
	fmt.Printf("- Is '%s' a valid email? %t\n", email1, emailRegex.MatchString(email1))
	fmt.Printf("- Is '%s' a valid email? %t\n", email2, emailRegex.MatchString(email2))
	fmt.Printf("- Is '%s' a valid email? %t\n", notEmail, emailRegex.MatchString(notEmail))
	fmt.Printf("- Extracted email indexes: %v\n", emailRegex.FindAllStringIndex(text, -1)) // -1 means find all matches not limited
	fmt.Printf("- Extracted emails: %v\n", emailRegex.FindAllString(text, -1))             // -1 means find all matches not limited
	fmt.Printf("- Replaced emails with '[EMAIL]': %s\n", emailRegex.ReplaceAllString(text, "[EMAIL]"))
}

func main() {
	stringsExample()
	convertExample()
	regexExample()
}
