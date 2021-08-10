package grpc

import (
	"context"

	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
)

type Socket struct {
	cancel     context.CancelFunc
	sourcePipe Pipe
	// sinkPipe   Pipe
}

type Pipe interface {
	Send(*grpcpb.PipeMessage) error
	Recv() (*grpcpb.PipeMessage, error)
}

func (s *Socket) Close() {
	s.cancel()
}
