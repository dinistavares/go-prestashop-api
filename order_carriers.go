package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// OrderCarrier service
type OrderCarrierService service

type ResponseOrderCarrier struct {
	XMLName           xml.Name          `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink             string            `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	OrderCarrier      *OrderCarrier     `xml:"order_carrier,omitempty" json:"order_carrier,omitempty"`
	OrderCarriersData *OrderCarrierData `xml:"order_carriers,omitempty" json:"order_carriers,omitempty"`
}

type OrderCarrierData struct {
	OrderCarriers *[]OrderCarrier `xml:"order_carrier,omitempty" json:"order_carrier,omitempty"`
}

type OrderCarrier struct {
	ID                  int    `xml:"id,omitempty" json:"id,omitempty"`
	IDOrder             int    `xml:"id_order,omitempty" json:"id_order,omitempty"`
	IDCarrier           int    `xml:"id_carrier,omitempty" json:"id_carrier,omitempty"`
	IDOrderInvoice      int    `xml:"id_order_invoice,omitempty" json:"id_order_invoice,omitempty"`
	Weight              string `xml:"weight,omitempty" json:"weight,omitempty"`
	ShippingCostTaxExcl string `xml:"shipping_cost_tax_excl,omitempty" json:"shipping_cost_tax_excl,omitempty"`
	ShippingCostTaxIncl string `xml:"shipping_cost_tax_incl,omitempty" json:"shipping_cost_tax_incl,omitempty"`
	TrackingNumber      string `xml:"tracking_number,omitempty" json:"tracking_number,omitempty"`
	DateAdd             string `xml:"date_add,omitempty" json:"date_add,omitempty"`
}

func (service *OrderCarrierService) Create(orderCarrier *OrderCarrier) (*OrderCarrier, *http.Response, error) {
	createdOrderCarrier := new(OrderCarrier)

	body := ResponseOrderCarrier{
		Xlink:        "http://www.w3.org/1999/xlink",
		OrderCarrier: orderCarrier,
	}

	_url := "order_carriers"
	req, _ := service.client.NewRequest("POST", _url, body)

	orderCarriersResponse := new(ResponseOrderCarrier)
	response, err := service.client.Do(req, orderCarriersResponse)

	if err != nil {
		return nil, response, err
	}

	if orderCarriersResponse != nil && orderCarriersResponse.OrderCarrier != nil {
		createdOrderCarrier = orderCarriersResponse.OrderCarrier
	}

	return createdOrderCarrier, response, nil
}

func (service *OrderCarrierService) Get(orderCarrierID int, params *ServiceListParams) (*OrderCarrier, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("order_carriers/%d", orderCarrierID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	orderCarrier := new(OrderCarrier)
	orderCarriersResponse := new(ResponseOrderCarrier)
	response, err := service.client.Do(req, orderCarriersResponse)

	if err != nil {
		return nil, response, err
	}

	if orderCarriersResponse != nil {
		if orderCarriersResponse.OrderCarrier != nil {
			orderCarrier = orderCarriersResponse.OrderCarrier
		}

		// Use first matching order carrier
		if orderCarriersResponse.OrderCarriersData != nil && orderCarriersResponse.OrderCarriersData.OrderCarriers != nil &&
			len(*orderCarriersResponse.OrderCarriersData.OrderCarriers) > 0 {
			orderCarrier = &(*orderCarriersResponse.OrderCarriersData.OrderCarriers)[0]
		}
	}

	return orderCarrier, response, nil
}

func (service *OrderCarrierService) List(params *ServiceListParams) (*[]OrderCarrier, *http.Response, error) {
	orderCarriers := new([]OrderCarrier)
	_url := makeResourceUrl("order_carriers", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	orderCarriersResponse := new(ResponseOrderCarrier)
	response, err := service.client.Do(req, orderCarriersResponse)

	if err != nil {
		return nil, response, err
	}

	if orderCarriersResponse != nil && orderCarriersResponse.OrderCarriersData != nil &&
		orderCarriersResponse.OrderCarriersData.OrderCarriers != nil {
		orderCarriers = orderCarriersResponse.OrderCarriersData.OrderCarriers
	}

	return orderCarriers, response, nil
}

// ListByOrderID lists order carriers for a given order ID, returning full records.
func (service *OrderCarrierService) ListByOrderID(orderID string, params *ServiceListParams) (*[]OrderCarrier, *http.Response, error) {
	searchParams := ServiceListParams{
		Display: &ServiceListDisplay{
			"full",
		},
		Filter: &ServiceListFilter{
			Key:      "id_order",
			Values:   []string{orderID},
			Operator: ListFilterOperatorLiteral,
		},
	}

	if params != nil {
		if params.Display != nil {
			searchParams.Display = params.Display
		}

		if params.Limit != nil {
			searchParams.Limit = params.Limit
		}

		if params.Sort != nil {
			searchParams.Sort = params.Sort
		}
	}

	return service.List(&searchParams)
}
