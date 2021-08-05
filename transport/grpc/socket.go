package grpc

import "google.golang.org/grpc"

type Socket grpc.ClientConn

func (s *Socket) Close() error {
	return (*grpc.ClientConn)(s).Close()
}

// type Socket gstruct {

// }

// func (s *Socket) Close() error {
// 	return (*grpc.ClientConn)(s).Close()
// }
