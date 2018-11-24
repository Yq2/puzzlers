package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringsBuilder(b *testing.B) {
	str := strings.Builder{}
	for i := 0; i < b.N; i++ {
		str.WriteString("strings_builder_test")
		_ = str.String()
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	str := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		str.WriteString("bytes_buffer_test")
		_ = str.String()
	}
}

//测试证明 strings.Builder.String()方法比bytes.Buffer.String()更高效


