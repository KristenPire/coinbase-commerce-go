# coinbase-commerce-go
Coinbase Commerce Golang

# Table of contents

<!--ts-->
   * [Documentation](#documentation)
   * [Installation](#installation)
   * [Usage](#usage)
      * [Checkouts](#checkouts)
      * [Charges](#charges)
      * [Events](#events)
<!--te-->


## Documentation

For more details visit [Coinbase API docs](https://commerce.coinbase.com/docs/api/).

To start using library, you need to register on [Commerce SignUp](https://commerce.coinbase.com/signup).
And get your ``API_KEY`` from user settings.

Next create a ``APIClient`` object for interacting with the API:
```golang
import "github.com/KristenPire/coinbase-commerce-go"

client := coinbase.Client(API_KEY)
```

``Client`` contains links to an every Golang Class representations of the API resources
``Checkout, Charge, Event``

You can call ``Create, List, Get, Update, Delete`` methods from an API resource classes

```golang
client.Charge.Create
client.Checkout.List 
client.Event.Get
client.Checkout.Update
client.Checkout.Delete
```
as well as ``Save, Delete, Refresh`` methods from API resource class instances.
```golang
checkout := client.Checkout.Get(<id>)
checkout.Refresh()
checkout.Save()
checkout.Delete()
```

Each API method returns an ``API ressource instances`` (``APICharge, APICheckout, APIEvent``) representing the response from the API, all of the models are dumpable with JSON.\
The response data is parsed into Golang objects, the appropriate ``APIObject`` subclasses will be used automatically.

Client support Common API Errors and Warnings handling.
All errors occuring during interaction with the API will be return.


| Error                    | Status Code |
|--------------------------|-------------|
| APIError                 |      *      |   
| InvalidRequestError      |     400     |   
| ParamRequiredError       |     400     |  
| ValidationError          |     400     |  
| AuthenticationError      |     401     |  
| ResourceNotFoundError    |     404     |
| RateLimitExceededError   |     429     |
| InternalServerError      |     500     |
| ServiceUnavailableError  |     503     |

## Installation

Install with ``go get``:

    go get "github.com/KristenPire/coinbase-commerce-go"


## Usage
```golang
import "github.com/KristenPire/coinbase-commerce-go"

client := coinbase.Client(API_KEY)
```
## Checkouts 
[Checkouts API docs](https://commerce.coinbase.com/docs/api/#checkouts)
### Get
```golang
checkout := client.Checkout.Get(<checkout_id>)
```
### Create
```golang
#by struct
checkout, err := client.Checkout.Create(coinbase.APICheckoutData{
    Name:"The Sovereign Individual",
    Description: "Mastering the Transition to the Information Age",
    Pricing_type: "fixed_price",
    Local_price: coinbase.Money{Amount : 100.00, Currency: "USD"},
    Requested_info: []string{"email", "name"},
   })

#or directly by json
checkout_info := `{
    "name": "The Sovereign Individual",
    "description": "Mastering the Transition to the Information Age",
    "pricing_type": "fixed_price",
    "local_price": {
        "amount": "100.00",
        "currency": "USD"
    },
    "requested_info": ["name", "email"]
}`
checkout, err := client.Checkout.Create(checkout_info)
```
### Update
```golang
#by object method
checkout := client.Checkout.Get(<checkout_id>)
checkout.Data.Name := "new name"
checkout.Save()

#by API method and json
checkout_info := `{"name": "newName"}`

checkout := client.Checkout.Update('<checkout_id>', checkout_info)

#or by API method and object
checkout := coinbase.APICheckoutData{}
checkout.Name := "new name"

checkout := client.Checkout.Update('<checkout_id>', checkout)
```

### Delete
```golang
#by object method
checkout := client.Checkout.Get(<checkout_id>)
checkout.Delete()

#by API method
client.Checkout.Delete('<checkout_id>')
```
### List
```golang
checkouts, err := client.Checkout.List()
```

### Iterations
```golang
checkouts, err := client.Checkout.List()
for err, checkout := range checkouts.Data{
    checkout.Delete()
}
```
## Charges
[Charges API docs](https://commerce.coinbase.com/docs/api/#charges)
### Retrieve
```golang
charge := client.Charge.Get(<charge_id>)
```
### Create
```golang
#by struct
charge, err := client.Charge.Create(coinbase.APIChargeData{
    Name:"The Sovereign Individual",
    Description: "Mastering the Transition to the Information Age",
    Pricing_type: "fixed_price",
    Local_price: coinbase.Money{Amount : 100.00, Currency: "USD"},
   })

#or directly by json
charge_info := `{
    "name": "The Sovereign Individual",
    "description": "Mastering the Transition to the Information Age",
    "pricing_type": "fixed_price",
    "local_price": {
        "amount": "100.00",
        "currency": "USD"title
    }
}`
charge, err := client.Charge.Create(charge_info)
```
### List
```golang
charges, err := client.Charge.List()
```
### Iterations
```golang
charges, err := client.Charge.List()
for _, charge := range charges.Data{
    jsonStr, _ := json.Marshal(charge)
    fmt.Println(string(jsonStr))
}
```
## Events
[Events API Docs](https://commerce.coinbase.com/docs/api/#events)
### Retrieve
```golang
event := client.Event.Get(<event_id>)
```
### List
```golang
events, err := client.Event.List()
```
### Iterations
```golang
events, err := client.Event.List()
for _, event := range events.Data{
    jsonStr, _ := json.Marshal(event)
    fmt.Println(string(jsonStr))
}
```

## Types
### Checkout
#### APICheckoutData
```golang
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
```
#### APICheckout
```golang
type APICheckout struct {
	father *ACheckout
	Data   APICheckoutData `json:"data,omitempty"`
	Errors []APIError      `json:"errors,omitempty"`
}
```
#### APICheckouts
```golang
type APICheckouts struct {
	Pagination APIPagination `json:"pagination,omitempty"`
	Checkouts  []APICheckout `json:"data,omitempty"`
	Errors     []APIError    `json:"errors,omitempty"`
}
```
### Charge
#### APIChargeData
```golang
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
	Payments []map[string]interface{} `json:"payements,omitempty"`
	Addresses struct {
		Bitcoin     string `json:"bitcoin,omitempty"`
		Bitcoincash string `json:"bitcoincash,omitempty"`
		Ethereum    string `json:"ethereum,omitempty"`
		Litecoin    string `json:"litecoin,omitempty"`
	} `json:"addresses,omitempty"`
	Local_price Money `json:"local_price,omitempty"`
}
```
#### APICharge
```golang
type APICharge struct {
	father *ACharge
	Data   APIChargeData `json:"data,omitempty"`
	Errors []APIError    `json:"errors,omitempty"`
}
```
#### APICharges
```golang
type APICharges struct {
	Pagination APIPagination `json:"pagination,omitempty"`
	Charges    []APICharge   `json:"data,omitempty"`
	Errors     []APIError    `json:"errors,omitempty"`
}
```

### Event
#### APIEventData
```golang

type APIEventData struct {
	Id          string        `json:"id,omitempty"`
	Resource    string        `json:"ressource,omitempty"`
	Created_at  *time.Time    `json:"created_at,omitempty"`
	Api_version string        `json:"api_version,omitempty"`
	Data        APIChargeData `json:"data,omitempty"`
}
```
#### APIEvent
```golang
type APIEvent struct {
	father *AEvent
	Data   APIEventData `json:"data,omitempty"`
	Errors []APIError   `json:"errors,omitempty"`
}
```
#### APIEvents
```golang
type APIEvents struct {
	Pagination APIPagination `json:"pagination,omitempty"`
	Events     []APIEvent    `json:"data,omitempty"`
	Errors     []APIError    `json:"errors,omitempty"`
}
```

### API
#### APIError
```golang
type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
```
#### APIPagination
```golang
type APIPagination struct {
	Order          string
	Starting_after string
	Ending_before  string
	Total          int
	Limit          int
	Previous_uri   string
	Next_uri       string
	Yielded        int
	Cursor_range   []string
}
```
#### Money
```golang
type Money struct {
	Amount   float64 `json:"amount,string,omitempty"`
	Currency string  `json:"currency,omitempty"`
}
```
