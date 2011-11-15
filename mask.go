package main

import (
    "fmt"
    "bufio"
    "os"
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

func handleLine(line string, masker CardMasker) string {
    return masker.Mask(line)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    var masker IdentityMasker
    for err == nil {
        fmt.Print(handleLine(line, masker))
        line, err = reader.ReadString('\n')
    }
    if err != os.EOF {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err);
    }
}
