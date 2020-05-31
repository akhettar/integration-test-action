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

// NewReadinessCheck create an instance of the HealthCheck
func NewReadinessCheck() *ReadinessCheck {
	timeout := DefaultTimeout
	if t := os.Getenv(InputTimeout); t != "" {
		if n, err := strconv.Atoi(t); err == nil {
			timeout = time.Duration(n) * time.Second
		}
	}

	url, ok := os.LookupEnv(InputReadinessEndpiont)
	//url, ok := os.LookupEnv("GITHUB_SERVER_URL")
	if !ok {
		log.Fatal("the readiness endpoint must be provided")
	}

	log.Printf("*****GITHUB_SERVER_URL: %s\n", url)
	log.Printf("*****ACTIONS_RUNTIME_URL: %s\n", os.Getenv("ACTIONS_RUNTIME_URL"))
	c := &http.Client{Timeout: 1 * time.Second}
	// return &ReadinessCheck{client: c, endpoint: "http://" + url + ":8080/v1/health", timeout: timeout}
	return &ReadinessCheck{client: c, endpoint: url, timeout: timeout}
}

func (h *ReadinessCheck) check() error {
	start := time.Now()
	for time.Since(start) < h.timeout {
		if res, err := h.client.Get(h.endpoint); err == nil && res.StatusCode == http.StatusOK {
			return nil
		}
		log.Printf("checking the health endpoint: %s\n", h.endpoint)
	}
	return fmt.Errorf("failed to check the readiness of the given endpoint on time: %v", h.timeout)
}
