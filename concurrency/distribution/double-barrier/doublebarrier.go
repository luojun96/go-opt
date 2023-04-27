package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd address")
	barrierName = flag.String("name", "my-test-doublebarrier", "barrier name")
	count       = flag.Int("c", 2, "")
)

func main() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	fmt.Println("connected to etcd")

	// create a session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		fmt.Println(err)
	}
	defer s1.Close()
	fmt.Println("created session")

	b := recipe.NewDoubleBarrier(s1, *barrierName, *count)
	fmt.Println("created double barrier")

	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "enter":
			b.Enter()
			fmt.Println("enter")
		case "leave":
			b.Leave()
			fmt.Println("leave")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
