package coinbase

import (
	"io"
	"bytes"
  "encoding/json"
  "net/http"
	"fmt"
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
  Key string
  Endpoint string
  ApiVersion string
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
    data, err := json.Marshal(body)
    if err != nil {
      return err
    }
		fmt.Println("ici print", string(data))
    bodyBuffered = bytes.NewBuffer([]byte(string(data)))
  }
  req, err := http.NewRequest(method, a.Endpoint + path, bodyBuffered)
  if err != nil {
    return err
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-CC-Version", a.ApiVersion)
  // do not authenticate on public time api call
  if path[len(path)-4:] != "time" {
    err = a.Authenticate(path, req, body)
    if err != nil {
      return err
    }
  }
  resp, err := client.Do(req)
  if err != nil {
		fmt.Println("ici maggle")
    return err
  }
	fmt.Println("ici\n\n", resp, "\n")
  err = json.NewDecoder(resp.Body).Decode(result)
  if err != nil {
    return err
  }
  return nil
}

// Authenticate works with the Fetch call and adds certain Headers
// to the http request. This includes the actual API key and the
// timestamp of the request. Also a signature which is encoded
// with hmac and the API secret key.
func (a *APIClient) Authenticate(path string, req *http.Request, body interface{}) error {
  req.Header.Set("X-CC-Api-Key", a.Key)
  return nil
}
