package coinbase

type Money struct {
	Amount   float64 `json:"amount,string,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type APICheckoutData struct {
	Id             string   `json:"id,omitempty"`
	Resource       string   `json:"ressource,omitempty"`
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Logo_url       string   `json:"logo_url,omitempty"`
	Requested_info []string `json:"requested_info,omitempty"`
	Pricing_type   string   `json:"pricing_type,omitempty"`
	Local_price    Money    `json:"local_price,omitempty"`
}

type APICheckout struct {
	Data   APICheckoutData `json:"data,omitempty"`
	Errors []Error         `json:"errors,omitempty"`
}

type APICheckouts struct {
	Pagination APIPagination     `json:"pagination,omitempty"`
	Data       []APICheckoutData `json:"data,omitempty"`
	Errors     []Error           `json:"errors,omitempty"`
}

func (a *APIClient) Checkout(id string) (checkout APICheckout, err error) {
	err = a.Fetch("GET", "/checkouts/"+id, nil, &checkout)
	return
}

func (a *APIClient) Checkouts() (checkouts APICheckouts, err error) {
	err = a.Fetch("GET", "/checkouts/", nil, &checkouts)
	return
}

func (a *APIClient) CreateCheckout(data interface{}) (checkout APICheckout, err error) {
	err = a.Fetch("POST", "/checkouts/", data, &checkout)
	return
}

func (a *APIClient) UpdateCheckout(id string, data interface{}) (checkout APICheckout, err error) {
	err = a.Fetch("PUT", "/checkouts/"+id, data, &checkout)
	return
}

func (a *APIClient) DeleteCheckout(id string) (checkout APICheckout, err error) {
	err = a.Fetch("DELETE", "/checkouts/"+id, nil, &checkout)
	return
}
