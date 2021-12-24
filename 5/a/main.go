package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hacel/adventofcode/internal/scan"
)

type Vent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	vents := []*Vent{}
	width := 0
	height := 0
	for s := range scan.S("5/input.txt") {
		fields := strings.FieldsFunc(s, func(r rune) bool {
			if r == ',' || r == ' ' || r == '-' || r == '>' {
				return true
			}
			return false
		})
		x1, _ := strconv.Atoi(fields[0])
		if x1 > width {
			width = x1
		}
		y1, _ := strconv.Atoi(fields[1])
		if y1 > height {
			height = y1
		}
		x2, _ := strconv.Atoi(fields[2])
		if x2 > width {
			width = x1
		}
		y2, _ := strconv.Atoi(fields[3])
		if y2 > height {
			height = y2
		}
		if x1 == x2 || y1 == y2 {
			vents = append(vents, &Vent{x1, y1, x2, y2})
		}
	}

	diagram := map[int]map[int]int{}
	for _, v := range vents {
		for v.y1 != v.y2 || v.x1 != v.x2 {
			if _, ok := diagram[v.y1]; !ok {
				diagram[v.y1] = make(map[int]int)
			}
			diagram[v.y1][v.x1]++
			if v.y1 < v.y2 {
				v.y1++
			} else if v.y1 > v.y2 {
				v.y1--
			}
			if v.x1 < v.x2 {
				v.x1++
			} else if v.x1 > v.x2 {
				v.x1--
			}
		}
		if _, ok := diagram[v.y1]; !ok {
			diagram[v.y1] = make(map[int]int)
		}
		diagram[v.y1][v.x1]++
	}

	overlap := 0
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if diagram[y][x] > 1 {
				overlap++
			}
		}
	}

	fmt.Printf("Overlap: %v\n", overlap)
}
