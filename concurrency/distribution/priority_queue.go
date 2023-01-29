package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr      = flag.String("addr", "https://192.168.31.23:2379", "etcd addresses")
	queueName = flag.String("name", "my-test-queue", "queue name")
)

func exec() {
	flag.Parse()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(*addr, ","),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	q := recipe.NewPriorityQueue(cli, *queueName)

	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "push":
			if len(items) != 3 {
				fmt.Println("must set value and priority to push")
				continue
			}
			pr, err := strconv.Atoi(items[2])
			if err != nil {
				fmt.Println("must set uint16 as priority")
				continue
			}
			q.Enqueue(items[1], unit16(pr))
		case "pop":
			v, err := q.Dequeue()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(v)
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
