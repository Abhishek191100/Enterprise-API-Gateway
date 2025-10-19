package router

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Path     string   `yaml:"path"`
	PathType string   `yaml:"pathType"`
	Method   string   `yaml:"method"`
	Host     string   `yaml:"host,omitempty"`
	Backends []string `yaml:"backends"`
}

type RoutingTable struct {
	mu     sync.RWMutex
	Routes []*Route
}

// LoadRoutingTable parses the YAML config and creates the routing table.
func LoadRoutingTable(configFile string) (*RoutingTable, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data struct {
		Routes []*Route `yaml:"routes"`
	}
	dec := yaml.NewDecoder(file)
	if err = dec.Decode(&data); err != nil {
		return nil, err
	}
	return &RoutingTable{Routes: data.Routes}, nil
}

func (rt *RoutingTable) Match(r *http.Request) (*Route, error) {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	method := r.Method
	path := r.URL.Path
	host := r.Host

	for _, route := range rt.Routes {
		if route.PathType == "exact" && route.Method == method && route.Path == path && (route.Host == host || route.Host == "") {
			return route, nil
		}
	}
	for _, route := range rt.Routes {
		if route.PathType == "prefix" && route.Method == method && strings.HasPrefix(path, route.Path) && (route.Host == host || route.Host == "") {
			return route, nil
		}
	}

	return nil, errors.New("no matching route found")
}

func (r *Route) NextBackend() string {
	if len(r.Backends) == 0 {
		return ""
	}
	return r.Backends[0]
}
