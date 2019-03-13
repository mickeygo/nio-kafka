package main

import "net/http"

// Host a server host
type Host struct {
	Address string
}

// BuildHost create a new host
func BuildHost() *Host {
	return &Host{}
}

// Start the host
func (h *Host) Start() {
	http.ListenAndServe(h.Address, nil)
}
