package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Calc(lineints []int) int {
	a, b := lineints[0], lineints[1]
	count := 0
	for n := a; n <= b; n++ {
		var o []int
		y := (len(strconv.Itoa(n)) - 1)
		for z := 0; z < y; z++ {
			x := int(math.Pow(10, float64(z+1)))
			expons := int(math.Pow(10, float64(y-z)))
			m := int(float64(n)/float64(x)) + n%int(x)*expons
			if int(m) > n && int(m) <= b {
				o = AppendIfMissing(o, int(m))
			}
		}
		count += len(o)
	}
	return count
}

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func Integerify(a string) []int {
	var returnarray []int
	arr := strings.Split(a, " ")
	for _, s := range arr {
		converted, _ := strconv.Atoi(s)
		returnarray = append(returnarray, converted)
	}
	return returnarray
}

func main() {
	argsWithoutProgram := os.Args[1]
	lines := GrabLines(argsWithoutProgram)[1:]
	for i, line := range lines {
		lineints := Integerify(line)
		fmt.Println("Case #", i+1, ":", Calc(lineints))
	}
}
