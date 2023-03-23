package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Define command-line flags
	urlFlag := flag.String("url", "", "URL to JavaScript code to beautify")
	localFlag := flag.String("local", "", "File path to local JavaScript code to beautify")
	flag.Parse()

	var jsCode string
	if *urlFlag != "" {
		// Get JavaScript code from URL
		resp, err := http.Get(*urlFlag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		jsCode = string(body)
	} else if *localFlag != "" {
		// Get JavaScript code from local file
		codeBytes, err := ioutil.ReadFile(*localFlag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		jsCode = string(codeBytes)
	} else {
		// No input source specified
		fmt.Println("Error: Please specify an input source with -url or -local")
		return
	}

	// Beautify JavaScript code
	var beautifiedCode bytes.Buffer
	err := jsBeautify(strings.NewReader(jsCode), &beautifiedCode)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print beautified code to console
	fmt.Println(beautifiedCode.String())
}

func jsBeautify(input *strings.Reader, output *bytes.Buffer) error {
	const indentSize = 2
	const indentChar = ' '
	var indentLevel int
	var inString bool
	var currentByte byte
	var lastByte byte
	var err error

	for {
		currentByte, err = input.ReadByte()
		if err != nil {
			break
		}

		if currentByte == '"' && lastByte != '\\' {
			inString = !inString
		}

		if !inString {
			switch currentByte {
			case '{':
				fmt.Fprint(output, "{\n")
				indentLevel++
				fmt.Fprint(output, strings.Repeat(string(indentChar), indentSize*indentLevel))
			case '}':
				if indentLevel > 0 {
					indentLevel--
				}
				fmt.Fprint(output, "\n", strings.Repeat(string(indentChar), indentSize*indentLevel), "}")
			case ';':
				fmt.Fprint(output, ";\n", strings.Repeat(string(indentChar), indentSize*indentLevel))
			}
		}

		fmt.Fprint(output, string(currentByte))
		if currentByte == '\n' {
			fmt.Fprint(output, strings.Repeat(string(indentChar), indentSize*indentLevel))
		}

		lastByte = currentByte
	}

	return nil
}
