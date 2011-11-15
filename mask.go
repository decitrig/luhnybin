package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

var (
    digit = regexp.MustCompile("^[0-9].*")
    validCardRune = regexp.MustCompile("[\\- 0-9]")
)

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

type IdentityMasker struct {}

func (masker IdentityMasker) Mask(line string) string {
    return line
}

type IterativeMasker struct {}

func (masker IterativeMasker) Mask(line string) string {
    slice := make([]int, len(line))
    for idx, rune := range line {
        slice[idx] = rune
    }
    return string(slice)
}

// Return the position of the first digit with a position >= start.
func firstDigitAfter(runes string, start int) int {
    pos := start
    for !digit.MatchString(runes[pos:]) {
        pos++;
    }
    return pos
}

func handleLine(line string, masker CardMasker) string {
    return masker.Mask(line)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    var masker IterativeMasker
    for err == nil {
        fmt.Print(handleLine(line, masker))
        line, err = reader.ReadString('\n')
    }
    if err != os.EOF {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err);
    }
}
