package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr      = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	queueName = flag.String("name", "my-test-queue", "queue name")
)

func main() {
	flag.Parse()

	log.Println("connecting to etcd:", *addr)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(*addr, ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	log.Println("creating priority queue:", *queueName)
	q := recipe.NewPriorityQueue(cli, *queueName)

	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "push":
			if len(items) != 3 {
				log.Println("must set value and priority to push")
				continue
			}
			pr, err := strconv.Atoi(items[2])
			if err != nil {
				log.Println("must set uint16 as priority")
				continue
			}
			q.Enqueue(items[1], uint16(pr))
		case "pop":
			v, err := q.Dequeue()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(v)
		case "quit", "exit":
			return
		default:
			log.Println("unknown action")
		}
	}
}
