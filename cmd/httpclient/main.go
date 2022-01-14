package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	timeoutPtr := flag.Int("timeout", 1, "Timeout")
	urlPtr := flag.String("url", "http://localhost:8080", "URL")
	flag.Parse()

	log.Println("Timeout: ", *timeoutPtr)
	log.Println("URL: ", *urlPtr)

	client := &http.Client{
		Timeout: time.Duration(*timeoutPtr) * time.Second,
	}

	request, err := http.NewRequest("GET", *urlPtr, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(body))

}
