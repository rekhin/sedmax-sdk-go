package grpc

type Receiver struct {
	pipe Pipe
}

func NewReceiver(socket *Socket) *Receiver {
	return &Receiver{
		pipe: socket.pipe,
	}
}

func (s *Receiver) Receive() ([]byte, error) {
	message, err := s.pipe.Recv()
	if err != nil {
		return nil, err
	}
	return message.GetData(), nil
}
