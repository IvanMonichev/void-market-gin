package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Clients struct {
	User  *resty.Client
	Order *resty.Client
}

// SetClient создает клиента с базовым URL и логами
func SetClient(baseURL string) *resty.Client {
	return resty.New().
		SetBaseURL(baseURL).
		SetHeader("Content-Type", "application/json").
		OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
			fmt.Printf("[OUT] %s %s\n", req.Method, req.URL)
			return nil
		}).
		OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
			fmt.Printf("[IN ] %d %s\n", resp.StatusCode(), resp.Request.URL)
			return nil
		})
}

func NewClients(userURL, orderURL string) *Clients {
	return &Clients{
		User:  SetClient(userURL),
		Order: SetClient(orderURL),
	}
}
