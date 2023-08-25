package geecache

import (
	"fmt"
	"geecache/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

const defaultBasePath = "/_geecache/"
const defaultReplicas = 50

type HTTPPool struct {
	self        string
	basePath    string
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter
	mu          sync.Mutex
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("httppool serving unexpected path" + r.URL.Path)
	}
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		log.Println("bad request")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Println("show part", parts)
	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)
	log.Println("show group", groupName, key, group)
	if group == nil {
		log.Println("no such group" + groupName)
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}
	view, err := group.Get(key)
	log.Println("res", view, err, key)
	if err != nil {
		log.Println("get err", view, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
	log.Println("end", view)
}

func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{
			baseUrl: peer + p.basePath,
		}
	}
}

func (p *HTTPPool) PickPeer(key string) (*httpGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		return p.httpGetters[peer], true
	}
	return nil, false
}

type httpGetter struct {
	baseUrl string
}

func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf("%v%v%v", h.baseUrl, group, key)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned :%v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body :%v", err)
	}
	return bytes, nil
}
