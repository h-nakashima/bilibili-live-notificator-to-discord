package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/xerrors"
)

type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
	UserAgent   string
}

func NewClient(endpointURL string, httpClient *http.Client, userAgent string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse url: %s: %w", endpointURL, err)
	}

	client := &Client{
		EndpointURL: parsedURL,
		HTTPClient:  httpClient,
		UserAgent:   userAgent,
	}
	return client, nil
}

func (client *Client) NewRequest(ctx context.Context, method string, subPath string, query string, body io.Reader) (*http.Request, error) {
	// TODO: endpoint"URL"ではないので名前を変更
	endpointURL := *client.EndpointURL
	endpointURL.Path = path.Join(client.EndpointURL.Path, subPath)
	endpointURL.RawQuery = query

	req, err := http.NewRequest(method, endpointURL.String(), body)
	if err != nil {
		return nil, xerrors.Errorf("failed to create new request: %s: %w", endpointURL.String(), err)
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", client.UserAgent)

	return req, nil
}

func (client *Client) DecodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
