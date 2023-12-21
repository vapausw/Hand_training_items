package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() {
	out := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(out, "test")
}
func main() {
	run()
}
