package string

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	// Buffer类型用于字节slice的缓存,可动态增长
	// 可用作一个I/O的输入和输出对象
	var buf bytes.Buffer // 不需初始化,零值也有效
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
