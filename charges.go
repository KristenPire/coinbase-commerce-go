package coinbase

import "time"
// ACharge is a class hosted inside the APIClient providing methods about charge
type ACharge struct {
	Api *APIClient
}

// APIChargeData is the golang struct equivalent of the Charge resource. It's findable inside APICharge object
type APIChargeData struct {
	Id           string     `json:"id,omitempty"`
	Ressource     string     `json:"ressource,omitempty"`
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
	Payments []struct {
		Network string`json:"network,omitepty"`
		Transaction_id string `json:"transaction_id,omitepty"`
		Status string `json:"status,omitepty"`
		Value struct {
			Local struct {
				Amount string `json:"amount,omitepty"`
				Currency string `json:"currency,omitepty"`
			}
			Crypto struct {
				Amount string `json:"amount,omitepty"`
				Currency string `json:"currency,omitepty"`
			}`json:"crypto,omitepty"`
		} `json:"value,omitepty"`
		Block struct {
			Height int `json:"height,omitepty"`
			Hash string `json:"hash,omitepty"`
			Confirmations_accumulated int `json:"confirmations_accumulated ,omitepty"`
			Confirmations_required int `json:"confirmations_required,omitepty"`
		} `json:"block,omitepty"`
	} `json:"payments,omitempty"`
	Addresses struct {
		Bitcoin     string `json:"bitcoin,omitempty"`
		Bitcoincash string `json:"bitcoincash,omitempty"`
		Ethereum    string `json:"ethereum,omitempty"`
		Litecoin    string `json:"litecoin,omitempty"`
	} `json:"addresses,omitempty"`
	Local_price Money `json:"local_price,omitempty"`
}

// APICharge is the object API object returned by the api routes.
type APICharge struct {
	father *ACharge
	Data   APIChargeData `json:"data,omitempty"`
	Errors []APIError    `json:"errors,omitempty"`
}

// APICharges is the object API object filled with a list of APICharge object alongside of Pagination information and a list of Errors.
type APICharges struct {
	Pagination APIPagination `json:"pagination,omitempty"`
	Charges    []APICharge   `json:"data,omitempty"`
	Errors     []APIError    `json:"errors,omitempty"`
}

// APICharge is the golang struct equivalent of the List charges routes
type APIChargesRequest struct {
	Pagination APIPagination   `json:"pagination,omitempty"`
	Data       []APIChargeData `json:"data,omitempty"`
	Errors     []APIError      `json:"errors,omitempty"`
}

// Get provides the APICharge instance of the given charge id.
func (a *ACharge) Get(id string) (charge APICharge, err error) {
	err = a.Api.Fetch("GET", "/charges/"+id, nil, &charge)
	charge.father = a
	return
}

// Refresh will update attributes and all nested data by making a fresh GET request to the relevant API endpoint.
func (a *APICharge) Refresh() (err error) {
	err = a.father.Api.Fetch("GET", "/charges/"+a.Data.Id, nil, &a.Data)
	return
}

// List create a APICharges object with a list of APICharge instance.
func (a *ACharge) List() (charges APICharges, err error) {
	temp := APIChargesRequest{}
	err = a.Api.Fetch("GET", "/charges/", nil, &temp)
	charges.Pagination = temp.Pagination
	charges.Errors = temp.Errors
	for _, data := range temp.Data {
		charges.Charges = append(charges.Charges, APICharge{father: a, Data: data, Errors: temp.Errors})
	}
	return
}

// Create a new charge and return his golang instance
func (a *ACharge) Create(data interface{}) (charge APICharge, err error) {
	err = a.Api.Fetch("POST", "/charges/", data, &charge)
	charge.father = a
	return
}
