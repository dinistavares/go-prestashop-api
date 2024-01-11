# go-prestashop-api

A PrestaShop API Golang Wrapper for the [Prestashop Webservice API)](https://devdocs.prestashop-project.org/8/webservice/)

# Install

```console
$ go get github.com/dinistavares/go-prestashop-api
```

# Usage

Create a new API client and authenticate with your Webservice key. You should specify the URL protocol by prefixing your domain name with `https://` or `http://`. Follow the [Prestashop documentation](https://devdocs.prestashop-project.org/8/webservice/tutorials/creating-access/) to enable the webservice [through the Prestashop user interface](https://devdocs.prestashop-project.org/8/webservice/tutorials/creating-access/#via-the-user-interface) or [programmatically](https://devdocs.prestashop-project.org/8/webservice/tutorials/creating-access/#programmatically) then create your API keys through the [Prestashop user interface](https://devdocs.prestashop-project.org/8/webservice/tutorials/creating-access/#via-the-user-interface-1) or [programmatically](https://devdocs.prestashop-project.org/8/webservice/tutorials/creating-access/#programatically).

```go
import (
  "github.com/dinistavares/go-prestashop-api"
)

func main(){
  shopURL := "https://example.com"
  key := "xxxxxxx"

  client, err := prestashop.New(shopURL)

  if err != nil {
    // handle error

    return
  }

  client.Authenticate(key)
}

```

The API routes are broken down into services, the supported services are:

- Carts `(Create, Get, List, ListCartsByCustomerID)`
- Customers `(Create, Get, GetCustomersByEmail, List)`
- Orders `(Create, Get, List, ListOrdersByCustomerID)`
- Products `(Create, Get, List)`

Get Customer by ID with all fields.

```go
func (client *prestashop.Client) listOrders() {
  customer, _, err := client.Customer.Get(2, nil)

  if err != nil {
    // Handle errors

    return
  }


  // ....
}

```

Get Customer by ID with specified fields.

```go
func (client *prestashop.Client) listOrders() {
  params := &prestashop.ServiceListParams{
    // Results should only show firstname & email
    Display: &prestashop.ServiceListDisplay{
      "firstname",
      "email"
    },
  }

  customer, _, err := client.Customer.Get(2, params)

  if err != nil {
    // Handle errors

    return
  }


  // ....
}

```

List customer with firstname 'bob'.

```go
func (client *prestashop.Client) listOrders() {
  params := &prestashop.ServiceListParams{
    // Results should only show customer id, firstname & email
    Display: &prestashop.ServiceListDisplay{
      "id",
      "firstname",
      "email"
    },
    // Filter customers with firstname as 'bob'
    Filter: &prestashop.ServiceListFilter{
      Key: "firstname",
      Values: []string{"bob"},
      Operator: prestashop.ListFilterOperatorLiteral,
    },
  }

  customer, _, err := client.Customer.List(params)

  if err != nil {
    // Handle errors

    return
  }


  // ....
}

```
