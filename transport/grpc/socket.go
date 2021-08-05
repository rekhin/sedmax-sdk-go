package grpc

import (
	"context"
)

// type Socket grpc.ClientConn

// func (s *Socket) Close() error {
// 	return (*grpc.ClientConn)(s).Close()
// }

type Socket struct {
	cancel context.CancelFunc
	pipe   Pipe
}

func (s *Socket) Close() {
	s.cancel()
}
