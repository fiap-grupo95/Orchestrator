package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
	"net/http"
	"strings"
)

type EntityAPIClient struct {
	base string
	hc   *http.Client
}

func NewEntityAPIClient(base string, hc *http.Client) *EntityAPIClient {
	return &EntityAPIClient{base: strings.TrimRight(base, "/"), hc: hc}
}

func (c *EntityAPIClient) ReleasePartsSupply(ctx context.Context, partsSupply []response.ServiceOrderPartsSupplyResponse) error {
	body, _ := json.Marshal(partsSupply)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/v1/parts-supply/release", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("entity-api status=%d", resp.StatusCode)
	}
	return nil
}
