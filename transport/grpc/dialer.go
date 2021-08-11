package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/rekhin/sedmax-sdk-go/transport/grpc/grpcpb"
	"google.golang.org/grpc"
)

type Dialer struct {
	uri string
}

func NewDialer(uri string) *Dialer {
	return &Dialer{
		uri: uri,
	}
}

func (d *Dialer) Dial(ctx context.Context) (*Socket, error) {
	conn, err := grpc.DialContext(ctx, d.uri,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	client := grpcpb.NewGrpcClient(conn)
	socket, err := client.Socket(ctx)
	if err != nil {
		return nil, fmt.Errorf("create socket failed: %s", err)
	}
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Printf("close failed: %s", err) // TODO get rid of the logger
		}
	}()
	return &Socket{
		cancel: cancel,
		socket: socket,
	}, nil
}
