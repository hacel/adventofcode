package scan

import (
	"bufio"
	"os"
	"strconv"
)

func B(path string) chan []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	ch := make(chan []byte)
	go func() {
		defer f.Close()
		defer close(ch)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Bytes()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	return ch
}

func S(path string) chan string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	return ch
}

func I(path string) chan int {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	ch := make(chan int)
	go func() {
		defer f.Close()
		defer close(ch)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			ch <- i
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	return ch
}
