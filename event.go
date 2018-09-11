package coinbase

import "time"

type AEvent struct {
	Api *APIClient
}

type APIEventData struct {
	Id          string        `json:"id,omitempty"`
	Resource    string        `json:"ressource,omitempty"`
	Created_at  *time.Time    `json:"created_at,omitempty"`
	Api_version string        `json:"api_version,omitempty"`
	Data        APIChargeData `json:"data,omitempty"`
}

type APIEvent struct {
	father *AEvent
	Data   APIEventData `json:"data,omitempty"`
	Errors []APIError   `json:"errors,omitempty"`
}

type APIEvents struct {
	Pagination APIPagination  `json:"pagination,omitempty"`
	Data       []APIEventData `json:"data,omitempty"`
	Errors     []APIError     `json:"errors,omitempty"`
}

func (a *AEvent) Get(id string) (event APIEvent, err error) {
	err = a.Api.Fetch("GET", "/events/"+id, nil, &event)
	event.father = a
	return
}

func (a *APIEvent) Refresh() (err error) {
	err = a.father.Api.Fetch("GET", "/events/"+a.Data.Id, nil, a)
	return
}

func (a *AEvent) List() (events APIEvents, err error) {
	err = a.Api.Fetch("GET", "/events/", nil, &events)
	return
}
