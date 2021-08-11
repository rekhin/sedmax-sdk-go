package grpc

import (
	"context"
	"log"

	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
)

type Sender struct {
	sequenceNumber       uint32
	acknowledgmentNumber uint32
	socket               grpcSocket
	buffer               Buffer
}

type Buffer interface {
	GetNumber() uint32
	Push(data []byte) uint32
	Pop() []byte
	Acknowledge(number uint32)
}

func NewSender(socket *Socket) *Sender {
	return &Sender{
		socket: socket.socket,
	}
}

func (s *Sender) Start(ctx context.Context) {
	go func() {
		for {
			message, err := s.socket.Recv()
			if err != nil {
				log.Printf("receive failed: %s", err) // TODO log -> GroupErr
			}
			if message.AcknowledgmentNumber < s.buffer.GetNumber() { //atomic.LoadUint32(&s.sequenceNumber) {
				s.buffer.Pop()
			}
			s.buffer.Acknowledge(message.AcknowledgmentNumber)
			// s.acknowledgmentNumber = message.AcknowledgmentNumber
		}
	}()
}

// TODO везде в Send и Receive добавить ctx
func (s *Sender) Send(data []byte) error {
	sequenceNumber := s.buffer.Push(data)
	message := &grpcpb.SocketMessage{
		SequenceNumber:       sequenceNumber, //atomic.AddUint32(&s.sequenceNumber, 1),
		AcknowledgmentNumber: 0,
		Data:                 data,
	}
	if err := s.socket.Send(message); err != nil { // TODO отправку переместить в общий цикл, причем данные для отправки брать из буфера
		return err
	}
	return nil
}
