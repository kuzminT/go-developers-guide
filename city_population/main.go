package main

import "fmt"

/*
 * Группировка городов по численности населения в тысячах человек
 * от 10 до 100, от 100 до 1000 и более 1000:
 * groupCity map[int][]string{
 *	 10: []string{...},
 *	 100: []string{...},
 *	 1000: []string{...},
 * }
 *
 * Население городов в тысячах человек:
 * cityPopulation map[string]int{...}
 */

func main() {
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Сад и огород"},        // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город", "Село"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"},                      // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Село":      100,
		"Миллионик": 500,
	}

	contains := func(s []string, e string) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	for key := range cityPopulation {
		if !contains(groupCity[100], key) {
			delete(cityPopulation, key)
		}
	}

	fmt.Println(cityPopulation)

}
