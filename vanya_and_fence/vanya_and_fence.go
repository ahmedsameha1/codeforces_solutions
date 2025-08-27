package main

import (
	"fmt"
)

func VanyaAndFence(heightOfFence int, heightOfPeople []int) int {
	result := 0
	for _, height := range heightOfPeople {
		if height > heightOfFence {
			result += 2
		} else {
			result += 1
		}
	}
	return result
}

func main() {
	var countOfPeople, heightOfFence int
	fmt.Scanln(&countOfPeople, &heightOfFence)
	var heightOfPeople []int = make([]int, countOfPeople)
	for i := range countOfPeople {
		fmt.Scanf("%d", &heightOfPeople[i])
	}
	fmt.Println(VanyaAndFence(heightOfFence, heightOfPeople))
}
