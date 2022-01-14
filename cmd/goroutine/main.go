package main

import (
	"context"
	"flag"
	"fmt"
)

func exec(maxcount int) {
	ch := make(chan int)
	cx, cancel := context.WithCancel(context.Background())

	n := 0
	go func(ctx context.Context) {
		// Loop forever
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}(cx)

	for {
		i := <-ch
		if i == maxcount {
			cancel() // send cancel to context
			fmt.Printf("Index: %v Done\n", i)
			return
		} else {
			fmt.Printf("Index: %v\n", i)
		}
	}
}

func main() {
	maxCountPtr := flag.Int("maxcount", 10, "Max count value")
	flag.Parse()
	exec(*maxCountPtr)
}
