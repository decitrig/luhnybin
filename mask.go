package main

import (
    "fmt"
    "bufio"
    "os"
)

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
