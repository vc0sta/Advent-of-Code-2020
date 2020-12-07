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
	seatList := stringToSeat(list)
	fmt.Printf("Part I: %d\n", partI(seatList))
	seatBool := seatToBool(seatList)
	fmt.Printf("Part II: %d\n", partII(seatBool))
}

type seat struct {
	row int64
	col int64
}

func stringToSeat(list []string) []seat {
	var seatList []seat

	for _, v := range list {

		str := strings.ReplaceAll(v, "R", "1")
		str = strings.ReplaceAll(str, "L", "0")
		str = strings.ReplaceAll(str, "B", "1")
		str = strings.ReplaceAll(str, "F", "0")

		row := str[0:7]
		col := str[7:]

		seat := seat{}

		if i, err := strconv.ParseInt(col, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			seat.col = i
		}
		if i, err := strconv.ParseInt(row, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			seat.row = i
		}

		seatList = append(seatList, seat)

	}

	return seatList
}

func partI(listSeat []seat) int64 {
	var bigger int64

	for _, seat := range listSeat {
		seatID := seat.row*8 + seat.col

		if bigger < seatID {
			bigger = seatID
		}

	}
	return bigger
}

func partII(list [128][8]bool) int {
	total := 1

	for row := 15; row < 120; row++ {
		for col := 0; col < 8; col++ {
			if !list[row][col] {
				return row*8 + col
			}
		}
	}

	return total
}

func seatToBool(list []seat) [128][8]bool {
	var plane [128][8]bool

	for _, v := range list {
		plane[v.row][v.col] = true
	}

	return plane
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
