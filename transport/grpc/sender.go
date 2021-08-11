package grpc

import (
	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
)

type Sender struct {
	socket grpcSocket
}

func NewSender(socket *Socket) *Sender {
	return &Sender{
		socket: socket.socket,
	}
}

func (s *Sender) Send(data []byte) error {
	message := &grpcpb.SocketMessage{
		SequenceNumber:       0,
		AcknowledgmentNumber: 0,
		Data:                 data,
	}
	if err := s.socket.Send(message); err != nil {
		return err
	}
	return nil
}
