package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr        = flag.String("addr", "https://192.168.31.23:2379", "etcd address")
	barrierName = flag.String("name", "my-test-doublebarrier", "barrier name")
	count       = flag.Int("c", 2, "")
)

func exec() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	b := recipe.NewDoubleBarrier(s1, *barrierName, *count)

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
