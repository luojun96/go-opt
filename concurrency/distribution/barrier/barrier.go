package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	barrierName = flag.String("name", "my-test-queue", "barrier name")
)

func main() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")
	fmt.Println("endpoints:", endpoints)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("new client error:", err)
	}
	defer cli.Close()

	fmt.Println("new client success", cli)
	b := recipe.NewBarrier(cli, *barrierName)

	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "hold":
			b.Hold()
			fmt.Println("hold")
		case "release":
			b.Release()
			fmt.Println("released")
		case "wait":
			b.Wait()
			fmt.Println("after wait")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
