package validation

import (
	"fmt"
)

func ConsistsOnlyFromAsciiChars(str string) error {
	for _, strLetterIndex := range str {
		if !(strLetterIndex >= 32 && strLetterIndex <= 126) && strLetterIndex != 13 && strLetterIndex != 10 {
			return fmt.Errorf("Only ascci chars between 32 and 126")
		}
	}
	return nil
}

func OnlyHasNewLines(words []string) bool {
	for _, wordToRead := range words {
		if wordToRead != "" {
			return false
		}
	}
	return true
}
