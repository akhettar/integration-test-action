package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// DefaultTimeout the default timeout
const (
	DefaultTimeout              = 40 * time.Second
	DefaultRunDockerComposeFlag = false
	InputTimeout                = "INPUT_TIMEOUT"
	InputReadinessEndpiont      = "INPUT_READINESS_ENDPOINT"
)

// ReadinessCheck type
type ReadinessCheck struct {
	timeout  time.Duration
	endpoint string
	client   *http.Client
}

// New create an instance of the HealthCheck
func New() *ReadinessCheck {
	timeout := DefaultTimeout
	if t := os.Getenv(InputTimeout); t != "" {
		if n, err := strconv.Atoi(t); err == nil {
			timeout = time.Duration(n) * time.Second
		}
	}

	url, ok := os.LookupEnv(InputReadinessEndpiont)
	if !ok {
		log.Fatal("the readiness endpoint must be provided")
	}

	c := &http.Client{Timeout: 1 * time.Second}
	return &ReadinessCheck{client: c, endpoint: url, timeout: timeout}
}

func (h *ReadinessCheck) check() error {
	start := time.Now()
	for time.Since(start) < h.timeout {
		if res, err := h.client.Get(h.endpoint); err == nil && res.StatusCode == http.StatusOK {
			return nil
		}
		log.Println("proving the health endpoint")
	}
	return fmt.Errorf("failed to check the readiness of the given endpoint on time: %v", h.timeout)
}

func main() {
	if err := New().check(); err != nil {
		log.Fatalf("failed to prove the readiness endpoint with error: %v", err)
	}
	log.Println("The service is up and running")
}
