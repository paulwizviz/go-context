# Overview

A series of examples to demonstrate Go context use cases.

## Support tools

Most of the examples here use a combination of client and server to demonstrate the use of Go context.

The server, known as `helloserver`, takes a GET request which is deloyed for 5 seconds before a response. To run the server, use this command `go run cmd/helloserver/main.go`.

## Context between http client and server

Start a http server.

Use the client `go run cmd/httpclient/main.go -timeout=<second> -url="<url of server>"`

## Context to stop a goroutine

In this example, you have a goroutine that loop forever and another main routine loop until a given value is met. When the value is met, the context is cancelled.

Run the app `go run cmd/goroutine/main.go -maxcount=<int value to make rountine stop>`

## Context to cancel a process with to http client calls

In this example, the main process two clients calling the hello server, consecutively. Each call to the server will timeout or cancalled depending on arguments provided.

* `timeout` is the duration the client will stay alive.
* `cancelin` is the duration from the start of the main process before the request context is cancelled.

Steps:

1. Start a http server.
1. Run the app this way `go run cmd/clientctx/main.go -cancelin=<time in seconds from start of app to cancel main process> -timeout=<time in second for client to time out> -url=<url to hello server>`

## Conext with timeout

This example uses `context.WithTimeout(context.Background(), cancelin*time.Second)` to set timeout.

To see the example in action, run this command `go run cmd/clientctx/main.go -cancelin=<time in seconds from start of app to cancel main process> -timeout=<time in second for client to time out> -url=<url to hello server>`

Scenarios:

Scenario 1. Context cancel in 1 secs

```
go run cmd/ctxtimeout/main.go -timeout=100 -cancelin=1 url="http://localhost:8090/hello"
2022/01/17 19:01:15 Get "http://localhost:8090/hello": context deadline exceeded
exit status 1
```

Scenario 2. Context cancel in 10 secs

```
go run cmd/ctxtimeout/main.go -timeout=100 -cancelin=10 url="http://localhost:8090/hello"
hello
```

Scenario 3. Client timeout in 1 sec. Context timeout in 10 secs.

```
go run cmd/ctxtimeout/main.go -timeout=1 -cancelin=10 url="http://localhost:8090/hello"
2022/01/17 19:06:59 Get "http://localhost:8090/hello": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
exit status 1
```

## Client server context with value

This example demonstrate the fact that you can't send a context value from a client to server using the standard package request.

On the client side:
```
go run cmd/ctxvalue/main.go client
2022/01/18 22:14:36 Request: What's up!
2022/01/18 22:14:36 Hello World
```

On the servier side:
```
2022/01/18 22:14:36 Conext value from client: <nil>
```