package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1)

	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2)
}

// var connMu sync.Mutex
// var conn net.Conn
//
// func getConn() net.Conn {
// 	connMu.Lock()
// 	defer connMu.Unlock()
//
// 	if conn != nil {
// 		return conn
// 	}
//
// 	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
// 	return conn
// }
//
// func main() {
//
// 	conn := getConn()
// 	if conn == nil {
// 		panic("conn is nil")
// 	}
// }
