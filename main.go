package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(nums[0])
	m, _ := strconv.Atoi(nums[1])

	for ii := 1; ii <= n; ii++ {
		fmt.Printf("%s", strings.Repeat("X", m))
		fmt.Printf("%s", strings.Repeat(".", m))
		fmt.Printf("%s\n", strings.Repeat("X", m))
	}
	for ii := 1; ii <= n; ii++ {
		fmt.Printf("%s", strings.Repeat(".", m))
		fmt.Printf("%s", strings.Repeat("X", m))
		fmt.Printf("%s\n", strings.Repeat(".", m))
	}
	for ii := 1; ii <= n; ii++ {
		fmt.Printf("%s", strings.Repeat("X", m))
		fmt.Printf("%s", strings.Repeat(".", m))
		fmt.Printf("%s\n", strings.Repeat("X", m))
	}

}
