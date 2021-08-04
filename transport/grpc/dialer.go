package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type Dialer struct{}

func NewDialer() *Dialer {
	return &Dialer{}
}

func (Dialer) Dial(ctx context.Context, uri string) (*Socket, error) {
	conn, err := grpc.DialContext(ctx, uri, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return (*Socket)(conn), nil
}
