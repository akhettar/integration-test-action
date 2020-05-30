package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReadinessCheck_check(t *testing.T) {
	type fields struct {
		timeout  string
		endpoint string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Service down",
			fields:  fields{timeout: "2", endpoint: "http://localhost:8080/health"},
			wantErr: true,
		},

		{
			name:    "Service up and running",
			fields:  fields{timeout: "2", endpoint: startTestServer()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEndpiontEnvVariable(tt.fields.endpoint, tt.fields.timeout)
			h := New()
			if err := h.check(); (err != nil) != tt.wantErr {
				t.Errorf("ReadinessCheck.check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func setEndpiontEnvVariable(url string, timeout string) {
	os.Setenv(InputReadinessEndpiont, url)
	os.Setenv(InputTimeout, timeout)
}

func startTestServer() (url string) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "service up and running")
	}))
	return svr.URL
}
