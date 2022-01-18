package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Key string
type Value string

const (
	ContextKey Key = Key("Hello")
)

func hdlr(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Printf("Conext value from client: %v", ctx.Value(ContextKey))
	ctx = context.WithValue(ctx, ContextKey, Value("World"))
	fmt.Fprintf(w, "Hello %v", ctx.Value(ContextKey))
}

func server() {
	http.HandleFunc("/", hdlr)
	http.ListenAndServe(":9090", nil)
}

func get() ([]byte, error) {
	ctx := context.WithValue(context.Background(), ContextKey, Value("What's up!"))
	client := &http.Client{}
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:9090/", nil)
	if err != nil {
		return nil, err
	}
	log.Printf("Request: %v", request.Context().Value(ContextKey))
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func main() {

	if len(os.Args) == 1 {
		log.Fatalf("Insufficient arguments")
	}
	arg := os.Args[1]
	switch arg {
	case "server":
		log.Println("starting server")
		server()
	case "client":
		body, err := get()
		if err != nil {
			log.Printf("Error: %v", err)
		}
		log.Printf("%v", string(body))
	}
}
