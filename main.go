package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strings"
)

func main() {
	r := csv.NewReader(bufio.NewReader(os.Stdin))
	fileName := os.Args[1]
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		keyword := record[0]
		merged := strings.Split(record[1], "#")
		userId := merged[0]
		tag := merged[1]

		score := strings.TrimSpace(merged[2])

		values := []string{keyword, userId, tag, score}

		err = writer.Write(values)
		if err != nil {
			fmt.Println(err)
		}
	}

}
