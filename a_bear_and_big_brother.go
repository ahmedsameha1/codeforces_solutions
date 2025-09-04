package main

import (
	"bufio"
	"fmt"
	"os"
)

func A_Bear_and_Big_Brother(bear1, bear2 int) int {
	counter := 0
	for {
		counter++
		bear1 *= 3
		bear2 *= 2
		if bear1 > bear2 {
			break
		}
	}
	return counter
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var bear1weight int
	var bear2weight int
	fmt.Fscanf(reader, "%d %d\n", &bear1weight, &bear2weight)
	fmt.Println(A_Bear_and_Big_Brother(bear1weight, bear2weight))
}
