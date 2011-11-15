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

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    for err == nil {
        fmt.Print(line)
        line, err = reader.ReadString('\n')
    }
    if err != os.EOF {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err);
    }
}
