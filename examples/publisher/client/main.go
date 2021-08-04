package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rekhin/sedmax-sdk-go/codec/proto/datapb"
	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	host := flag.String("host", "localhost", "")
	port := flag.Uint("port", 0, "")
	flag.Parse()
	var opts []grpc.DialOption = []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *host, *port), opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("failed to close: %v", err)
		}
	}()
	log.Println(conn.Target()) // TODO
	sinkClient := grpcpb.NewSinkClient(conn)
	pipe, err := sinkClient.Pipe(context.Background())
	if err != nil {
		log.Fatalf("failed to publish: %v", err)
	}
	for {
		series := &datapb.Series{
			Sets: []*datapb.Set{
				{
					Tag: "temperature",
					Points: []*datapb.Point{
						{
							Timestamp: time.Now().UnixNano(),
							Status:    0x00,
							Value:     &datapb.Value{Variant: &datapb.Value_Float64{Float64: 123.456}},
						},
						{
							Timestamp: time.Now().UnixNano(),
							Status:    0x00,
							Value:     &datapb.Value{Variant: &datapb.Value_Float64{Float64: 654.321}},
						},
						{
							Timestamp: time.Now().UnixNano(),
							Status:    0x00,
							Value: &datapb.Value{Variant: &datapb.Value_Array{Array: &datapb.Array{StringArray: []string{
								"foo", "bar", "baz",
							}}}},
						},
					},
				},
			},
		}
		log.Println(series)
		data, err := proto.Marshal(series)
		if err != nil {
			panic(err)
		}
		message := &grpcpb.PipeMessage{
			SequenceNumber:       0,
			AcknowledgmentNumber: 0,
			Data:                 data,
		}
		if err := pipe.Send(message); err != nil {
			log.Fatalf("failed to send %s", err)
		}
		time.Sleep(time.Second)
	}
}
