package registry

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type GeeRegistry struct {
	mu      sync.Mutex
	servers map[string]*ServerItem
	timeout time.Duration
}
type ServerItem struct {
	Addr  string
	start time.Time
}

var DefaultGeeRegister = New(defaultTimeout)

const (
	defaultTimeout = time.Minute * 5
)

func New(timeout time.Duration) *GeeRegistry {
	return &GeeRegistry{
		servers: make(map[string]*ServerItem),
		timeout: timeout,
	}
}

// ServeHTTP(w http.ResponseWriter, r *http.Request)
func (r *GeeRegistry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set("X-Geerpc-Servers", strings.Join(r.GetAliveServers(), ","))
		log.Println("show servers", r.GetAliveServers())
	case "POST":
		addr := req.Header.Get("X-Geerpc-Server")
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.PutServer(addr)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// r.
	}
}
func (r *GeeRegistry) PutServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.servers[addr]
	if s == nil {
		r.servers[addr] = &ServerItem{
			Addr:  addr,
			start: time.Now(),
		}
	} else {
		s.start = time.Now()
	}
}
func (r *GeeRegistry) GetAliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for _, s := range r.servers {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, s.Addr)
		} else {
			delete(r.servers, s.Addr)
		}
	}
	// alive = append(alive, "127.0.0.1:7777")
	sort.Strings(alive)
	return alive
}

func HeartBeat(registry, addr string, duration time.Duration) {
	if duration == 0 {
		duration = defaultTimeout - time.Duration(1)*time.Minute
	}
	var err error
	err = sendHeartBeat(registry, addr)
	go func() {
		t := time.NewTicker(duration)
		for err == nil {
			<-t.C
			err = sendHeartBeat(registry, addr)
		}
	}()
}

func sendHeartBeat(registry, addr string) error {
	log.Println(addr, "send heart beat to re", registry, addr)
	req, _ := http.NewRequest("POST", registry, nil)
	req.Header.Set("X-Geerpc-Server", addr)
	httpclient := &http.Client{}
	if _, err := httpclient.Do(req); err != nil {
		log.Println("rpc server err:", err)
		return err
	}
	return nil
}
