package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pinger/internal/docker"
	"pinger/internal/models"
	"time"
)

const (
	backendURL = "http://backend:8082/api/pinger"
	interval   = 10 * time.Second
)

// sendContainerUpdate отправляет данные о контейнере в бэкенд
func sendContainerUpdate(container models.Container) error {
	jsonData, err := json.Marshal(container)
	if err != nil {
		return fmt.Errorf("error marshaling container data: %w", err)
	}

	resp, err := http.Post(backendURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending data to backend: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Updated container: %s (ID: %s, Status: %s)",
		container.Name, container.ID, container.Status)
	return nil
}

// updateContainers обновляет данные о контейнерах и отправляет их в бэкенд
func updateContainers(client *docker.DockerClient) error {
	containers, err := client.GetContainers()
	if err != nil {
		return fmt.Errorf("error getting containers: %w", err)
	}

	for _, container := range containers {
		if err := sendContainerUpdate(container); err != nil {
			log.Printf("Error updating container %s: %v", container.ID, err)
		}
	}

	log.Printf("Successfully updated %d containers", len(containers))
	return nil
}

func main() {
	client, err := docker.NewDockerClient()
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	log.Printf("Starting container monitoring (interval: %s)", interval)

	// Первичное обновление
	if err := updateContainers(client); err != nil {
		log.Printf("Initial container update failed: %v", err)
	}

	// Периодические обновления
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if err := updateContainers(client); err != nil {
			log.Printf("Container update failed: %v", err)
		}
	}
}
