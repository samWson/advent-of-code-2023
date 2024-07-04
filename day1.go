package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
	"unicode/utf8"
)

func isNumber(r rune) (rune, error) {
	_, err := strconv.ParseInt(string(r), 0, 64)

	if err != nil {
		return 0, errors.New("Not a number")
	}

	return r, nil
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
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

		stringReader := strings.NewReader(trimedLine)

		var builder strings.Builder

		for {
			ch, _, err := stringReader.ReadRune()
			if err != nil {
				break
			}

			num, err := isNumber(ch)
			if err != nil {
				continue
			}

			builder.WriteRune(num)
		}

		var finalNumber strings.Builder

		firstRune, _ := utf8.DecodeRuneInString(builder.String())

		finalNumber.WriteRune(firstRune)

		lastRune, _ := utf8.DecodeLastRuneInString(builder.String())

		finalNumber.WriteRune(lastRune)

		lastNumber, _ := strconv.ParseInt(finalNumber.String(), 0, 64)

		numbers = append(numbers, lastNumber)
	}

	var sum int64

	for _, v := range numbers {
		sum += v
	}

	fmt.Println(sum)
}
