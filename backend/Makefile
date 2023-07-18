generate-proto-user:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/user/user.proto

gqlgen:
	go run github.com/99designs/gqlgen