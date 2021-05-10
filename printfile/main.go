package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	filePath := os.Args[1]

	file, err := os.Open(filePath) // For read access.
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
	file.Close()
}
