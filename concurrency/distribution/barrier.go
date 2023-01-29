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
	addr        = flag.String("addr", "https://192.168.31.23:2379", "etcd addresses")
	barrierName = flag.String("name", "my-test-queue", "barrier name")
)

func exec() {
	flag.Parse()

	endpoints := rstrings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

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
