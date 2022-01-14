package helloserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHdlr(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println("server: handler started")
	defer log.Println("server: handler ended")

	select {
	case <-time.After(5 * time.Second):
		// Wait until five seconds and then send response
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// If client end is cancelled before 5 seconds this will
		// be called
		err := ctx.Err()
		log.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func Exec() {
	http.HandleFunc("/hello", helloHdlr)
	http.ListenAndServe(":8090", nil)
}
