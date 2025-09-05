package main

import "fmt"

func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temps {
		var groupKey int
		if temp < 0 {
			groupKey = (int(temp) / 10) * 10
			if int(temp)%10 != 0 {
				groupKey -= 10
			}
		} else {
			groupKey = (int(temp) / 10) * 10
		}

		groups[groupKey] = append(groups[groupKey], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	grouped := groupTemperatures(temperatures)

	fmt.Println("Группировка температур:")
	for key := -30; key <= 40; key += 10 {
		if values, exists := grouped[key]; exists {
			fmt.Printf("%d: %v\n", key, values)
		}
	}
}
