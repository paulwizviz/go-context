# Overview

A series of examples to demonstrate Go context use cases.

## Context between http client and server

Start a http server. `go run cmd/helloserver/main.go`

Use the client `go run cmd/httpclient/main.go -timeout=<second> -url="<url of server>"`

## Context to stop a goroutine

In this example, you have a goroutine that loop forever and another main routine loop until a given value is met. When the value is met, the context is cancelled.

Run the app `go run cmd/goroutine/main.go -maxcount=<int value to make rountine stop>`

## Context to cancel a process with to http client calls

In this example, the app has two clients calling the hello server. The hello server takes 5 seconds to response. The process will invoke a context cancel at a specified duration from start time.

Start a http server. `go run cmd/helloserver/main.go`

Run the app this way `go run cmd/clientctx/main.go -cancelin=<time in seconds from start of app to cancel main process> -timeout=<time in second for client to time out> -url=<url to hello server>`

If context cancel is called before response is done, the main routine will cancel.