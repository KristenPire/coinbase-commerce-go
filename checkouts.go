package coinbase

type Money struct {
	Amount float64 `json:"amount,string,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type APICheckoutCreate struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Pricing_type  string `json:"pricing_type"`
	Local_price Money `json:"local_price"`
	Requested_info []string `json:"requested_info"`
}

type APICheckoutData struct {
	Id string `json:"id"`
	Resource string `json:"ressource"`
	Name string `json:"name"`
	Description string `json:"description"`
	Logo_url string `json:"logo_url"`
	Requested_info []string `json:"requested_info"`
	Pricing_type  string `json:"pricing_type"`
	Local_price Money `json:"local_price"`
}

type APICheckout struct {
	Data APICheckoutData `json:"data"`
	Errors []Error `json:"errors"`
}

type APICheckouts struct {
	Pagination APIPagination `json:"pagination"`
	Data []APICheckoutData `json:"data"`
	Errors []Error `json:"errors"`
}


func (a *APIClient) ListCheckouts() (checkouts APICheckouts, err error){
	err = a.Fetch("GET", "/checkouts/", nil, &checkouts)
	if err != nil {
		return
	}
	return
}

func (a *APIClient) CreateCheckout(data APICheckoutCreate) (checkout APICheckout, err error){
	err = a.Fetch("POST", "/checkouts/", data, &checkout)
	if err != nil {
		return
	}
	return
}

func (a *APIClient) UpdateCheckout(id string, data APICheckoutData) (checkout APICheckout, err error){
	err = a.Fetch("PUT", "/checkouts/" + id, data, &checkout)
	if err != nil {
		return
	}
	return
}

func (a *APIClient) ShowCheckout(id string) (checkout APICheckout, err error){
	err = a.Fetch("GET", "/checkouts/" + id, nil, &checkout)
	if err != nil {
		return
	}
	return
}

func (a *APIClient) DeleteCheckout(id string) (checkout APICheckout, err error){
	err = a.Fetch("DELETE", "/checkouts/" + id, nil, &checkout)
	if err != nil {
		return
	}
	return
}
