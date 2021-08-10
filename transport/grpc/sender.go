package grpc

import (
	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
)

type Sender struct {
	pipe Pipe
}

func NewSender(socket *Socket) *Sender {
	return &Sender{
		pipe: socket.sourcePipe,
	}
}

func (s *Sender) Send(data []byte) error {
	message := &grpcpb.PipeMessage{
		SequenceNumber:       0,
		AcknowledgmentNumber: 0,
		Data:                 data,
	}
	if err := s.pipe.Send(message); err != nil {
		return err
	}
	return nil
}
