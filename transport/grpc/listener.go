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

func (l *Listener) Listen(ctx context.Context, f func(*Socket)) error {
	listener, err := net.Listen("tcp", l.uri)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcpb.RegisterSinkServer(grpcServer, NewSinkServer(f))
	grpcServer.Serve(listener)
	return nil
}

type SinkServer struct {
	grpcpb.UnimplementedSinkServer

	f func(*Socket)
}

func NewSinkServer(f func(*Socket)) *SinkServer {
	return &SinkServer{
		f: f,
	}
}

func (s *SinkServer) Pipe(pipe grpcpb.Sink_PipeServer) error {
	// s.f()

	// for {
	// 	message, err := pipe.Recv()
	// 	if err == io.EOF {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		return fmt.Errorf("receive failed: %s", err)
	// 	}
	// 	series := new(datapb.Series)
	// 	err = proto.Unmarshal(message.GetData(), series)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(series)
	// }

	return nil
}
