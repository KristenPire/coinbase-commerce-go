package coinbase

import (
	"fmt"
)

type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("type: %s\nmessage:%s", e.Type, e.Message)
}
