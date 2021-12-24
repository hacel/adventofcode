package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hacel/adventofcode/internal/scan"
)

func main() {
	var fish [9]int
	for s := range scan.S("6/input.txt") {
		for _, v := range strings.Split(s, ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			fish[n]++
		}
	}

	for i := 0; i < 80; i++ {
		spawned := fish[0]
		for i := 1; i < len(fish); i++ {
			fish[i-1] = fish[i]
		}
		fish[8] = spawned
		fish[6] += spawned
	}

	count := 0
	for i := 0; i < len(fish); i++ {
		count += fish[i]
	}
	fmt.Printf("Number of fish: %d\n", count)
}
