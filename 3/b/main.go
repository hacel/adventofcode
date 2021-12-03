package main

import (
	"fmt"
	"strconv"

	"github.com/hacel/adventofcode/internal/scan"
)

const WIDTH = 12

func bitAt(n int64, i int) int64 {
	return (n >> (WIDTH - i - 1)) & 0x1
}

func main() {
	var o2 []int64
	var co2 []int64
	for line := range scan.B("3/input.txt") {
		n, err := strconv.ParseInt(string(line), 2, 64)
		if err != nil {
			panic(err)
		}
		o2 = append(o2, n)
		co2 = append(co2, n)
	}

	// Oxygen generator rating
	for i := 0; i < WIDTH; i++ {
		if len(o2) == 1 {
			break
		}

		// Get most common bit
		var n0, n1 int
		for j := 0; j < len(o2); j++ {
			if bitAt(o2[j], i) == 1 {
				n1++
			} else {
				n0++
			}
		}
		var bit int64
		if n1 >= n0 {
			bit = 1
		}

		// Filter
		for j := 0; j < len(o2); {
			if bitAt(o2[j], i) != bit {
				o2 = append(o2[:j], o2[j+1:]...)
				continue
			}
			j++
		}
	}

	// CO2 scrubber rating
	for i := 0; i < WIDTH; i++ {
		if len(co2) == 1 {
			break
		}

		// Get most common bit
		var n0, n1 int
		for j := 0; j < len(co2); j++ {
			if bitAt(co2[j], i) == 1 {
				n1++
			} else {
				n0++
			}
		}
		var bit int64
		if n1 < n0 {
			bit = 1
		}

		// Filter
		for j := 0; j < len(co2); {
			if bitAt(co2[j], i) != bit {
				co2 = append(co2[:j], co2[j+1:]...)
				continue
			}
			j++
		}
	}

	fmt.Printf("Oxygen generator: %d\n", o2[0])
	fmt.Printf("CO2 scrubber: %d\n", co2[0])
	fmt.Printf("Life support rating: %d\n", co2[0]*o2[0])
}
