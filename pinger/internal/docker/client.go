package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"pinger/internal/models"
	"time"
)

type DockerClient struct {
	cli *client.Client
}

// NewDockerClient создает новый экземпляр DockerClient
func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error creating docker client: %w", err)
	}
	return &DockerClient{cli: cli}, nil
}

// GetContainers возвращает список контейнеров
func (c *DockerClient) GetContainers() ([]models.Container, error) {
	containers, err := c.cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %w", err)
	}

	var result []models.Container
	for _, cont := range containers {
		inspect, err := c.cli.ContainerInspect(context.Background(), cont.ID)
		if err != nil {
			continue
		}

		var ip string
		if inspect.NetworkSettings != nil {
			for _, network := range inspect.NetworkSettings.Networks {
				ip = network.IPAddress
				break // Берем первый IP
			}
		}

		result = append(result, models.Container{
			ID:           cont.ID[:12],
			Name:         cont.Names[0][1:],
			Status:       cont.State,
			IP:           ip,
			LastPingTime: time.Now(),
			IsActive:     cont.State == "running",
		})
	}

	return result, nil
}
