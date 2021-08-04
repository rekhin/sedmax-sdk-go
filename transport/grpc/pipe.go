package grpc

import "github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"

type Pipe interface {
	Send(*grpcpb.PipeMessage) error
	Recv() (*grpcpb.PipeMessage, error)
}
