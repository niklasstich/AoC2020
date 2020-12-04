package AoC2020

import (
    "reflect"
    "regexp"
    "strconv"
    "strings"
)

type passport struct {
    Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

func Day4_1() (validPassports int) {
    lines, err := ReadLinesSeperatedByEmptyLines("../inputs/04.txt")
    if err != nil {
        panic(err)
    }
    for _, line := range lines {
        if checkPassportValidity(line, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, nil) {
            validPassports++
        }
    }
    return
}

func Day4_2() (validPassports int) {
    lines, err := ReadLinesSeperatedByEmptyLines("../inputs/04.txt")
    if err != nil {
        panic(err)
    }
    validator := map[string]func(string) bool{
        "byr": func(s string) bool {
            year, err := strconv.Atoi(s)
            if err != nil {
                return false
            }
            return year >= 1920 && year <= 2002
        },
        "iyr": func(s string) bool {
            year, err := strconv.Atoi(s)
            if err != nil {
                return false
            }
            return year >= 2010 && year <= 2020
        },
        "eyr": func(s string) bool {
            year, err := strconv.Atoi(s)
            if err != nil {
                return false
            }
            return year >= 2020 && year <= 2030
        },
        "hgt": func(s string) bool {
            if strings.HasSuffix(s, "in"){
                s = strings.Trim(s, "in")
                h, err := strconv.Atoi(s)
                if err != nil {
                    return false
                }
                return h >= 59 && h <= 76
            } else if strings.HasSuffix(s, "cm") {
                s = strings.Trim(s, "cm")
                h, err := strconv.Atoi(s)
                if err != nil {
                    return false
                }
                return h >= 150 && h <= 193
            } else {
                return false
            }
        },
        "hcl": func(s string) bool {
            regex, err := regexp.Compile("^#[0-9a-f]{6}$")
            if err != nil {
                panic(err)
            }
            return regex.MatchString(s)
        },
        "ecl": func(s string) bool {
            allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
            return Contains(allowed, s)
        },
        "pid": func(s string) bool {
            regex, err := regexp.Compile("^[0-9]{9}$")
            if err != nil {
                panic(err)
            }
            return regex.MatchString(s)
        },
        "cid": func(s string) bool {
           return true
        },
    }
    for _, line := range lines {
        if checkPassportValidity(line, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, validator) {
            validPassports++
        }
    }
    return
}

func checkPassportValidity(passp string, requiredFields []string, validation map[string]func(string) bool) bool {
    // get reflection value of an instance of the struct
    s := reflect.ValueOf(&passport{}).Elem()
    if s.Kind() != reflect.Struct {
        panic("s isn't reflect.Struct")
    }
    // set all found values in passport struct
    pairs := strings.Split(passp, " ")
    for _, pair := range pairs {
        tuple := strings.Split(pair, ":")
        s.FieldByName(strings.Title(tuple[0])).SetString(tuple[1])
    }

    for _, reqField := range requiredFields {
        field := s.FieldByName(strings.Title(reqField))
        if field.IsZero() {
            return false
        }
        if validation != nil && !validation[reqField](field.String()) {
            return false
        }
    }
    return true
}