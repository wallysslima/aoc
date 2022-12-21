package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	current, max := 0, 0

	for input.Scan() {
		snack, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Fprintf(os.Stdin, "err: %v", err)
			current = 0
			continue
		}
		current += snack
		if current > max {
			max = current
		}
	}

	fmt.Println(max)
}
