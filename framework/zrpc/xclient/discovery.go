package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

type selectMode int

const (
	RandomSelect selectMode = iota
	RoundRobinSelect
)

type Discovery interface {
	Refresh() error //refresh from remote registry
	Update(servers []string) error
	Get(mode selectMode) (string, error)
	GetAll() ([]string, error)
}

// MultiServerDiscovery is a discovery for multi servers without a registry center
// user provides the server addresses explicitly instead
type MultiServerDiscovery struct {
	r       *rand.Rand
	mu      sync.Mutex
	servers []string
	index   int //record the selected position for robin algorithm
}

func NewMultiServerDiscovery(servers []string) *MultiServerDiscovery {
	d := &MultiServerDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().Unix())),
	}
	d.index = d.r.Intn(math.MaxInt32 - 1)
	return d
}

func (d *MultiServerDiscovery) Refresh() error {
	return nil
}

func (d *MultiServerDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

func (d *MultiServerDiscovery) Get(mode selectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discover: no available servers")
	}
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		s := d.servers[d.index%n]
		d.index = (d.index + 1) & n
		return s, nil

	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
	return "", nil
}

func (d *MultiServerDiscovery) GetAll() ([]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	allServers := make([]string, len(d.servers), len(d.servers))
	copy(allServers, d.servers)
	//allServers = d.servers
	return allServers, nil
}
