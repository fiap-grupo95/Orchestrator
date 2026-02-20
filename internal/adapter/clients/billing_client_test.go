package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBillingClient_CreateBudget(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/budgets" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"id":"b-1"}`))
		}))
		defer srv.Close()

		c := NewBillingClient(srv.URL, srv.Client())
		id, err := c.CreateBudget(context.Background(), "os-1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != "b-1" {
			t.Fatalf("expected b-1, got %s", id)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer srv.Close()

		c := NewBillingClient(srv.URL, srv.Client())
		if _, err := c.CreateBudget(context.Background(), "os-1"); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("empty id", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id":""}`))
		}))
		defer srv.Close()

		c := NewBillingClient(srv.URL, srv.Client())
		if _, err := c.CreateBudget(context.Background(), "os-1"); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestBillingClient_CancelBudget(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/budgets/b-1/cancel" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
		}))
		defer srv.Close()

		c := NewBillingClient(srv.URL, srv.Client())
		if err := c.CancelBudget(context.Background(), "b-1"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer srv.Close()

		c := NewBillingClient(srv.URL, srv.Client())
		if err := c.CancelBudget(context.Background(), "b-1"); err == nil {
			t.Fatal("expected error")
		}
	})
}
