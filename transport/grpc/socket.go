package grpc

import (
	"context"

	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
)

type Socket struct {
	cancel context.CancelFunc
	socket grpcSocket
}

type grpcSocket interface {
	Send(*grpcpb.SocketMessage) error
	Recv() (*grpcpb.SocketMessage, error)
}

func (s *Socket) Close() {
	s.cancel()
}
