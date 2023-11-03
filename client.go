package godashvector

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client dashvector client
type Client struct {
	APIKey      string
	Endpoint    string
	RestyClient *resty.Client
}

// NewClient creates a new dashvector client
func NewClient(apiKey, endpoint string) *Client {
	rc := resty.New()
	rc.SetHeader("dashvector-auth-token", apiKey)
	rc.SetHeader("Content-Type", "application/json")
	rc.SetBaseURL(fmt.Sprintf("https://dashvector.%s.aliyuncs.com", endpoint))
	return &Client{
		APIKey:      apiKey,
		Endpoint:    endpoint,
		RestyClient: rc,
	}
}
