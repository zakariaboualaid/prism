package backend

import (
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type BackendURL struct {
	*url.URL
}

// Backend hold data about a server
type Backend struct {
	URL          BackendURL `yaml:"url"`
	ReverseProxy *httputil.ReverseProxy
	Alive        bool
}

func (b *Backend) setAlive() (alive bool) {
	b.Alive = alive
	return
}

func (b *Backend) isAlive() (alive bool) {
	alive = b.Alive
	return
}

// Holds data about reachable backends
type ServerPool struct {
	Backends []*Backend
	Current  uint64
}

// AddBackend to server pool
func (s *ServerPool) AddBackend(backend *Backend) {
	s.Backends = append(s.Backends, backend)
}

func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex()
	l := len(s.Backends) + next
	for i := next; i < l; i++ {
		idx := i % len(s.Backends)
		if s.Backends[idx].isAlive() {
			if i != next {
				atomic.StoreUint64(&s.Current, uint64(idx))
			}
			return s.Backends[idx]
		}
	}
	return nil
}

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.Current, uint64(1)) % uint64(len(s.Backends)))
}

func (b *BackendURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	url, err := url.Parse(s)
	b.URL = url
	return nil
}
