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
	grpcpb.RegisterSinkServer(grpcServer, NewSinkServer(ctx, f))
	grpcServer.Serve(listener)
	return nil
}

type SinkServer struct {
	grpcpb.UnimplementedSinkServer
	ctx context.Context
	f   func(*Socket)
}

func NewSinkServer(ctx context.Context, f func(*Socket)) *SinkServer {
	return &SinkServer{
		ctx: ctx,
		f:   f,
	}
}

func (s *SinkServer) Pipe(pipe grpcpb.Sink_PipeServer) error {
	ctx, cancel := context.WithCancel(s.ctx)
	s.f(&Socket{
		cancel: cancel,
		pipe:   pipe,
	})
	<-ctx.Done()
	return ctx.Err()
}
