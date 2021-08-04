package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rekhin/sedmax-sdk-go/codec/proto/datapb"
	"github.com/rekhin/sedmax-sdk-go/transport/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	host := flag.String("host", "localhost", "")
	port := flag.Uint("port", 0, "")
	flag.Parse()

	ctx := context.Background()
	dialer := grpc.NewDialer()
	socket, err := dialer.Dial(ctx, fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	defer func() {
		if err := socket.Close(); err != nil {
			log.Fatalf("close failed: %v", err)
		}
	}()
	// log.Println(conn.Target()) // TODO
	sender := grpc.NewSender(socket)
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
		if err := sender.Send(data); err != nil {
			log.Printf("send failed: %s", err)
		}
		time.Sleep(time.Second)
	}
}
