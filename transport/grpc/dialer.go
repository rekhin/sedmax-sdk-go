package grpc

import (
	"context"
	"fmt"

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
	sinkClient := grpcpb.NewSinkClient(conn)
	pipe, err := sinkClient.Pipe(ctx)
	if err != nil {
		return nil, fmt.Errorf("create pipe failed: %s", err)
	}
	return &Socket{
		pipe: pipe,
	}, nil
}
