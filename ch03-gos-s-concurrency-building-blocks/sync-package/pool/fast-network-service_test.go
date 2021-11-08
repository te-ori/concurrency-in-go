package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

var address string = "localhost:8010"

func fconnectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: fconnectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func fstartNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("cannot accept conneciton: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	daemonStarted := fstartNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequestF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			b.Fatalf("can not dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
