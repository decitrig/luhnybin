package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

var (
	digit         = regexp.MustCompile("^[0-9].*")
	validCardRune = regexp.MustCompile(`^([\-0-9 ]).*`)
)

func assert(cond bool, message string) {
	if !cond {
		panic(message)
	}
}

func debug(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// left <- first digit
// right <- left + 1
// increment right until 14th digit
//    if invalid char found, restart with next possible
// check right, right + 1 digit, right + 2 digits
// keep track of last character output, if check succeeds output from pos to right
// start again with next character after right

type CardMasker interface {
	Mask(line string) string
}

type IdentityMasker struct{}

func (masker IdentityMasker) Mask(line string) string {
	return line
}

type IterativeMasker struct{}

func (masker IterativeMasker) Mask(line string) string {
	src := utf8.NewString(line)
	output := make([]int, len(line))

	left := firstDigitAfter(line, 0)
	right := findNthDigitAfter(line, left, 14)
	nextOut := 0
	for nextOut < left {
		output[nextOut] = src.At(nextOut)
		nextOut++
	}
	assert(left == nextOut, "next output must be first digit")

	return string(output)
}

// Return the position of the first digit with a position >= start.
func firstDigitAfter(runes string, start int) int {
	pos := start
	for pos < len(runes) && !digit.MatchString(runes[pos:]) {
		pos++
	}
	return pos
}

// Return one past the position of the nth digit after start. Digits may be
// interspersed with spaces and dashes, but if an invalid card rune is found
// before n digits are consumed, return -1;
func findNthDigitAfter(runes string, start int, n int) int {
	pos := start
	for n > 1 {
		if !validCardRune.MatchString(runes[pos:]) {
			return -1
		}
		if digit.MatchString(runes[pos:]) {
			n--
		}
		pos++
	}
	return pos
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	var masker IterativeMasker
	for err == nil {
		fmt.Print(masker.Mask(line))
		line, err = reader.ReadString('\n')
	}
	if err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
