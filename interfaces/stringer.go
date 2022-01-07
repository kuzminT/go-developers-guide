package main

import (
	"fmt"
	"strings"
)

// пакет используется для проверки ответа, не удаляйте его

type battery struct {
	energy string
}

func (b battery) String() string {
	formatStr := make([]string, len(b.energy))
	for _, v := range b.energy {
		vs := string(v)
		if vs == "1" {
			formatStr = append(formatStr, "X")
		} else {
			formatStr = append([]string{" "}, formatStr...)
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(formatStr, ""))
}

func main() {
	// batteryForTest - не забывайте об имени
	// } Скобка, закрывающая функцию main() вам не видна, но она есть
	var energy string
	fmt.Scan(&energy)

	batteryForTest := battery{energy}

	fmt.Printf("%s", batteryForTest)

}
