package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var connMu sync.Mutex
var conn net.Conn
var once sync.Once
var addr = "baidu.com:80"

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	if conn != nil {
		return conn
	}

	conn, _ = net.DialTimeout("tcp", addr, 10*time.Second)
	return conn
}

func getConnByOnce() net.Conn {
	once.Do(func() {
		var err error
		fmt.Println("get connection by sync.Once")
		conn, err = net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			panic(err)
		}
	})
	return conn
}

func main() {

	fmt.Println("get connection")
	conn = getConnByOnce()
	if conn == nil {
		panic("conn is nil")
	}
	fmt.Println("get connection success")
	getConnByOnce()
	fmt.Printf("conn: %v\n", conn.RemoteAddr().String())
}
