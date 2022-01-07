package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	// if strings.Contains(info.Name(), "csv") {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(data) > 1 {
		fmt.Println(path)
		fmt.Printf("find data: %s\n", data[4][2])
	}
	// }

	// fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	root, _ := filepath.Abs("csv_tasks/task") // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Some error: %v\n", err)
	}

}
