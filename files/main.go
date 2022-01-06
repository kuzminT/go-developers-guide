package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// создаем файл
	// file, _ := os.Create("text.txt")
	// file.WriteString("1 строка \n")
	// file.WriteString("2 строка \n")
	// defer file.Close()
	// // переименовываем файл
	// // os.Rename("text.txt", "new_text.txt")
	// // удаляем файл
	// // os.Remove("new_text.txt")
	// info, err := file.Stat() // функция Stat возвращает информацию о файле и ошибку
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%+v", info) // true

	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	// Я заранее записал в файл 5 цифр, каждую на новой строке
	for s.Scan() { // возвращает true, пока файл не будет прочитан до конца
		fmt.Printf("New line: %s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
	}

}
