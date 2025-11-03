package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inFile, err := os.Open("in.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Printf("Error creating out file: %v\n", err)
		return
	}
	defer outFile.Close()

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	var buffer strings.Builder
	insideTag := false

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error read file: %v\n", err)
			return
		}
		if r == '\xa0' {
			r = ' '
		}
		if r == '<' {
			insideTag = true
			continue
		}
		if r == '>' && insideTag {
			insideTag = false
			continue
		}
		if !insideTag {
			buffer.WriteRune(r)
		}
	}
	_, err = writer.WriteString(buffer.String())
	if err != nil {
		fmt.Printf("Error write to file: %v\n", err)
		return
	}

}
