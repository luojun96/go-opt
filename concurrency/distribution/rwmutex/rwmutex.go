package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr     = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	lockName = flag.String("name", "my-test-lock", "lock name")
	action   = flag.String("rw", "w", "r means acquiring read lock, w means acquiring write lock")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	log.Println("connecting to etcd")
	endpoints := strings.Split(*addr, ",")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	log.Println("creating etcd session")
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	log.Println("creating etcd lock")
	m1 := recipe.NewRWMutex(s1, *lockName)
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		switch action {
		case "w":
			testWriteLocker(m1)
		case "r":
			testReadLocker(m1)
		default:
			fmt.Println("unknown action")
		}
	}
}

func testWriteLocker(m1 *recipe.RWMutex) {
	log.Println("acquiring write lock")
	if err := m1.Lock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired write lock")

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if err := m1.Unlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released write lock")
}

func testReadLocker(m1 *recipe.RWMutex) {
	log.Println("acquiring read lock")
	if err := m1.RLock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired read lock")

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if err := m1.RUnlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released read lock")
}
