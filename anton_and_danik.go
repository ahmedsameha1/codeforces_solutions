package main

import (
	"bufio"
	"fmt"
	"os"
)

func AntonAndDanik(results string) string {
	anton := 0
	danik := 0
	for _, winner := range results {
		if winner == 'A' {
			anton += 1
		} else {
			danik += 1
		}
	}
	if anton > danik {
		return "Anton"
	} else if danik > anton {
		return "Danik"
	} else {
		return "Friendship"
	}
}

func main() {
	var unneeded int
	var results string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &unneeded)
	fmt.Fscanf(reader, "%s\n", &results)
	fmt.Println(AntonAndDanik(results))
}
