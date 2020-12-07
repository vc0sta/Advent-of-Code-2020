package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"regexp"
	// "strconv"
)

func main() {
	list := readInputs("input.txt")
	total, valid := partI(list)
	fmt.Printf("Part I: %d\n", total)
	fmt.Printf("Part II: %d\n", partII(valid))
}

func partI(list [][]string) (int, [][]string) {
	var valid [][]string
	total := 0
	for _, doc := range list {
		a := strings.Join(doc, " ")
		cid, _ := regexp.MatchString(`cid:`, a)

		if (!cid && len(doc) == 7) || len(doc) == 8 {
			total += 1
			valid = append(valid, doc)
		}
	}
	return total, valid
}

func partII(list [][]string) int {
	total := 0
	for _, doc := range list {
		passport := mapDocument(doc)
		if isValid(passport) {
			total += 1
		}

	}

	return total
}

func isValid(passport map[string]string) bool {
	byr, _ := strconv.Atoi(passport["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr, _ := strconv.Atoi(passport["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, _ := strconv.Atoi(passport["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	if !validHeight(passport["hgt"]) {
		return false
	}

	hcl, _ := regexp.Match("^#(?:[0-9a-fA-F]{3}){1,2}$", []byte(passport["hcl"]))
	if !hcl {
		return false
	}

	pid, _ := regexp.Match("^[0-9]{9}$", []byte(passport["pid"]))
	if !pid {
		return false
	}

	valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range valid {
		if color == passport["ecl"] {
			return true
		}
	}

	return false
}

func validHeight(value string) bool {
	in, _ := regexp.MatchString(`in`, value)
	cm, _ := regexp.MatchString(`cm`, value)

	if cm {
		size, _ := strconv.Atoi(value[:len(value)-2])
		if size < 150 || size > 193 {
			return false
		}
	} else if in {
		size, _ := strconv.Atoi(value[:len(value)-2])
		if size < 59 || size > 76 {
			return false
		}
	} else {
		return false
	}
	return true
}

func mapDocument(document []string) map[string]string {
	fields := make(map[string]string)

	for _, v := range document {
		values := strings.Split(v, ":")
		fields[values[0]] = values[1]
	}
	return fields
}

func readInputs(input string) [][]string {

	var list [][]string

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	doc := []string{}
	for scanner.Scan() {

		section := scanner.Text()
		if section == "" {
			list = append(list, doc)
			doc = []string{}
		} else {
			section := splitLines(section)
			for _, v := range section {

				doc = append(doc, v)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list

}

func splitLines(line string) []string {
	return strings.Split(line, " ")
}
