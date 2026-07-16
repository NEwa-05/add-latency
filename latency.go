// Package latency
package addlatency

import (
	"context"
	"net/http"
	"time"
)

// Config set the needed plugin parameter.
type Config struct {
	AddedLatency int `json:"addedLatency,omitempty"`
}

// CreateConfig create the plugin config.
func CreateConfig() *Config {
	return &Config{
		AddedLatency: 3,
	}
}

// Latency plugin setup.
type Latency struct {
	next         http.Handler
	addedLatency int
	name         string
}

// New create the new latency plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Latency{
		addedLatency: config.AddedLatency,
		next:         next,
		name:         name,
	}, nil
}

func (a *Latency) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Duration(a.addedLatency) * time.Second)
	a.next.ServeHTTP(rw, req)
}
