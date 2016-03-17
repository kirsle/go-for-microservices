package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Array
	words := []string{
		"Hello",
		"world",
	}

	// Map
	fruits := map[string]string{
		"Apple":  "red",
		"Banana": "yellow",
	}

	// Loop over an array.
	for i, word := range words {
		fmt.Printf("Word %d: %s\n", i, word)
	}

	// Loop over a map.
	for key, value := range fruits {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Read a file.
	fh, err := os.Open("/etc/fstab")
	if err != nil {
		panic(fmt.Sprintf("Couldn't read the file: %s", err))
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Printf("Line: %s\n", scanner.Text())
	}
}
