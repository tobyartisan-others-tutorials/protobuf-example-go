# protobuf-example-go

Generating Go code as part of the Udemy course https://www.udemy.com/course/protocol-buffers/.

## Generating Go code from protobuf

1. Command to generate Go code from `simple.proto`: `protoc -I src/ --go_out=src/ src/simple/simple.proto`
2. Then, `go get .\src\example.simple\`
3. Then, `go build .\src\example.simple\simple.pb.go`
4. Then, `go run src/main.go`
