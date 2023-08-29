package topicserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

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

func TestHandlePost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/topic/", HandlerTopic)
	p := &topic{
		Title:     "test",
		Content:   "hello world",
		Creator:   "apple",
		Create_ts: 1693237249,
	}
	by, _ := json.Marshal(p)
	buf := bytes.NewBuffer(by)
	r, _ := http.NewRequest(http.MethodPost, "/topic/", buf)

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	resp := w.Result()
	// resp.Body
	// resp.body
	rsp_body, _ := ioutil.ReadAll(resp.Body)
	index, _ := strconv.Atoi(string(rsp_body))
	t.Log(rsp_body, index)
	// resp.body
	if resp.StatusCode != http.StatusOK {
		t.Errorf("response code is %v", resp.StatusCode)
	}

	r2, _ := http.NewRequest(http.MethodGet, "/topic/"+string(rsp_body), nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	resp := w.Result()

	if resp.

}
