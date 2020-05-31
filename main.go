package main

import "log"

func main() {
	if err := NewReadinessCheck().check(); err != nil {
		log.Fatalf("failed to prove the readiness endpoint with error: %v", err)
	}
	log.Println("The service is up and running")
}
