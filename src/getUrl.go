package main

import (
	"github.com/phayes/freeport"
)

type url struct{
	scheme string
	domain string
	port int
}

func NewUrl() *url{
	p := new(url)
	p.scheme = "http://"
	p.domain = "localhost"
	p.setPort() // auto set free port
	return p
}

func (s *url)setPort(){
	var portErr error
	s.port, portErr = freeport.GetFreePort()
	if portErr != nil {
		panic(portErr)
	}
}