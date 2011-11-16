package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func assert(cond bool, message string) {
	if !cond {
		panic(message)
	}
}

func debug(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func isDigit(r rune) bool {
    return (r >= '0') && (r <= '9')
}

func isValidCardRune(r rune) bool {
    return isDigit(r) || r == ' ' || r == '-'
}

func isLunhy(digits []int) bool {
    sum := 0
    for offset := 1; offset <= len(digits); offset++ {
        pos := len(digits) - offset;
        if offset % 2 == 0 {
            digits[pos] *= 2
        }
        sum += sumDigits(digits[pos])
    }
    return (sum % 10) == 0
}

func sumDigits(n int) (sum int) {
    sum = 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return
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

func (buffer RuneBuffer) String() string {
    return string(buffer.output)
}

func (buffer *RuneBuffer) MaskDigits(left, right int) {
    for left < right {
        r := buffer.source[left]
        if isDigit(r) {
            buffer.output[left] = 'X'
        }
        left++
    }
}

func (buffer *RuneBuffer) GetDigits(left, right int) (digits []int) {
    for left < right {
        r := buffer.source[left]
        if isDigit(r) {
            digits = append(digits, r - '0')
        }
        left++
    }
    return
}

func (buffer RuneBuffer) firstDigit() int {
    for start := 0; start < len(buffer.source); start++ {
        if isDigit(buffer.source[start]) {
            return start;
        }
    }
    return -1;
}

func (buffer RuneBuffer) findNthDigitAfter(start, n int) (pos int, ok bool) {
	pos = start
    ok = false
    if pos < 0 {
        return
    }
	for pos < len(buffer.source) && n > 0 {
        r := buffer.source[pos]
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

func (buffer RuneBuffer) size() int {
    return len(buffer.source)
}

func mask(line string) string {
    output := NewRuneBuffer(line)
    left := output.firstDigit()
    for left < output.size() {
        for n := 14; n <= 16; n++ {
            right, ok := output.findNthDigitAfter(left, n)
            if ok {
                digits := output.GetDigits(left, right)
                if isLunhy(digits) {
                    output.MaskDigits(left, right)
                }
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Print(mask(line))
		line, err = reader.ReadString('\n')
	}
	if err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
