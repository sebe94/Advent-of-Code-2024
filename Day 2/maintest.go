package main

import "fmt"

func isSafe(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		//Differenz zwischen zwei aufeinander folgenden Zahlen
		diff := numbers[i+1] - numbers[i]

		//Bedingung größer / kleiner als 3
		if diff > 3 || diff < -3 {
			return false
		}

		//Bedingung weder auf/abstieg
		if diff == 0 {
			return false
		}

		//bedingung positiv und negativ
		if i > 0 {
			prevDiff := numbers[i] - numbers[i-1]
			if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
				return false
			}
		}
	}
	return true
}

func main() {
	//Anlegen der reports
	testCases := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	for _, testCase := range testCases {
		if isSafe(testCase) {
			fmt.Println(testCase, "-> Safe")
		} else {
			fmt.Println(testCase, "-> Unsafe")
		}
	}

}
