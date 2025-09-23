package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Please, enter numbers separated by spaces: ")
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}

		num, numErr := strconv.Atoi(n)
		if numErr != nil {
			fmt.Println(numErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println(inputErr)
		}
	}
	fmt.Println(sum)
}
