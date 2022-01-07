package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func readTask() (interface{}, interface{}, interface{}) {
	return 2.0, 0.5, "/"
}

var calcMap = map[string]func(val1, val2 float64){
	"+": func(val1, val2 float64) {
		fmt.Printf("%.4f", val1+val2)
	},
	"-": func(val1, val2 float64) {
		fmt.Printf("%.4f", val1-val2)
	},
	"*": func(val1, val2 float64) {
		fmt.Printf("%.4f", val1*val2)
	},
	"/": func(val1, val2 float64) {
		fmt.Printf("%.4f", val1/val2)
	},
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	val1, ok := value1.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
	val2, ok := value2.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}

	operationName, ok := operation.(string)
	if !ok {
		fmt.Println("неизвестная операция")
		return
	}

	if action, ok := calcMap[operationName]; ok {
		action(val1, val2)
	} else {
		fmt.Println("неизвестная операция")
		return
	}
}
