package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr      = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	queueName = flag.String("name", "my-test-queue", "queue name")
)

func main() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")

	log.Println("connect to etcd:", endpoints)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	log.Println("create queue:", *queueName)
	q := recipe.NewQueue(cli, *queueName)

	log.Println("start console")
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "push":
			if len(items) != 2 {
				log.Println("must set value to push")
				continue
			}
			q.Enqueue(items[1])
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
