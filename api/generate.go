//go:generate protoc -I . --go_out=plugins=grpc:. v1/oauth.proto
//go:generate protoc -I . --swagger_out=logtostderr=true:. v1/oauth.proto
package api
