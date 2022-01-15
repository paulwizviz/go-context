package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/paulwizviz/go-context/internal/testclient"
)

func main() {

	cancelInPtr := flag.Int("cancelin", 1, "Cancel context int")
	timeoutPtr := flag.Int("timeout", 1, "Timeout")
	urlPtr := flag.String("url", "http://localhost:8090/hello", "URL")
	flag.Parse()

	timeout := time.Duration(*timeoutPtr) * time.Second
	url := *urlPtr

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Duration(*cancelInPtr)*time.Second, func() {
		cancel()
	})
	resp, err := testclient.Get(ctx, timeout, url)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("First response: ", string(resp))

	resp, err = testclient.Get(ctx, timeout, url)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Second response: ", string(resp))

}
