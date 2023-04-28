package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	nodeID    = flag.Int("id", 0, "node ID")
	addr      = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	electName = flag.String("name", "my-test-elect", "election name")
)

var count int

func main() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")

	log.Println("endpoints:", endpoints)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	log.Println("connected to etcd")

	log.Println("start to create session")
	session, err := concurrency.NewSession(cli)
	defer session.Close()

	log.Println("start to elect for ID:", *nodeID)
	e1 := concurrency.NewElection(session, *electName)

	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		switch action {
		case "elect":
			go elect(e1)
		case "proclaim":
			proclaim(e1)
		case "resign":
			resign(e1)
		case "watch":
			go watch(e1)
		case "query":
			query(e1)
		case "rev":
			rev(e1)
		default:
			log.Println("unknown action")
		}
	}
}

func elect(e1 *concurrency.Election) {
	log.Println("acampaigning for ID:", *nodeID)

	if err := e1.Campaign(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

func proclaim(e1 *concurrency.Election) {
	log.Println("proclaiming for ID:", *nodeID)
	if err := e1.Proclaim(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("proclaimed for ID:", *nodeID)
}

func resign(e1 *concurrency.Election) {
	log.Println("resigning for ID:", *nodeID)
	if err := e1.Resign(context.Background()); err != nil {
		log.Println(err)
	}
	log.Println("resigned for ID:", *nodeID)
}

func query(e1 *concurrency.Election) {
	resp, err := e1.Leader(context.Background())
	if err != nil {
		log.Printf("failed to get the current leader: %v", err)
	}
	log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
}

func rev(e1 *concurrency.Election) {
	rev := e1.Rev()
	log.Println("current rev:", rev)
}

func watch(e1 *concurrency.Election) {
	ch := e1.Observe(context.TODO())
	log.Println("start to watch for ID:", *nodeID)
	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader changed to", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}
