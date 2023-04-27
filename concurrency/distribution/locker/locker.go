package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr     = flag.String("addr", "http://127.0.0.1:2379", "etcd address")
	lockName = flag.String("name", "my-test-lock", "lock name")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	endpoints := strings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	fmt.Println("cli:", cli)

	fmt.Println("start to use lock")
	useLock(cli)

	fmt.Println("start to use mutex")
	useMutex(cli)
}

func useLock(cli *clientv3.Client) {
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println("err:", err)
	}
	defer s1.Close()
	log.Println("s1:", s1)

	locker := concurrency.NewLocker(s1, *lockName)

	log.Println("acquiring lock")
	locker.Lock()
	log.Println("acquired lock")

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)

	log.Println("releasing lock")
	locker.Unlock()

	log.Println("released lock")
}

func useMutex(cli *clientv3.Client) {
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println("err:", err)
	}
	defer s1.Close()

	log.Println("s1:", s1)

	m1 := concurrency.NewMutex(s1, *lockName)

	log.Printf("before acquiring. key: %s", m1.Key())
	log.Println("acquiring lock")
	if err := m1.Lock(context.TODO()); err != nil {
		log.Println("err:", err)
	}
	log.Printf("acquired lock. key: %s", m1.Key())

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)

	log.Println("releasing lock")
	if err := m1.Unlock(context.TODO()); err != nil {
		fmt.Println(err)
	}
	fmt.Println("released lock")

}
