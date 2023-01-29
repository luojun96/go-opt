package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr     = flag.String("addr", "https://192.168.31.23:2379", "etcd addresses")
	lockName = flag.String("name", "my-test-lock", "lock name")
)

func exec() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	endpoints := strings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
}

func useLock(cli *clientv3.Client) {
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	locker := concurrency.NewLocker(s1, *lockName)

	log.Println("acquiring lock")
	locker.Lock()
	log.Println("acquired lock")

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	locker.Unlock()

	log.Println("released lock")
}

func useMutex(cli *clientv3.Client) {
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	m1 := concurrency.NewMutex(s1, *lockName)

	log.Printf("before acquiring. key: %s", m1.Key())
	log.Println("acquiring lock")
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Printf("acquired lock. key: %s", m1.Key())

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("released lock")

}
