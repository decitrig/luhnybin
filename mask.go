package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

func isDigit(r rune) bool {
    return (r >= '0') && (r <= '9')
}

func isValidCardRune(r rune) bool {
    return isDigit(r) || r == ' ' || r == '-'
}


type RuneBuffer struct {
    source []rune
    output []rune
}

func NewRuneBuffer(s string) RuneBuffer {
    source := make([]rune, len(s))
    for i, r := range s {
        source[i] = r
    }
    output := make([]rune, len(s))
    copy(output, source)
    return RuneBuffer{source, output}
}

func (runeString RuneBuffer) String() string {
    return string(runeString.output)
}

func (runeString *RuneBuffer) MaskDigits(left, right int) {
    for left <= right {
        r := runeString.source[left]
        if isDigit(r) {
            runeString.output[left] = 'X'
        }
        left++
    }
}

func (runeString RuneBuffer) firstDigitAfter(start int) int {
    for start < len(runeString.source) {
        if isDigit(runeString.source[start]) {
            return start;
        }
        start++
    }
    return -1;
}

func (runeString RuneBuffer) findNthDigitAfter(start, n int) (pos int, ok bool) {
	pos = start
    ok = false
    if pos < 0 {
        return
    }
	for pos < len(runeString.source) && n > 1 {
        r := runeString.source[pos]
        if !isValidCardRune(r) {
            return
        }
        if isDigit(r) {
            n--
		}
		pos++
	}
    ok = true
	return
}

func (runeString RuneBuffer) size() int {
    return len(runeString.source)
}

type IterativeMasker struct{}

func (masker IterativeMasker) Mask(line string) string {
    output := NewRuneBuffer(line)
    left := output.firstDigitAfter(0)
    for left < output.size() {
        for n := 14; n <= 16; n++ {
            right, ok := output.findNthDigitAfter(left, n)
            if ok {
                output.MaskDigits(left, right)
            } else {
                // skip a short run of digits
                left = right
                break;
            }
        }
        left++
    }
	return output.String()
}

func maskDigits(runes []rune) {
    for pos, r := range runes {
        if digit.MatchString(string(r)) {
            runes[pos] = 'X'
        }
    }
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
