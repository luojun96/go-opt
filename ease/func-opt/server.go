package main

import (
	"crypto/x509"
	"errors"
	"time"
)

type TLSConfig struct {
	Host               string         `yaml:"host"`
	Proxy              string         `yaml:"proxy"`
	InsecureSkipVerify bool           `yaml:"insecure_skip_verify"`
	RootCAPemPath      string         `yaml:"root_ca_pem_path"`
	RootCaPem          string         `yaml:"root_ca_pem"`
	RootCAs            *x509.CertPool `yaml:"root_c_as"`
	ExpireSkipVerify   bool           `yaml:"expire_skip_verify"`
	AlertExpireBefore  time.Duration  `yaml:"alert_expire_before"`
}

type Config struct {
	Protocol string        `yaml:"protocol"`
	Timeout  time.Duration `yaml:"timeout"`
	MaxCount int           `yaml:"max_count"`
	TLS      *TLSConfig    `yaml:"tls"`
}

type Server struct {
	Addr string  `yaml:"addr"`
	Port int     `yaml:"port"`
	Conf *Config `yaml:config`
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	if len(addr) == 0 {
		return nil, errors.New("address should not be empty")
	}
	if port == 0 {
		return nil, errors.New("port should not be empty")
	}
	return &Server{
		Addr: addr,
		Port: port,
		Conf: conf,
	}, nil
}
