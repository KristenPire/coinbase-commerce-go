package coinbase

import (
	"fmt"
)

type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code		int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Code: %d type: %s\nmessage:%s",e.Code, e.Type, e.Message)
}
