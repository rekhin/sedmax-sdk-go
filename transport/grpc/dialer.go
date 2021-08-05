package grpc

import (
	"context"

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
	conn, err := grpc.DialContext(ctx, d.uri, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return (*Socket)(conn), nil
}
