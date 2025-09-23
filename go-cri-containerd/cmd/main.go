package main

import (
	"fmt"
	"os"
	"github.com/example/go-cri-containerd/internal/cri"
	"github.com/example/go-cri-containerd/internal/containers"
)

func main() {
	socketPath := "/run/containerd/containerd.sock" // Default containerd socket path
	if len(os.Args) > 1 {
		socketPath = os.Args[1]
	}
	client, err := cri.NewCRIClient(socketPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to CRI: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	containersList, err := containers.ListActiveContainers(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list containers: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Active containers:")
	for _, c := range containersList {
		fmt.Printf("ID: %s, Name: %s, SandboxID: %s\n", c.Id, c.Metadata.Name, c.PodSandboxId)
	}
}
