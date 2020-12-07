package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	list := readInputs("input.txt")

	fmt.Printf("Part I: %d\n", partI(list))
	fmt.Printf("Part II: %d\n", partII(list))
}

type Password struct {
	min  int
	max  int
	pass string
	char string
}

func readInputs(input string) []Password {

	var list []Password

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		section := strings.Split(scanner.Text(), " ")

		times := strings.Split(section[0], "-")

		char := strings.Split(section[1], ":")

		min, _ := strconv.Atoi(times[0])
		max, _ := strconv.Atoi(times[1])

		pass := Password{
			min:  min,
			max:  max,
			pass: section[2],
			char: char[0]}

		list = append(list, pass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list

}

func partI(list []Password) int {
	total := 0
	for _, pass := range list {
		count := strings.Count(pass.pass, pass.char)

		if count >= pass.min && count <= pass.max {
			total += 1
		}

	}

	return total
}

func partII(list []Password) int {
	var first, last bool
	total := 0
	for _, pass := range list {
		first = pass.pass[pass.min-1] == []byte(pass.char)[0]
		last = pass.pass[pass.max-1] == []byte(pass.char)[0]

		if first != last {
			total += 1
		}
	}

	return total
}
