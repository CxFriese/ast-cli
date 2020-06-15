package wrappers

import "fmt"

type HealthStatus struct {
	Success bool
	Message string
}

func (h *HealthStatus) String() string {
	if h.Success {
		return "Success"
	}

	return fmt.Sprintf("Failure, due to %v", h.Message)
}

type HealthCheckWrapper interface {
	RunWebAppCheck() (*HealthStatus, error)
}