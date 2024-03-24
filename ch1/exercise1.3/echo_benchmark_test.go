// To run the benchmarks, execute the following command:
// go test -bench=.
// The output will be similar to:
// goos: darwin
// goarch: amd64
// pkg: github.com/striversity/gotraining/ch1/exercise1.3
// BenchmarkEcho1-4    10000000               134 ns/op
// BenchmarkEcho2-4    10000000               134 ns/op
// BenchmarkEcho3-4    10000000               134 ns/op
// PASS
// ok      github.com/striversity/gotraining/ch1/exercise1.3  4.013s

package main

import (
	"os"
	"strings"
	"testing"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

func echo2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

func echo3() {
	strings.Join(os.Args[1:], " ")
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
