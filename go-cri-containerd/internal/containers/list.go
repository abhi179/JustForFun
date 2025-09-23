package containers

import (
	"context"
	"fmt"
	"time"
	"github.com/example/go-cri-containerd/internal/cri"
	runtimeapi "github.com/containerd/cri/pkg/apis/runtime/v1"
)

func ListActiveContainers(client *cri.CRIClient) ([]*runtimeapi.Container, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &runtimeapi.ListContainersRequest{
		Filter: &runtimeapi.ContainerFilter{
			State: &runtimeapi.ContainerStateValue{
				State: runtimeapi.ContainerState_CONTAINER_RUNNING,
			},
		},
	}
	resp, err := client.RuntimeService.ListContainers(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}
	return resp.Containers, nil
}
