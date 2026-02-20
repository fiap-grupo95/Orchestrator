package clients

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
)

func TestOSClient_GetOS(t *testing.T) {
	t.Run("success with auth header", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v1/service-orders/123" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			if r.URL.RawQuery != "full_data=true" {
				t.Fatalf("unexpected query: %s", r.URL.RawQuery)
			}
			if got := r.Header.Get("Authorization"); got != "Bearer token-1" {
				t.Fatalf("unexpected auth header: %s", got)
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id":"123","status":"PENDING"}`))
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "token-1", srv.Client())
		out, err := c.GetOS(context.Background(), "123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out == nil || out.ID != "123" {
			t.Fatalf("unexpected output: %+v", out)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		if _, err := c.GetOS(context.Background(), "x"); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{"))
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		if _, err := c.GetOS(context.Background(), "x"); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestOSClient_CreateOS(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Fatalf("unexpected method: %s", r.Method)
			}
			if r.URL.Path != "/v1/service-orders/create" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			raw, _ := io.ReadAll(r.Body)
			body := string(raw)
			if !strings.Contains(body, `"customer_id":"c1"`) {
				t.Fatalf("unexpected body: %s", body)
			}
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"id":"os-1"}`))
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		id, err := c.CreateOS(context.Background(), request.StartInput{CustomerID: "c1", VehicleID: "v1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != "os-1" {
			t.Fatalf("expected os-1, got %s", id)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		if _, err := c.CreateOS(context.Background(), request.StartInput{}); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("empty id", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"id":""}`))
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		if _, err := c.CreateOS(context.Background(), request.StartInput{}); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestOSClient_CancelOS(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v1/service-orders/os-2/cancel" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		if err := c.CancelOS(context.Background(), "os-2"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer srv.Close()

		c := NewOSClient(srv.URL, "", srv.Client())
		if err := c.CancelOS(context.Background(), "x"); err == nil {
			t.Fatal("expected error")
		}
	})
}
