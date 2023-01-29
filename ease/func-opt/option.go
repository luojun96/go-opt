package main

import "time"

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Conf.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Conf.Timeout = timeout
	}
}

func MaxConns(maxcount int) Option {
	return func(s *Server) {
		s.Conf.MaxCount = maxcount
	}
}

func TLS(tls *TLSConfig) Option {
	return func(s *Server) {
		s.Conf.TLS = tls
	}
}

func NewServerWithOptions(addr string, port int, options ...func(*Server)) (*Server, error) {
	conf := Config{
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxCount: 1000,
		TLS:      nil,
	}

	srv := Server{
		Addr: addr,
		Port: port,
		Conf: &conf,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}
