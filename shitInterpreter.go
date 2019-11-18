package main

import (
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	file_bytes, _ := ioutil.ReadFile(filename)

	PARSE_INSTRUCTIONS(LEX_TEXT(string(file_bytes)))
}
