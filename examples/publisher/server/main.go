package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/rekhin/sedmax-sdk-go/codec/proto/datapb"
	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	host := flag.String("host", "localhost", "")
	port := flag.Uint("port", 0, "")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println(lis.Addr().String()) // TODO
	var opts []grpc.ServerOption
	// ...
	grpcServer := grpc.NewServer(opts...)
	grpcpb.RegisterSinkServer(grpcServer, NewSinkServer())
	grpcServer.Serve(lis)
}

type SinkServer struct {
	grpcpb.UnimplementedSinkServer
}

func NewSinkServer() *SinkServer {
	return &SinkServer{}
}

func (SinkServer) Pipe(pipe grpcpb.Sink_PipeServer) error {
	grpc.NewReceiver(pipe)
	for {
		message, err := pipe.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("receive failed: %s", err)
		}
		series := new(datapb.Series)
		err = proto.Unmarshal(message.GetData(), series)
		if err != nil {
			panic(err)
		}
		log.Println(series)
	}
}
