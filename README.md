# protobuf-example-go

Generating Go code as part of the Udemy course https://www.udemy.com/course/protocol-buffers/.

## Generating Go code from protobuf

1. Generate Go code: `protoc -I src/ --go_out=src/ src/simple/simple.proto`
2. Build: `go build main.go`
3. Add more packages using `go get ...` as directed.
4. Start program: `go run src/main.go`
