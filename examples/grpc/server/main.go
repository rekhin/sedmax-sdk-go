package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/rekhin/sedmax-sdk-go/codec/proto/datapb"
	"github.com/rekhin/sedmax-sdk-go/transport/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	host := flag.String("host", "localhost", "")
	port := flag.Uint("port", 0, "")
	flag.Parse()

	ctx := context.Background()

	listener := grpc.NewListener(fmt.Sprintf("%s:%d", *host, *port))
	err := listener.Listen(ctx, func(socket *grpc.Socket) {
		// TODO separate socket to Source and Sink. Now socket argument is Source or Sink
		defer socket.Close()
		receiver := grpc.NewReceiver(socket)
		for {
			data, err := receiver.Receive()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("receive failed: %s", err)
				return
			}
			series := new(datapb.Series)
			err = proto.Unmarshal(data, series)
			if err != nil {
				panic(err)
			}
			log.Println(series)
		}
	})
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}
}
