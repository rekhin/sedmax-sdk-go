package grpc

import (
	"context"
	"fmt"

	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
	"google.golang.org/grpc"
)

type Sender struct {
	sinkClient grpcpb.SinkClient
	pipe       Pipe
}

func NewSender(socket *Socket) *Sender {
	sinkClient := grpcpb.NewSinkClient((*grpc.ClientConn)(socket))
	return &Sender{
		sinkClient: sinkClient,
	}
}

func (s *Sender) Send(data []byte) error {
	if s.pipe == nil {
		pipe, err := s.sinkClient.Pipe(context.Background())
		if err != nil {
			return fmt.Errorf("create pipe failed: %s", err)
		}
		s.pipe = pipe
	}
	message := &grpcpb.PipeMessage{
		SequenceNumber:       0,
		AcknowledgmentNumber: 0,
		Data:                 data,
	}
	if err := s.pipe.Send(message); err != nil {
		s.pipe = nil // ?
		return err
	}
	return nil
}
