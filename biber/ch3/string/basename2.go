package string

import "strings"

// 使用了strings.LastIndex
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if not found
	s = s[slash + 1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
// 另:path和path/filepath包提供了关于文件路径名更一般的函数操作
