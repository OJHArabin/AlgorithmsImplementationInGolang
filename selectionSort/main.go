package main

import "fmt"

func main() {

	var toSort = []int{24, 36, 3, 20, 1}
	selectionSort(toSort)
	fmt.Println(toSort)
}

func selectionSort(toSort []int) {
	for i := 0; i < len(toSort); i++ {
		minindex := i
		//searching smallest vlaue from array
		for j := i; j < len(toSort); j++ {
			if toSort[j] < toSort[minindex] {
				minindex = j
			}
		}
		//sorting
		toSort[i], toSort[minindex] = toSort[minindex], toSort[i]
	}
}
