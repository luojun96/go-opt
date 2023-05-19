package main

import (
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func initLog() {
	logger, _ = zap.NewProduction()
}

func main() {
	var (
		server *Server
		err    error
	)

	initLog()
	logger.Info("create server without options")
	server, err = NewServer("localhost", 9000, nil)
	if err != nil {
		logger.Error("create server failed", zap.Error(err))
	}
	logger.Info("server without options created", zap.String("addr", server.Addr))

	logger.Info("create server with options")
	conf := Config{
		Protocol: "tcp",
		Timeout:  time.Duration(60),
	}
	server, err = NewServer("localhost", 9000, &conf)
	if err != nil {
		logger.Error("create server failed", zap.Error(err))
	}
	logger.Info("server with options created", zap.String("addr", server.Addr), zap.Any("config", conf))

	logger.Info("create server with options using functional options")

	logger.Info("options is nil")
	server, err = NewServerWithOptions("localhost", 1024)
	if err != nil {
		logger.Error("create server failed", zap.Error(err))
	}

	logger.Info("set protocol to udp")
	server, err = NewServerWithOptions("localhost", 2048, Protocol("udp"))
	if err != nil {
		logger.Error("create server failed", zap.Error(err))
	}
	logger.Info("set protocol to udp", zap.Any("config", server.Config))

	logger.Info("set timeout to 30 seconds, and max connections to 10000")
	server, err = NewServerWithOptions("localhost", 8080, Timeout(300*time.Second), MaxConns(10000))
	if err != nil {
		logger.Error("create server failed", zap.Error(err))
	}
	logger.Info("set timeout to 30 seconds, and max connections to 10000", zap.Any("config", server.Config))
}
