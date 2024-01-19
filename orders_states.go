package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// OrderState service
type OrderStateService service

type ResponseOrderState struct {
	XMLName         xml.Name        `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink           string          `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	OrderState      *OrderState     `xml:"order_state,omitempty" json:"order_state,omitempty"`
	OrderStatesData *OrderStateData `xml:"order_states,omitempty" json:"order_states,omitempty"`
}

type OrderStateData struct {
	OrderStates *[]OrderState `xml:"order_state,omitempty" json:"order_state,omitempty"`
}

type OrderStates struct {
	OrderState []OrderState `xml:"order_state,omitempty" json:"order_state,omitempty"`
}

type OrderState struct {
	ID          string        `xml:"id,omitempty" json:"id,omitempty"`
	Unremovable string        `xml:"unremovable,omitempty" json:"unremovable,omitempty"`
	Delivery    string        `xml:"delivery,omitempty" json:"delivery,omitempty"`
	Hidden      string        `xml:"hidden,omitempty" json:"hidden,omitempty"`
	SendEmail   string        `xml:"send_email,omitempty" json:"send_email,omitempty"`
	ModuleName  string        `xml:"module_name,omitempty" json:"module_name,omitempty"`
	Invoice     string        `xml:"invoice,omitempty" json:"invoice,omitempty"`
	Color       string        `xml:"color,omitempty" json:"color,omitempty"`
	Logable     string        `xml:"logable,omitempty" json:"logable,omitempty"`
	Shipped     string        `xml:"shipped,omitempty" json:"shipped,omitempty"`
	Paid        string        `xml:"paid,omitempty" json:"paid,omitempty"`
	PdfDelivery string        `xml:"pdf_delivery,omitempty" json:"pdf_delivery,omitempty"`
	PdfInvoice  string        `xml:"pdf_invoice,omitempty" json:"pdf_invoice,omitempty"`
	Deleted     string        `xml:"deleted,omitempty" json:"deleted,omitempty"`
	Name        *LanguageData `xml:"name,omitempty" json:"name,omitempty"`
	Template    *LanguageData `xml:"template,omitempty" json:"template,omitempty"`
}

func (service *OrderStateService) Create(order_state *OrderState) (*OrderState, *http.Response, error) {
	createdOrderState := new(OrderState)

	body := ResponseOrderState{
		Xlink:      "http://www.w3.org/1999/xlink",
		OrderState: order_state,
	}

	_url := "order_states"
	req, _ := service.client.NewRequest("POST", _url, body)

	order_statesResponse := new(ResponseOrderState)
	response, err := service.client.Do(req, order_statesResponse)

	if err != nil {
		return nil, response, err
	}

	if order_statesResponse != nil && order_statesResponse.OrderState != nil {
		createdOrderState = order_statesResponse.OrderState
	}

	return createdOrderState, response, nil
}

func (service *OrderStateService) Get(order_stateID int, params *ServiceListParams) (*OrderState, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("order_states/%d", order_stateID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	orderState := new(OrderState)
	orderStatesResponse := new(ResponseOrderState)
	response, err := service.client.Do(req, orderStatesResponse)

	if err != nil {
		return nil, response, err
	}

	if orderStatesResponse != nil {
		if orderStatesResponse.OrderState != nil {
			orderState = orderStatesResponse.OrderState
		}

		// Use fisrt matching order_state
		if orderStatesResponse.OrderStatesData != nil && orderStatesResponse.OrderStatesData.OrderStates != nil &&
			len(*orderStatesResponse.OrderStatesData.OrderStates) > 0 {
			orderState = &(*orderStatesResponse.OrderStatesData.OrderStates)[0]
		}
	}

	return orderState, response, nil
}

func (service *OrderStateService) List(params *ServiceListParams) (*[]OrderState, *http.Response, error) {
	orderStates := new([]OrderState)
	_url := makeResourceUrl("order_states", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	orderStatesResponse := new(ResponseOrderState)
	response, err := service.client.Do(req, orderStatesResponse)

	if err != nil {
		return nil, response, err
	}

	if orderStatesResponse != nil && orderStatesResponse.OrderStatesData != nil &&
		orderStatesResponse.OrderStatesData.OrderStates != nil {
		orderStates = orderStatesResponse.OrderStatesData.OrderStates
	}

	return orderStates, response, nil
}

