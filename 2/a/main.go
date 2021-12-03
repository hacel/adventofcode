package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hacel/adventofcode/internal/scan"
)

func main() {
	x := 0
	y := 0
	for s := range scan.S("2/input.txt") {
		split := strings.Split(s, " ")
		direction := split[0]
		X, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "down":
			y += X
		case "up":
			y -= X
		case "forward":
			x += X
		}
	}
	fmt.Println(x * y)
}
