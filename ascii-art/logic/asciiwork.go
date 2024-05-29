package logic

import (
	"ascii-art-web/ascii-art/validation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertWordsToAsciiArtWithNewLinesStr(word string, letters [95][8]string) string {
	word = strings.ReplaceAll(word, "\n", "\\n")

	words := strings.Split(word, "\\n")

	resStr := ""
	if validation.OnlyHasNewLines(words) {
		words = words[1:]
	}

	for _, wordToRead := range words {
		if wordToRead == "" {

			resStr += "\n"
			continue
		}
		resStr += ConvertInputToAsciiArtStr(wordToRead, letters)
	}
	return resStr
}

func ConvertInputToAsciiArtStr(str string, letters [95][8]string) string {
	resStr := ""
	for i := 0; i < 8; i++ {
		for _, strLetterIndex := range str {
			if strLetterIndex >= 32 && strLetterIndex <= 126 {
				strLetterIndex -= 32
				for _, line := range letters[strLetterIndex][i] {
					resStr += string(line)
				}
			}
		}
		resStr += "\n"
	}

	return resStr
}

func GetAsciiArtLetters(filePath string) ([95][8]string, error) {
	var empty [95][8]string

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return empty, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var letters [95][8]string
	letterIndex := 0
	for scanner.Scan() {

		for i := 0; i < 8; i++ {
			scanner.Scan()
			line := scanner.Text()

			letters[letterIndex][i] = line
		}
		if err != nil {
			fmt.Println(err)
			return empty, err

		}
		letterIndex++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return empty, err

	}

	return letters, nil
}
