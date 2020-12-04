package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type eyeColor string

const (
	AMB eyeColor = "amb"
	BLU          = "blu"
	BRN          = "brn"
	GRY          = "gry"
	GRN          = "grn"
	HZL          = "hzl"
	OTH          = "oth"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	var valid int = 0
	var hasCid bool = false
	var field int = 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if string(line) == "" {
			if field == 8 || (field == 7) && hasCid == false {
				valid++
			}
			hasCid = false
			field = 0
			continue
		}

		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			key := strings.Split(pair, ":")[0]
			value := strings.Split(pair, ":")[1]
			switch key {
			case "cid":
				hasCid = true
				field++
			case "byr":
				byr, _ := strconv.Atoi(value)
				if byr >= 1920 && byr <= 2002 {
					field++
				}
			case "iyr":
				iyr, _ := strconv.Atoi(value)
				if iyr >= 2010 && iyr <= 2020 {
					field++
				}
			case "eyr":
				eyr, _ := strconv.Atoi(value)
				if eyr >= 2020 && eyr <= 2030 {
					field++
				}
			case "hgt":
				unit := value[len(value)-2:]
				hgt, _ := strconv.Atoi(value[0 : len(value)-2])
				if unit == "cm" {
					if hgt >= 150 && hgt <= 193 {
						field++
					}
				} else if unit == "in" {
					if hgt >= 59 && hgt <= 76 {
						field++
					}
				}
			case "hcl":
				var isValid bool = true
				for i, char := range value {
					if i == 0 {
						if string(char) != "#" {
							isValid = false
							break
						} else {
							continue
						}
					}
					if int(char)-int('0') <= 9 || int(char)-int('a') <= 5 {
						continue
					} else {
						isValid = false
						break
					}
				}
				if isValid == true {
					field++
				}
			case "ecl":
				ecl := eyeColor(value)
				if err := ecl.isValidEcl(); err == nil {
					field++
				}
			case "pid":
				var isValid bool = true
				if len(value) != 9 {
					isValid = false
				}
				if isValid == true {
					for _, char := range value {
						if int(char)-int('0') > 9 {
							isValid = false
						}
					}
				}
				if isValid == true {
					field++
				}
			}
		}
	}

	if field == 8 || (field == 7) && hasCid == false {
		valid++
	}

	fmt.Printf("There are %d valid passport(s).", valid)
}

func (ec eyeColor) isValidEcl() error {
	switch ec {
	case AMB, BLU, BRN, GRY, GRN, HZL, OTH:
		return nil
	}
	return errors.New("invalid ecl")
}
