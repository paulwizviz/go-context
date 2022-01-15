package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/paulwizviz/go-context/internal/testclient"
)

func main() {
	timeoutPtr := flag.Int("timeout", 1, "Timeout")
	urlPtr := flag.String("url", "http://localhost:8080", "URL")
	flag.Parse()

	log.Println("Timeout: ", *timeoutPtr)
	log.Println("URL: ", *urlPtr)

	resp, err := testclient.Get(context.TODO(), time.Duration(*timeoutPtr)*time.Second, *urlPtr)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Response: ", string(resp))

}
