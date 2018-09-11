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
checkouts := client.Checkout.List()
```

### Iterations
```golang
checkouts := client.Checkout.List()
for _, checkout := range checkouts.Data{
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
charges := client.Charge.List()
```
### Iterations
```golang
charges := client.Charge.List()
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
events := client.Event.List()
```
### Iterations
```golang
events := client.Event.List()
for _, event := range events.Data{
    jsonStr, _ := json.Marshal(event)
    fmt.Println(string(jsonStr))
}
```
