package main

import (
	"context"
	"flag"
	"fmt"
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
	cancelin := time.Duration(*cancelInPtr)
	url := *urlPtr

	ctx, cancel := context.WithTimeout(context.Background(), cancelin*time.Second)
	defer cancel()

	resp, err := testclient.Get(ctx, timeout, url)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(resp))

}
