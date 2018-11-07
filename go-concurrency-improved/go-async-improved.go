package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

//Go async average time for  1000  HEAD requests
//638.133765ms


func main() {
	start := time.Now()
	concurrency := flag.Int("concurrency", Concurrency, "Concurrent connection to the server")
	req := flag.String("req", Url, "Request url")
	keepAlive := flag.Bool("keepAlive", KeepAlive, "Whether to keep connection alive, if enabled keep alive for 5 min")
	iterCnt := flag.Int("iterCnt", IterationCount, "how many iterations to run")
	reqCnt := flag.Int("reqCnt", RequestCount, "how many request in an iteration")
	flag.Parse()

	bucket := make(chan struct{}, *iterCnt * *reqCnt)
	fin := make(chan struct{})
	var ops uint32 = 0

	go func() {
		for i := 0; i < *iterCnt * *reqCnt; i++{
			bucket <- struct{}{}
		}
	}()

	for i :=  0; i < *concurrency; i++{
		go func(){
			tr := &http.Transport{
				Dial: makeDialer(*keepAlive),
				TLSHandshakeTimeout: time.Second * 10,
				DisableKeepAlives: !(*keepAlive),
				MaxIdleConnsPerHost: *concurrency,
			}

			client := &http.Client{Transport: tr}

			for {
				<- bucket
				_, err := client.Head(*req)
				if err != nil{
					println(err)
				}
				atomic.AddUint32(&ops, 1)
			}
		}()
	}

	go func(){
		i := 1
		for{
			if ops >= uint32(*reqCnt * i){
				fmt.Println(time.Since(start))
				start = time.Now()
				i++
			}
			if ops == uint32(*iterCnt * *reqCnt){
				fin <- struct{}{}
			}
			time.Sleep(time.Millisecond)
		}
	}()

	<- fin

}

type DialerFunc func(network, addr string) (net.Conn, error)

func makeDialer(keepAlive bool) DialerFunc{
	return func(network, addr string) (net.Conn, error){
		conn, err := (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 1000 * time.Second,
		}).Dial(network, addr)

		if err != nil {
			return conn, err
		}
		if !keepAlive {
			conn.(*net.TCPConn).SetLinger(0)
		}
		return conn, err
	}
}
