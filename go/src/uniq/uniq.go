package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	lines := make(map[string]int)

	for input.Scan() {
		lines[input.Text()]++
	}

	for line, num := range lines {
		if len(line) > 0 {
			fmt.Printf("%s \t%d\n", line, num)
		}
	}
}
