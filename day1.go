package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
	"unicode/utf8"
)

func openFileFromArgs() (*os.File, error) {
	path := os.Args[1]

	file, err := os.Open(path)

	return file, err
}

func runeIsNumber(r rune) bool {
	_, err := strconv.ParseInt(string(r), 0, 64)

	if err != nil {
		return false
	}

	return true
}

func extractNumbersFromString(s string) *strings.Builder {
	stringReader := strings.NewReader(s)
	var builder strings.Builder

	for {
		ch, _, err := stringReader.ReadRune()
		if err != nil {
			break
		}

		if !runeIsNumber(ch) {
			continue
		}

		builder.WriteRune(ch)
	}

	return &builder
}

func parseFirstAndLastNumber(builder *strings.Builder) string {
		var firstAndLastNumber strings.Builder

		firstRune, _ := utf8.DecodeRuneInString(builder.String())

		firstAndLastNumber.WriteRune(firstRune)

		lastRune, _ := utf8.DecodeLastRuneInString(builder.String())

		firstAndLastNumber.WriteRune(lastRune)

		return firstAndLastNumber.String()
}

func sumSlice(numbers []int64) int64 {
	var sum int64

	for _, v := range numbers {
		sum += v
	}

	return sum
}


func main() {
	file, err := openFileFromArgs()
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var numbers [] int64

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		trimedLine := strings.TrimSpace(line)

		builder := extractNumbersFromString(trimedLine)

		firstAndLastNumber, _ := strconv.ParseInt(parseFirstAndLastNumber(builder), 0, 64)

		numbers = append(numbers, firstAndLastNumber)
	}

	sum := sumSlice(numbers)

	fmt.Println(sum)
}
