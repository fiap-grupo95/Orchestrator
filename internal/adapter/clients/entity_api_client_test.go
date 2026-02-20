package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
)

func TestEntityAPIClient_ReleasePartsSupply(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v1/parts-supply/release" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
		}))
		defer srv.Close()

		c := NewEntityAPIClient(srv.URL, srv.Client())
		err := c.ReleasePartsSupply(context.Background(), []response.ServiceOrderPartsSupplyResponse{
			{ID: "p1", Quantity: 1},
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer srv.Close()

		c := NewEntityAPIClient(srv.URL, srv.Client())
		if err := c.ReleasePartsSupply(context.Background(), nil); err == nil {
			t.Fatal("expected error")
		}
	})
}
