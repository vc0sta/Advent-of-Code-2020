package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strconv"
)

func main() {

	list := readInputs("input.txt")

	fmt.Printf("Part I: %d\n", partI(list))
	fmt.Printf("Part II: %d\n", partII(list))
}

func readInputs(input string) []int {

	var list []int

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		number, _ := strconv.Atoi(scanner.Text())
		list = append(list, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func partI(list []int) int {

	for _, x := range list {
		for _, y := range list {
			if x+y == 2020 {
				return x * y
			}

		}
	}

	return 0
}

func partII(list []int) int {

	for _, x := range list {
		for _, y := range list {
			for _, z := range list {
				if x+y+z == 2020 {
					return x * y * z
				}

			}

		}
	}
	return 0

}
