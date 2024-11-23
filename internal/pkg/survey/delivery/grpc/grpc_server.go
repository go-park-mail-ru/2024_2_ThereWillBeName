package grpc

//go:generate protoc -I . proto/survey.proto --go_out=./gen --go-grpc_out=./gen
