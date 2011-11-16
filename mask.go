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
	fmt.Fprintf(os.Stdout, format+"\n", args...)
}

func isDigit(r rune) bool {
    return (r >= '0') && (r <= '9')
}

func isValidCardRune(r rune) bool {
    return isDigit(r) || r == ' ' || r == '-'
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

func (buffer RuneBuffer) firstDigit() int {
    for start := 0; start < len(buffer.source); start++ {
        if isDigit(buffer.source[start]) {
            return start;
        }
    }
    return -1;
}

func (buffer RuneBuffer) tryToMask(start, n int) {
    if start < 0 {
        return
    }
    sum := 0
    digitsFound := 0
    toMask := make([]int, 0, n)
    for pos := start; pos < len(buffer.source); pos++ {
        r := buffer.source[pos]
        if !isValidCardRune(r) {
            return
        }
        if isDigit(r) {
            digitsFound++
            toMask = append(toMask, pos)
            digit := r - '0'
            if ((n - digitsFound) % 2 == 0) {
                sum += sumDigits(digit)
            } else {
                sum += sumDigits(digit * 2)
            }
        }
        if digitsFound == n {
            break;
        }
    }
    if (sum % 10) != 0 {
        return
    }
    for _, idx := range toMask {
        buffer.output[idx] = 'X'
    }
}

func (buffer RuneBuffer) size() int {
    return len(buffer.source)
}

func mask(line string) string {
    output := NewRuneBuffer(line)
    left := output.firstDigit()
    for left < output.size() {
        for n := 14; n <= 16; n++ {
            output.tryToMask(left, n)
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
