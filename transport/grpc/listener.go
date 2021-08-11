package grpc

import (
	"context"
	"log"
	"net"

	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
	"google.golang.org/grpc"
)

type Listener struct {
	uri string
}

func NewListener(uri string) *Listener {
	return &Listener{
		uri: uri,
	}
}

func (l *Listener) Listen(ctx context.Context, f func(socket *Socket)) error {
	listener, err := net.Listen("tcp", l.uri)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcpb.RegisterGrpcServer(grpcServer, NewGrpcServer(ctx, f))
	grpcServer.Serve(listener)
	return nil
}

type GrpcServer struct {
	grpcpb.UnimplementedGrpcServer
	ctx context.Context
	f   func(*Socket)
}

func NewGrpcServer(ctx context.Context, f func(*Socket)) *GrpcServer {
	return &GrpcServer{
		ctx: ctx,
		f:   f,
	}
}

func (s *GrpcServer) Socket(socket grpcpb.Grpc_SocketServer) error {
	ctx, cancel := context.WithCancel(s.ctx)
	s.f(&Socket{
		cancel: cancel,
		socket: socket,
	})
	<-ctx.Done()
	return ctx.Err()
}
