package camel_case_four

import (
	"fmt"
	"strings"
	"unicode"
)

func CamelCaseFour(line string) {
	parts := strings.Split(line, ";")
	if len(parts) != 3 {
		fmt.Println("Invalid input format")
		return
	}

	operation, typeWord, wordValue := parts[0], parts[1], parts[2]

	switch operation {
	case "S":
		switch typeWord {
		case "M":
			fmt.Println(convertCamelCaseToSpaces(wordValue[:len(wordValue)-2]))
		case "C", "V":
			fmt.Println(convertCamelCaseToSpaces(wordValue))
		}
	case "C":
		switch typeWord {
		case "M":
			fmt.Println(convertSpacesToCamelCase(wordValue) + "()")
		case "C":
			fmt.Println(convertSpacesToPascalCase(wordValue))
		case "V":
			fmt.Println(convertSpacesToCamelCase(wordValue))
		}
	default:
		fmt.Println("Invalid operation")
	}
}

func convertSpacesToPascalCase(input string) string {
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder
	for _, word := range words {
		for i, r := range word {
			if i == 0 {
				result.WriteRune(unicode.ToUpper(r))
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}

func convertSpacesToCamelCase(input string) string {
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder
	result.WriteString(words[0])

	for _, word := range words[1:] {
		for i, r := range word {
			if i == 0 {
				result.WriteRune(unicode.ToUpper(r))
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}

func convertCamelCaseToSpaces(input string) string {
	var result strings.Builder
	for i, r := range input {
		if unicode.IsUpper(r) && i != 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
