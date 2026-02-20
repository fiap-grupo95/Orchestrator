package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExecutionClient_StartExecution(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/executions/start" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
		}))
		defer srv.Close()

		c := NewExecutionClient(srv.URL, srv.Client())
		if err := c.StartExecution(context.Background(), "os-1"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
		}))
		defer srv.Close()

		c := NewExecutionClient(srv.URL, srv.Client())
		if err := c.StartExecution(context.Background(), "os-1"); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestExecutionClient_CancelExecution(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v1/executions/cancel/os-1" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
		}))
		defer srv.Close()

		c := NewExecutionClient(srv.URL, srv.Client())
		if err := c.CancelExecution(context.Background(), "os-1"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer srv.Close()

		c := NewExecutionClient(srv.URL, srv.Client())
		if err := c.CancelExecution(context.Background(), "os-1"); err == nil {
			t.Fatal("expected error")
		}
	})
}
