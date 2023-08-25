package xclient

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SelectMode int

const (
	RandomSelect SelectMode = iota
	RoundRobinSelect
)

type GeeRegistryDiscovery struct {
	registry   string
	servers    []string
	index      int
	mu         sync.Mutex
	r          *rand.Rand
	lastUpdate time.Time
}

func NewGeeRegistryDiscovery(registry string) (*GeeRegistryDiscovery, error) {
	gee := &GeeRegistryDiscovery{
		registry: registry,
		servers:  make([]string, 0, 50),
		index:    0,
		r:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return gee, nil
}
func (d *GeeRegistryDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	d.lastUpdate = time.Now()
	return nil
}

func (d *GeeRegistryDiscovery) Refresh() error {
	// d.mu.Lock()
	// defer d.mu.Unlock()
	resp, err := http.Get(d.registry)
	if err != nil {
		log.Println("http get err", err)
		return err
	}
	server_str := resp.Header.Get("X-Geerpc-Servers")
	servers := strings.Split(server_str, ",")
	log.Println("show server", servers)

	d.servers = make([]string, 0, len(servers))
	for _, server := range servers {
		if strings.TrimSpace(server) != "" {
			d.servers = append(d.servers, strings.TrimSpace(server))
		}
	}
	d.lastUpdate = time.Now()
	log.Println("end", d.servers, d.lastUpdate)
	return nil
}
func (d *GeeRegistryDiscovery) Get(mode SelectMode) (string, error) {
	if err := d.Refresh(); err != nil {
		return "", err
	}
	n := len(d.servers)
	log.Println(n, mode)
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		s := d.servers[d.index%n]
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("get err")
	}

}
