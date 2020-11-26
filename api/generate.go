//go:generate protoc -I . --go_out=plugins=grpc:. v1/oauth.proto v1/account.proto
//go:generate protoc -I . --swagger_out=logtostderr=true:. v1/oauth.proto v1/account.proto
package api
