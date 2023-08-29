package topicserver

import (
	"encoding/json"
	// "log"
	"testing"
)

func TestTopic(t *testing.T) {
	c := &topic{
		Title:     "test",
		Content:   "hello world",
		Creator:   "apple",
		Create_ts: 1693237249,
	}
	// json.Marshal()
	// var err error
	buf := make([]byte, 0)
	buf, _ = json.Marshal(c)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// log.Println(string(buf))
	t.Log(string(buf))
	// t.Log
}

// func Fib(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return Fib(n-1) + Fib(n-1)
// }
// func TestFib(t *testing.T) {
// 	var (
// 		in       = 7
// 		expected = 64
// 	)
// 	actual := Fib(in)
// 	if actual != expected {
// 		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
// 	}
// }

// func
