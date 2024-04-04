package prestashop

import (
	"encoding/xml"
	"net/http"
)

type ResponseApi struct {
	XMLName xml.Name `xml:"prestashop" json:"prestashop,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Xlink   string   `xml:"xlink,attr" json:"xlink,omitempty"`
	Data    *ApiData `xml:"api" json:"api,omitempty"`
}

type ApiData struct {
	Text         string       `xml:",chardata" json:"text,omitempty"`
	ShopName     string       `xml:"shopName,attr" json:"shopname,omitempty"`
	Carriers     *ApiResource `xml:"carriers" json:"carriers,omitempty"`
	Carts        *ApiResource `xml:"carts" json:"carts,omitempty"`
	Currencies   *ApiResource `xml:"currencies" json:"currencies,omitempty"`
	Customers    *ApiResource `xml:"customers" json:"customers,omitempty"`
	Orders       *ApiResource `xml:"orders" json:"orders,omitempty"`
	OrderDetails *ApiResource `xml:"order_details" json:"order_details,omitempty"`
	OrderStates  *ApiResource `xml:"order_states" json:"order_states,omitempty"`
	Products     *ApiResource `xml:"products" json:"products,omitempty"`
}

type ApiResource struct {
	Text        string                  `xml:",chardata" json:"text,omitempty"`
	Href        string                  `xml:"href,attr" json:"href,omitempty"`
	Get         string                  `xml:"get,attr" json:"get,omitempty"`
	Put         string                  `xml:"put,attr" json:"put,omitempty"`
	Post        string                  `xml:"post,attr" json:"post,omitempty"`
	Patch       string                  `xml:"patch,attr" json:"patch,omitempty"`
	Delete      string                  `xml:"delete,attr" json:"delete,omitempty"`
	Head        string                  `xml:"head,attr" json:"head,omitempty"`
	Description *ApiResourceDescription `xml:"description" json:"description,omitempty"`
	Schema      []ApiResourceSchema `xml:"schema" json:"schema,omitempty"`
}

type ApiResourceDescription struct {
	Text   string `xml:",chardata" json:"text,omitempty"`
	Href   string `xml:"href,attr" json:"href,omitempty"`
	Get    string `xml:"get,attr" json:"get,omitempty"`
	Put    string `xml:"put,attr" json:"put,omitempty"`
	Post   string `xml:"post,attr" json:"post,omitempty"`
	Patch  string `xml:"patch,attr" json:"patch,omitempty"`
	Delete string `xml:"delete,attr" json:"delete,omitempty"`
	Head   string `xml:"head,attr" json:"head,omitempty"`
}

type ApiResourceSchema struct {
	Text string `xml:",chardata" json:"text,omitempty"`
	Href string `xml:"href,attr" json:"href,omitempty"`
	Type string `xml:"type,attr" json:"type,omitempty"`
}

func (client *Client) GetPermissions() (*ApiData, *http.Response, error) {
	req, _ := client.NewRequest("GET", "", nil)

	apiData := new(ApiData)
	apiResponse := new(ResponseApi)
	response, err := client.Do(req, apiResponse)

	if err != nil {
		return nil, response, err
	}

	if apiResponse.Data != nil {
		apiData = apiResponse.Data
	}

	return apiData, response, nil
}
