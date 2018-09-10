package coinbase

import "time"

type APIChargeData struct {
	Id           string     `json:"id,omitempty"`
	Resource     string     `json:"ressource,omitempty"`
	Code         string     `json:"code,omitempty"`
	Name         string     `json:"name,omitempty"`
	Description  string     `json:"description,omitempty"`
	Logo_url     string     `json:"logo_url,omitempty"`
	Hosted_url   string     `json:"Hosted_url,omitempty"`
	Created_at   *time.Time `json:"created_at,omitempty"`
	Updated_at   *time.Time `json:"updated_at,omitempty"`
	Confirmed_at *time.Time `json:"confirmed_at,omitempty"`
	Checkout     struct {
		Id string `json:"id,omitempty"`
	} `json:"checkout,omitempty"`
	Timeline []struct {
		Time    *time.Time `json:"id,omitempty"`
		Status  string     `json:"status,omitempty"`
		Context string     `json:"context,omitempty"`
	} `json:"timeline,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Pricing_type string                 `json:"pricing_type,omitempty"`
	Pricing      struct {
		Local       Money `json:"local,omitempty"`
		Bitcoin     Money `json:"bitcoin,omitempty"`
		Bitcoincash Money `json:"bitcoincash,omitempty"`
		Ethereum    Money `json:"ethereum,omitempty"`
		Litecoin    Money `json:"litecoin,omitempty"`
	} `json:"pricing,omitempty"`
	Payements []map[string]interface{} `json:"payements,omitempty"`
	Addresses struct {
		Bitcoin     string `json:"bitcoin,omitempty"`
		Bitcoincash string `json:"bitcoincash,omitempty"`
		Ethereum    string `json:"ethereum,omitempty"`
		Litecoin    string `json:"litecoin,omitempty"`
	} `json:"addresses,omitempty"`
	Local_price Money `json:"local_price,omitempty"`
}

type APICharge struct {
	Data   APIChargeData `json:"data,omitempty"`
	Errors []Error       `json:"errors,omitempty"`
}

type APICharges struct {
	Pagination APIPagination   `json:"pagination,omitempty"`
	Data       []APIChargeData `json:"data,omitempty"`
	Errors     []Error         `json:"errors,omitempty"`
}

func (a *APIClient) Charge(id string) (charge APICharge, err error) {
	err = a.Fetch("GET", "/charges/"+id, nil, &charge)
	return
}

func (a *APIClient) Charges() (charges APICharges, err error) {
	err = a.Fetch("GET", "/charges/", nil, &charges)
	return
}

func (a *APIClient) CreateCharge(data APIChargeData) (charge APICharge, err error) {
	err = a.Fetch("POST", "/charges/", data, &charge)
	return
}
