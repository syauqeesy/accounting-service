# Accounting Service

Install:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Compile protobuf:
```
protoc --proto_path=. --go_out=proto --go-grpc_out=proto proto/*.proto
```