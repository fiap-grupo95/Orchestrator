package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
	"net/http"
	"strings"
)

type OSClient struct {
	base  string
	token string
	hc    *http.Client
}

func NewOSClient(base string, token string, hc *http.Client) *OSClient {
	return &OSClient{base: strings.TrimRight(base, "/"), token: token, hc: hc}
}

func (c *OSClient) GetOS(ctx context.Context, id string) (*response.ServiceOrderResponse, error) {
	path := fmt.Sprintf("/v1/service-orders/%s?full_data=true", id)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.base+path, nil)
	req.Header.Set("Content-Type", "application/json")
	c.setAuthHeader(req)

	var osResponse response.ServiceOrderResponse

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("os-service status=%d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&osResponse); err != nil {
		return nil, err
	}

	return &osResponse, nil
}

func (c *OSClient) CreateOS(ctx context.Context, in request.StartInput) (string, error) {
	body, _ := json.Marshal(in)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/v1/service-orders/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.setAuthHeader(req)

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
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/v1/service-orders/"+osID+"/cancel", nil)
	c.setAuthHeader(req)
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

func (c *OSClient) setAuthHeader(req *http.Request) {
	if strings.TrimSpace(c.token) != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
}
