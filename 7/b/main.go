package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hacel/adventofcode/internal/scan"
)

func Average(d []int) (avg int) {
	for _, n := range d {
		avg += n
	}
	avg /= len(d)
	return
}

func main() {
	var crabs []int
	for s := range scan.S("7/input.txt") {
		for _, v := range strings.Split(s, ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			crabs = append(crabs, n)
		}
	}

	// Test lower and upper bound
	var cheapestPosition, cost [2]int
	cheapestPosition[0] = Average(crabs)
	cheapestPosition[1] = Average(crabs) + 1
	for _, c := range crabs {
		steps := int(math.Abs(float64(cheapestPosition[0] - c)))
		cost[0] += steps * (steps + 1) / 2
		steps = int(math.Abs(float64(cheapestPosition[1] - c)))
		cost[1] += steps * (steps + 1) / 2
	}

	if cost[0] < cost[1] {
		fmt.Printf("Best position: %d\n", cheapestPosition[0])
		fmt.Printf("Cost: %d\n", cost[0])
	} else {
		fmt.Printf("upBest position: %d\n", cheapestPosition[1])
		fmt.Printf("Cost: %d\n", cost[1])
	}
}
