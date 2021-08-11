package sedmax

//go:generate protoc --go_out=plugins=grpc:. proto/measurement.proto

//go:generate protoc --go_out=. --go-grpc_out=. proto/grpc.proto
//go:generate protoc --go_out=. --go-grpc_out=. proto/data.proto
