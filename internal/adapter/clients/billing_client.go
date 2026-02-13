package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BillingClient struct {
	base string
	hc   *http.Client
}

func NewBillingClient(base string, hc *http.Client) *BillingClient {
	return &BillingClient{base: strings.TrimRight(base, "/"), hc: hc}
}

func (c *BillingClient) CreateBudget(ctx context.Context, osID string) (string, error) {
	body, _ := json.Marshal(map[string]string{"os_id": osID})

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/budgets", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("billing-service status=%d", resp.StatusCode)
	}

	var out struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if out.ID == "" {
		return "", fmt.Errorf("billing-service returned empty id")
	}
	return out.ID, nil
}

func (c *BillingClient) CancelBudget(ctx context.Context, budgetID string) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/budgets/"+budgetID+"/cancel", nil)
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("billing-service cancel status=%d", resp.StatusCode)
	}
	return nil
}
