package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	frequency := 0

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		diffString := scanner.Text()
		diffInt, err := strconv.Atoi(diffString)
		if err != nil {
			panic(err)
		}

		frequency += diffInt
	}

	fmt.Printf("Final frequency is %d\n", frequency)
}

func part2() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	frequencyChanges := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		diffString := scanner.Text()
		diffInt, err := strconv.Atoi(diffString)
		if err != nil {
			panic(err)
		}

		frequencyChanges = append(frequencyChanges, diffInt)
	}

	frequency := 0
	frequencyHistory := make(map[int]struct{}, 0)

	// initialize map with initial frequency (0)
	frequencyHistory[0] = struct{}{}
	for {
		for _, f := range frequencyChanges {
			frequency += f
			if _, ok := frequencyHistory[frequency]; ok {
				fmt.Printf("First frequency seen twice is %d\n", frequency)
				return
			}
			frequencyHistory[frequency] = struct{}{}
		}
	}
}
