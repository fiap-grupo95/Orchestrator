package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniloAleite/orchestrator/internal/usecase"
)

type OSClient struct {
	base string
	hc   *http.Client
}

func NewOSClient(base string, hc *http.Client) *OSClient {
	return &OSClient{base: strings.TrimRight(base, "/"), hc: hc}
}

func (c *OSClient) CreateOS(ctx context.Context, in usecase.StartInput) (string, error) {
	body, _ := json.Marshal(in)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/service-orders", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("os-service status=%d", resp.StatusCode)
	}

	var out struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if out.ID == "" {
		return "", fmt.Errorf("os-service returned empty id")
	}
	return out.ID, nil
}

func (c *OSClient) CancelOS(ctx context.Context, osID string) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/service-orders/"+osID+"/cancel", nil)
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("os-service cancel status=%d", resp.StatusCode)
	}
	return nil
}
