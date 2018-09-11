package coinbase

import "time"

type AEvent struct {
  Api	*APIClient
}

type APIEventData struct {
	Id          string        `json:"id,omitempty"`
	Resource    string        `json:"ressource,omitempty"`
	Created_at  *time.Time    `json:"created_at,omitempty"`
	Api_version string        `json:"api_version,omitempty"`
	Data        APIChargeData `json:"data,omitempty"`
}

type APIEvent struct {
	Data   APIEventData `json:"data,omitempty"`
	Errors []Error      `json:"errors,omitempty"`
}

type APIEvents struct {
	Pagination APIPagination  `json:"pagination,omitempty"`
	Data       []APIEventData `json:"data,omitempty"`
	Errors     []Error        `json:"errors,omitempty"`
}

func (a *APIClient) Event(id string) (event APIEvent, err error) {
	err = a.Fetch("GET", "/events/"+id, nil, &event)
	return
}

func (a *APIClient) Events() (events APIEvents, err error) {
	err = a.Fetch("GET", "/events/", nil, &events)
	return
}
