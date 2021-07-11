package ch4

import (
	"bufio"
	"fmt"
	"os"
)

// 读取多行输入，但是只打印第一次出现的行
func dedup()  {
	// 将这种忽略value的map当作字符串集合
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

