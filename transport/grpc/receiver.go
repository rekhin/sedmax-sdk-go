package grpc

type Receiver struct {
	socket grpcSocket
}

func NewReceiver(socket *Socket) *Receiver {
	return &Receiver{
		socket: socket.socket,
	}
}

func (s *Receiver) Receive() ([]byte, error) {
	message, err := s.socket.Recv()
	if err != nil {
		return nil, err
	}
	return message.GetData(), nil
}
