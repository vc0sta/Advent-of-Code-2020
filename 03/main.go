package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
)

func main() {

	list := readInputs("input.txt")

	fmt.Printf("Part I: %d\n", partI(list))
	fmt.Printf("Part II: %d\n", partII(list))
}

func partI(list []string) int {
	return countTrees(list, 3, 1)
}

func partII(list []string) int {
	total := 1
	total *= countTrees(list, 1, 1)
	total *= countTrees(list, 3, 1)
	total *= countTrees(list, 5, 1)
	total *= countTrees(list, 7, 1)
	total *= countTrees(list, 1, 2)
	return total
}

func readInputs(input string) []string {

	var list []string

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		section := scanner.Text()
		list = append(list, section)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list

}

func countTrees(list []string, right, down int) int {
	total := 0
	position := 0
	for index, section := range list {
		if index%down == 0 {

			if isTree(section[position]) {
				total += 1
			}
			position = (position + right) % len(section)
		}
	}
	return total
}

func isTree(char byte) bool {
	return char == []byte("#")[0]
}
