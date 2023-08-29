package topicserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func HandlerTopic(w http.ResponseWriter, r *http.Request) {
	// mp := topicserver.map_topic
	// log.Println(mp)
	// var err error
	switch r.Method {
	case http.MethodGet:
		HandlerGet(w, r)
	case http.MethodPost:
		HandlerPost(w, r)
	case http.MethodDelete:
		HandlerDelete(w, r)
	default:
		http.Error(w, "method err:"+r.Method, http.StatusBadRequest)
		return
	}
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {

	index, err := strconv.Atoi(r.URL.Path[len("/topic/"):])
	if err != nil {
		http.Error(w, "path err "+r.URL.Path, http.StatusBadRequest)
		return
	}

	t, ok := map_topic[uint64(index)]
	if !ok {
		http.Error(w, "index not find"+r.URL.Path, http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "marshal fail", http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	// log.Println(buf)
	if err != nil {
		http.Error(w, "body err", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var t topic
	json.Unmarshal(buf, &t)
	uniq_id++
	t.Id = uniq_id
	// var t topic
	map_topic[t.Id] = &t
	log.Println(t)
	w.Write([]byte(strconv.Itoa(int(t.Id))))
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(r.URL.Path[len("/topic/"):])
	if err != nil {
		http.Error(w, "path err "+r.URL.Path, http.StatusBadRequest)
		return
	}
	t, ok := map_topic[uint64(index)]
	if !ok {
		http.Error(w, "index not find"+r.URL.Path, http.StatusBadRequest)
		return
	}
	delete(map_topic, uint64(index))
	log.Println("delete succ", t)
}
