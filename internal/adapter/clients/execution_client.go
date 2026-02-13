package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ExecutionClient struct {
	base string
	hc   *http.Client
}

func NewExecutionClient(base string, hc *http.Client) *ExecutionClient {
	return &ExecutionClient{base: strings.TrimRight(base, "/"), hc: hc}
}

func (c *ExecutionClient) StartExecution(ctx context.Context, osID string) error {
	body, _ := json.Marshal(map[string]string{"os_id": osID})

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/executions/start", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("execution-service status=%d", resp.StatusCode)
	}
	return nil
}
