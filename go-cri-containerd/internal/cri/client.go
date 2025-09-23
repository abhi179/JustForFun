package cri

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	runtimeapi "github.com/containerd/cri/pkg/apis/runtime/v1"
)

type CRIClient struct {
	RuntimeService runtimeapi.RuntimeServiceClient
	conn           *grpc.ClientConn
}

func NewCRIClient(socketPath string) (*CRIClient, error) {
	conn, err := grpc.Dial(socketPath, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to CRI socket: %w", err)
	}
	runtimeClient := runtimeapi.NewRuntimeServiceClient(conn)
	return &CRIClient{
		RuntimeService: runtimeClient,
		conn:           conn,
	}, nil
}

func (c *CRIClient) Close() error {
	return c.conn.Close()
}
