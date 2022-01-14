# Overview

A series of examples to demonstrate Go context use cases.

## Context between http client and server

Start a http server. `go run cmd/helloserver/main.go`

Use the client `go run cmd/httpclient/main.go -timeout=<second> -url="<url of server>"`