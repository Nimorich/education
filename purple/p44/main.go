package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Open input file
	inFile, err := os.Open("in.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inFile.Close()

	// Create output file
	outFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
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
			fmt.Printf("Error reading from file: %v\n", err)
			return
		}
		if r == '\xa0' {
			r = ' '
		}
		if r == '<' {
			// Potential start of a tag
			insideTag = true
			continue
		}

		if r == '>' && insideTag {
			// End of a tag
			insideTag = false
			continue
		}

		// Only write characters that are not inside tags
		if !insideTag {
			buffer.WriteRune(r)
		}
	}

	// Write the processed text to the output file
	_, err = writer.WriteString(buffer.String())
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}
