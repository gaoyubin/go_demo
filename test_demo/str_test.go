package testdemo

import (
	"bytes"
	"math/rand"
	"testing"
	// "time"
)

// import "strconv"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// func TestFib2(t *testing.T) {
// 	// var (
// 	// 	in       = 7
// 	// 	expected = 12
// 	// )
// 	var fibTests = []struct {
// 		in       int
// 		expected int
// 	}{
// 		{1, 1},
// 		{2, 1},
// 		{3, 2},
// 		{4, 3},
// 		{5, 15},
// 		{6, 8},
// 		{7, 13},
// 	}

// 	for _, tt := range fibTests {
// 		actual := Fib(tt.in)
// 		if actual != tt.expected {
// 			t.Errorf("fib(%d)=%d,expected %d", tt.in, actual, tt.expected)
// 			// t.Fatal("fib(%d)=%d,expected %d", tt.in, actual, tt.expected)
// 			// strconv.

// 			// t.Fatal("fail" + strconv.Itoa(tt.in) + ":" + strconv.Itoa(tt.expected) + ":" + strconv.Itoa(actual))
// 		}
// 	}

// }

// func BenchmarkFib10(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Fib(10)
// 	}
// }

func BenchmarkFib(b *testing.B) {
	// time.Sleep(1 * time.Second)
	// b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Fib(30) // run fib(30) b.N times
	}
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}
func bufferConcat(n int, str string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()

}
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func benchmark(b *testing.B, f func(n int, str string) string) {
	// for i:=0; i<b.N; i++{
	// var
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(1000, str)
	}
	// }
}
func BenchmarkPlusConcat(b *testing.B) {
	benchmark(b, plusConcat)
}
func BenchmarkBufferConcat(b *testing.B) {
	benchmark(b, bufferConcat)
}
func BenchmarkByteConcat(b *testing.B) {
	benchmark(b, byteConcat)
}
