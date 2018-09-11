package coinbase

type ACheckout struct {
  Api	*APIClient
}

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
	father *ACheckout
	Data   APICheckoutData `json:"data,omitempty"`
	Errors []Error         `json:"errors,omitempty"`
}

type APICheckouts struct {
	father *ACheckout
	Pagination APIPagination     `json:"pagination,omitempty"`
	Data       []APICheckoutData `json:"data,omitempty"`
	Errors     []Error           `json:"errors,omitempty"`
}

func (a *ACheckout) Get(id string) (checkout APICheckout, err error) {
	err = a.Api.Fetch("GET", "/checkouts/"+id, nil, &checkout)
	checkout.father = a
	return
}

func (a *ACheckout) List() (checkouts APICheckouts, err error) {
	err = a.Api.Fetch("GET", "/checkouts/", nil, &checkouts)
	checkouts.father = a
	return
}

func (a *ACheckout) Create(data interface{}) (checkout APICheckout, err error) {
	err = a.Api.Fetch("POST", "/checkouts/", data, &checkout)
	checkout.father = a
	return
}

func (a *ACheckout) Update(id string, data interface{}) (checkout APICheckout, err error) {
	err = a.Api.Fetch("PUT", "/checkouts/"+id, data, &checkout)
	checkout.father = a
	return
}

func (a *ACheckout) Delete(id string) (checkout APICheckout, err error) {
	err = a.Api.Fetch("DELETE", "/checkouts/"+id, nil, &checkout)
	checkout.father = a
	return
}

func (a *APICheckout) Save() (err error) {
	err = a.father.Api.Fetch("PUT", "/checkouts/"+a.Data.Id, a.Data, a)
	return
}

func (a *APICheckout) Delete() (err error) {
	err = a.father.Api.Fetch("DELETE", "/checkouts/"+a.Data.Id, nil, a)
	return
}
