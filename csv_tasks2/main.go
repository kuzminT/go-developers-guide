package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("csv_tasks2/task.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// buf := make([]byte, 3000)
	// n, err := file.Read(buf)
	// fmt.Println(n, err)
	s := bufio.NewReader(file)

	// fmt.Printf("%+v", s)
	var i int
	for {
		i++
		c, err := s.ReadString(';')
		if err != nil {
			break
		}
		// fmt.Println(c)

		// buf := bytes.NewBuffer([]byte(s.Text()))
		if c == "0;" {
			fmt.Println(i)
			fmt.Println(c)
			return
		}
		// data := csv.NewReader(s.Split())
		// for {
		// 	row, err := r.Read()
		// 	if err != nil && err != io.EOF { // Здесь тоже нужно учитывать конец файла
		// 		panic(err)
		// 	}
		// 	for i, c := range row {
		// 		if strings.Contains(c, "0") {
		// 			// fmt.Println(row)
		// 			fmt.Println(i)
		// 			return
		// 		}
		// 	}

		// 	// fmt.Printf("New line: %s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
		// }
	}
}
