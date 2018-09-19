package coinbase

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	// ENDPOINT defaults to https://api.commerce.coinbase.com
	// but can be overridden for test purposes
	ENDPOINT = "https://api.commerce.coinbase.com"
	// API_VERSION since version two you have to
	// specify a API version in your http request headers
	API_VERSION = "2018-03-22"
)

// APIClient is the interface for most of the API calls
// If Endpoint or ApiVersion aren't defined the library
// will use the default https://api.commerce.coinbase.com
type APIClient struct {
	Key        string
	Endpoint   string
	ApiVersion string
	Checkout   *ACheckout
	Charge     *ACharge
	Event      *AEvent
}

func Client(api_key string) (client APIClient) {
	client.Key = api_key
	client.Checkout = new(ACheckout)
	client.Checkout.Api = &client
	client.Charge = new(ACharge)
	client.Charge.Api = &client
	client.Event = new(AEvent)
	client.Event.Api = &client

	return
}

// Fetch works as a wrapper for all kind of http requests. It requires a http method
// and a relative path to the API endpoint. It will try to decode all results into
// a single interface type which you can provide.
func (a *APIClient) Fetch(method, path string, body interface{}, result interface{}) error {
	if a.Endpoint == "" {
		// use default endpoint
		a.Endpoint = ENDPOINT
	}
	if a.ApiVersion == "" {
		// use default api version
		a.ApiVersion = API_VERSION
	}

	client := &http.Client{}
	var bodyBuffered io.Reader
	if body != nil {
		var data []byte
		var err error
		switch body.(type) {
		case string:
			data = []byte(body.(string))
		default:
			data, err = json.Marshal(body)
			if err != nil {
				return err
			}
		}
		bodyBuffered = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(method, a.Endpoint+path, bodyBuffered)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CC-Version", a.ApiVersion)
	req.Header.Set("X-CC-Api-Key", a.Key)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if (resp.StatusCode >= 400){
		return &APIError{Code: resp.StatusCode}
	}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}
