package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Order service
type OrderService service

type ResponseOrder struct {
	XMLName    xml.Name   `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink      string     `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	Order      *Order     `xml:"order,omitempty" json:"order,omitempty"`
	OrdersData *OrderData `xml:"orders,omitempty" json:"orders,omitempty"`
}

type OrderData struct {
	Orders *[]Order `xml:"order,omitempty" json:"order,omitempty"`
}

type Order struct {
	ID                    int                `xml:"id,omitempty" json:"id,omitempty"`
	IDAddressDelivery     int                `xml:"id_address_delivery,omitempty" json:"id_address_delivery,omitempty"`
	IDAddressInvoice      int                `xml:"id_address_invoice,omitempty" json:"id_address_invoice,omitempty"`
	IDCart                int                `xml:"id_cart,omitempty" json:"id_cart,omitempty"`
	IDCurrency            int                `xml:"id_currency,omitempty" json:"id_currency,omitempty"`
	IDLang                int                `xml:"id_lang,omitempty" json:"id_lang,omitempty"`
	IDCustomer            int                `xml:"id_customer,omitempty" json:"id_customer,omitempty"`
	IDCarrier             int                `xml:"id_carrier,omitempty" json:"id_carrier,omitempty"`
	CurrentState          int                `xml:"current_state,omitempty" json:"current_state,omitempty"`
	Module                string             `xml:"module,omitempty" json:"module,omitempty"`
	InvoiceNumber         int                `xml:"invoice_number,omitempty" json:"invoice_number,omitempty"`
	InvoiceDate           string             `xml:"invoice_date,omitempty" json:"invoice_date,omitempty"`
	DeliveryNumber        int                `xml:"delivery_number,omitempty" json:"delivery_number,omitempty"`
	DeliveryDate          string             `xml:"delivery_date,omitempty" json:"delivery_date,omitempty"`
	Valid                 string             `xml:"valid,omitempty" json:"valid,omitempty"`
	DateAdd               string             `xml:"date_add,omitempty" json:"date_add,omitempty"`
	DateUpdated           string             `xml:"date_upd,omitempty" json:"date_updated,omitempty"`
	ShippingNumber        string             `xml:"shipping_number,omitempty" json:"shipping_number,omitempty"`
	Note                  string             `xml:"note,omitempty" json:"note,omitempty"`
	IDShopGroup           int                `xml:"id_shop_group,omitempty" json:"id_shop_group,omitempty"`
	IDShop                int                `xml:"id_shop,omitempty" json:"id_shop,omitempty"`
	SecureKey             string             `xml:"secure_key,omitempty" json:"secure_key,omitempty"`
	Payment               string             `xml:"payment,omitempty" json:"payment,omitempty"`
	Recyclable            string             `xml:"recyclable,omitempty" json:"recyclable,omitempty"`
	Gift                  string             `xml:"gift,omitempty" json:"gift,omitempty"`
	GiftMessage           string             `xml:"gift_message,omitempty" json:"gift_message,omitempty"`
	MobileTheme           string             `xml:"mobile_theme,omitempty" json:"mobile_theme,omitempty"`
	TotalDiscounts        float64            `xml:"total_discounts,omitempty" json:"total_discounts,omitempty"`
	TotalDiscountsTaxIncl float64            `xml:"total_discounts_tax_incl,omitempty" json:"total_discounts_tax_incl,omitempty"`
	TotalDiscountsTaxExcl float64            `xml:"total_discounts_tax_excl,omitempty" json:"total_discounts_tax_excl,omitempty"`
	TotalPaid             float64            `xml:"total_paid,omitempty" json:"total_paid,omitempty"`
	TotalPaidTaxIncl      float64            `xml:"total_paid_tax_incl,omitempty" json:"total_paid_tax_incl,omitempty"`
	TotalPaidTaxExcl      float64            `xml:"total_paid_tax_excl,omitempty" json:"total_paid_tax_excl,omitempty"`
	TotalPaidReal         float64            `xml:"total_paid_real,omitempty" json:"total_paid_real,omitempty"`
	TotalProducts         float64            `xml:"total_products,omitempty" json:"total_products,omitempty"`
	TotalProductsWt       float64            `xml:"total_products_wt,omitempty" json:"total_products_wt,omitempty"`
	TotalShipping         float64            `xml:"total_shipping,omitempty" json:"total_shipping,omitempty"`
	TotalShippingTaxIncl  float64            `xml:"total_shipping_tax_incl,omitempty" json:"total_shipping_tax_incl,omitempty"`
	TotalShippingTaxExcl  float64            `xml:"total_shipping_tax_excl,omitempty" json:"total_shipping_tax_excl,omitempty"`
	CarrierTaxRate        float64            `xml:"carrier_tax_rate,omitempty" json:"carrier_tax_rate,omitempty"`
	TotalWrapping         float64            `xml:"total_wrapping,omitempty" json:"total_wrapping,omitempty"`
	TotalWrappingTaxIncl  float64            `xml:"total_wrapping_tax_incl,omitempty" json:"total_wrapping_tax_incl,omitempty"`
	TotalWrappingTaxExcl  float64            `xml:"total_wrapping_tax_excl,omitempty" json:"total_wrapping_tax_excl,omitempty"`
	RoundMode             int                `xml:"round_mode,omitempty" json:"round_mode,omitempty"`
	RoundType             int                `xml:"round_type,omitempty" json:"round_type,omitempty"`
	ConversionRate        float64            `xml:"conversion_rate,omitempty" json:"conversion_rate,omitempty"`
	Reference             string             `xml:"reference,omitempty" json:"reference,omitempty"`
	Associations          *OrderAssociations `xml:"associations,omitempty" json:"associations,omitempty"`
}

type OrderRow struct {
	ID                 int     `xml:"id,omitempty" json:"id,omitempty"`
	ProductID          int     `xml:"product_id,omitempty" json:"product_id,omitempty"`
	ProductAttributeID int     `xml:"product_attribute_id,omitempty" json:"product_attribute_id,omitempty"`
	ProductQuantity    int     `xml:"product_quantity,omitempty" json:"product_quantity,omitempty"`
	ProductName        string  `xml:"product_name,omitempty" json:"product_name,omitempty"`
	ProductReference   string  `xml:"product_reference,omitempty" json:"product_reference,omitempty"`
	ProductEan13       string  `xml:"product_ean13,omitempty" json:"product_ean13,omitempty"`
	ProductIsbn        string  `xml:"product_isbn,omitempty" json:"product_isbn,omitempty"`
	ProductUpc         string  `xml:"product_upc,omitempty" json:"product_upc,omitempty"`
	ProductPrice       string  `xml:"product_price,omitempty" json:"product_price,omitempty"`
	IDCustomization    int     `xml:"id_customization,omitempty" json:"id_customization,omitempty"`
	UnitPriceTaxIncl   float64 `xml:"unit_price_tax_incl,omitempty" json:"unit_price_tax_incl,omitempty"`
	UnitPriceTaxExcl   float64 `xml:"unit_price_tax_excl,omitempty" json:"unit_price_tax_excl,omitempty"`
}

type OrderAssociations struct {
	OrderRows *[]OrderRow `xml:"order_rows>order_row,omitempty" json:"order_rows,omitempty"`
}

func (service *OrderService) Create(order *Order) (*Order, *http.Response, error) {
	createdOrder := new(Order)

	body := ResponseOrder{
		Xlink: "http://www.w3.org/1999/xlink",
		Order: order,
	}

	_url := "orders"
	req, _ := service.client.NewRequest("POST", _url, body)

	ordersResponse := new(ResponseOrder)
	response, err := service.client.Do(req, ordersResponse)

	if err != nil {
		return nil, response, err
	}

	if ordersResponse != nil && ordersResponse.Order != nil {
		createdOrder = ordersResponse.Order
	}

	return createdOrder, response, nil
}

func (service *OrderService) Get(orderID int, params *ServiceListParams) (*Order, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("orders/%d", orderID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	order := new(Order)
	ordersResponse := new(ResponseOrder)
	response, err := service.client.Do(req, ordersResponse)

	if err != nil {
		return nil, response, err
	}

	if ordersResponse != nil {
		if ordersResponse.Order != nil {
			order = ordersResponse.Order
		}

		// Use fisrt matching order
		if ordersResponse.OrdersData != nil && ordersResponse.OrdersData.Orders != nil &&
			len(*ordersResponse.OrdersData.Orders) > 0 {
			order = &(*ordersResponse.OrdersData.Orders)[0]
		}
	}

	return order, response, nil
}

func (service *OrderService) List(params *ServiceListParams) (*[]Order, *http.Response, error) {
	orders := new([]Order)
	_url := makeResourceUrl("orders", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	ordersResponse := new(ResponseOrder)
	response, err := service.client.Do(req, ordersResponse)

	if err != nil {
		return nil, response, err
	}

	if ordersResponse != nil && ordersResponse.OrdersData != nil &&
		ordersResponse.OrdersData.Orders != nil {
		orders = ordersResponse.OrdersData.Orders
	}

	return orders, response, nil
}

func (service *OrderService) ListOrdersByCustomerID(customerID int, params *ServiceListParams) (*[]Order, *http.Response, error) {
	searchParams := ServiceListParams{
		Display: &ServiceListDisplay{
			"full",
		},
		Filter: &ServiceListFilter{
			Key:      "id_customer",
			Values:   []string{fmt.Sprintf("%d", customerID)},
			Operator: ListFilterOperatorLiteral,
		},
	}

	if params != nil {
		// Override display params
		if params.Display != nil {
			searchParams.Display = params.Display
		}

		// Set limits if defined
		if params.Limit != nil {
			searchParams.Limit = params.Limit
		}

		// Set sort order if defined
		if params.Sort != nil {
			searchParams.Sort = params.Sort
		}
	}

	orders, response, err := service.List(&searchParams)

	if err != nil {
		return nil, response, err
	}

	return orders, response, err
}
