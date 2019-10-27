package main

import "fmt"

func main() {

	var toSort = []int{24, 36, 3, 20, 1, -2}
	bubbleSort(toSort)
	fmt.Println(toSort)
}

func bubbleSort(toSort []int) {
	for i := 0; i < len(toSort)-1; i++ {

		for j := 0; j < len(toSort)-i-1; j++ {
			if toSort[j+1] < toSort[j] {
				//sorting
				toSort[j], toSort[j+1] = toSort[j+1], toSort[j]
			}
		}

	}
}
