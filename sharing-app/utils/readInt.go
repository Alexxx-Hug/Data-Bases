package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadInt(reader *bufio.Reader) int64 {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	value, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Invalid number")
		return 0
	}

	return value
}
