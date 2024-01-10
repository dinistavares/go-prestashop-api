package prestashop

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Order service
type OrderService service

type ResponseOrder struct {
	XMLName xml.Name `xml:"prestashop" json:"prestashop,omitempty"`
	Xlink   string   `xml:"xlink,attr" json:"xlink,omitempty"`
	OrderData
}

type OrderData struct {
	Order  *Order   `xml:"order" json:"order,omitempty"`
	Orders *[]Order `xml:"orders" json:"orders,omitempty"`
}

type Order struct {
	ID                    int                `xml:"id" json:"id,omitempty"`
	IDAddressDelivery     int                `xml:"id_address_delivery" json:"id_address_delivery,omitempty"`
	IDAddressInvoice      int                `xml:"id_address_invoice" json:"id_address_invoice,omitempty"`
	IDCart                int                `xml:"id_cart" json:"id_cart,omitempty"`
	IDCurrency            int                `xml:"id_currency" json:"id_currency,omitempty"`
	IDLang                int                `xml:"id_lang" json:"id_lang,omitempty"`
	IDCustomer            int                `xml:"id_customer" json:"id_customer,omitempty"`
	IDCarrier             int                `xml:"id_carrier" json:"id_carrier,omitempty"`
	CurrentState          int                `xml:"current_state" json:"current_state,omitempty"`
	Module                string             `xml:"module" json:"module,omitempty"`
	InvoiceNumber         int                `xml:"invoice_number" json:"invoice_number,omitempty"`
	InvoiceDate           string             `xml:"invoice_date" json:"invoice_date,omitempty"`
	DeliveryNumber        int                `xml:"delivery_number" json:"delivery_number,omitempty"`
	DeliveryDate          string             `xml:"delivery_date" json:"delivery_date,omitempty"`
	Valid                 string             `xml:"valid" json:"valid,omitempty"`
	DateAdd               string             `xml:"date_add" json:"date_add,omitempty"`
	DateUpd               string             `xml:"date_upd" json:"date_upd,omitempty"`
	ShippingNumber        string             `xml:"shipping_number" json:"shipping_number,omitempty"`
	Note                  string             `xml:"note" json:"note,omitempty"`
	IDShopGroup           int                `xml:"id_shop_group" json:"id_shop_group,omitempty"`
	IDShop                int                `xml:"id_shop" json:"id_shop,omitempty"`
	SecureKey             string             `xml:"secure_key" json:"secure_key,omitempty"`
	Payment               string             `xml:"payment" json:"payment,omitempty"`
	Recyclable            string             `xml:"recyclable" json:"recyclable,omitempty"`
	Gift                  string             `xml:"gift" json:"gift,omitempty"`
	GiftMessage           string             `xml:"gift_message" json:"gift_message,omitempty"`
	MobileTheme           string             `xml:"mobile_theme" json:"mobile_theme,omitempty"`
	TotalDiscounts        string             `xml:"total_discounts" json:"total_discounts,omitempty"`
	TotalDiscountsTaxIncl string             `xml:"total_discounts_tax_incl" json:"total_discounts_tax_incl,omitempty"`
	TotalDiscountsTaxExcl string             `xml:"total_discounts_tax_excl" json:"total_discounts_tax_excl,omitempty"`
	TotalPaid             string             `xml:"total_paid" json:"total_paid,omitempty"`
	TotalPaidTaxIncl      string             `xml:"total_paid_tax_incl" json:"total_paid_tax_incl,omitempty"`
	TotalPaidTaxExcl      string             `xml:"total_paid_tax_excl" json:"total_paid_tax_excl,omitempty"`
	TotalPaidReal         string             `xml:"total_paid_real" json:"total_paid_real,omitempty"`
	TotalProducts         string             `xml:"total_products" json:"total_products,omitempty"`
	TotalProductsWt       string             `xml:"total_products_wt" json:"total_products_wt,omitempty"`
	TotalShipping         string             `xml:"total_shipping" json:"total_shipping,omitempty"`
	TotalShippingTaxIncl  string             `xml:"total_shipping_tax_incl" json:"total_shipping_tax_incl,omitempty"`
	TotalShippingTaxExcl  string             `xml:"total_shipping_tax_excl" json:"total_shipping_tax_excl,omitempty"`
	CarrierTaxRate        string             `xml:"carrier_tax_rate" json:"carrier_tax_rate,omitempty"`
	TotalWrapping         string             `xml:"total_wrapping" json:"total_wrapping,omitempty"`
	TotalWrappingTaxIncl  string             `xml:"total_wrapping_tax_incl" json:"total_wrapping_tax_incl,omitempty"`
	TotalWrappingTaxExcl  string             `xml:"total_wrapping_tax_excl" json:"total_wrapping_tax_excl,omitempty"`
	RoundMode             int                `xml:"round_mode" json:"round_mode,omitempty"`
	RoundType             int                `xml:"round_type" json:"round_type,omitempty"`
	ConversionRate        string             `xml:"conversion_rate" json:"conversion_rate,omitempty"`
	Reference             string             `xml:"reference" json:"reference,omitempty"`
	Associations          *OrderAssociations `xml:"associations" json:"associations,omitempty"`
}

type OrderRows struct {
	ID                 int    `xml:"id" json:"id,omitempty"`
	ProductID          int    `xml:"product_id" json:"product_id,omitempty"`
	ProductAttributeID int    `xml:"product_attribute_id" json:"product_attribute_id,omitempty"`
	ProductQuantity    int    `xml:"product_quantity" json:"product_quantity,omitempty"`
	ProductName        string `xml:"product_name" json:"product_name,omitempty"`
	ProductReference   string `xml:"product_reference" json:"product_reference,omitempty"`
	ProductEan13       string `xml:"product_ean13" json:"product_ean13,omitempty"`
	ProductIsbn        string `xml:"product_isbn" json:"product_isbn,omitempty"`
	ProductUpc         string `xml:"product_upc" json:"product_upc,omitempty"`
	ProductPrice       string `xml:"product_price" json:"product_price,omitempty"`
	IDCustomization    int    `xml:"id_customization" json:"id_customization,omitempty"`
	UnitPriceTaxIncl   string `xml:"unit_price_tax_incl" json:"unit_price_tax_incl,omitempty"`
	UnitPriceTaxExcl   string `xml:"unit_price_tax_excl" json:"unit_price_tax_excl,omitempty"`
}

type OrderAssociations struct {
	OrderRows *[]OrderRows `xml:"order_rows" json:"order_rows,omitempty"`
}

func (service *OrderService) Get(orderID int, params *ServiceListParams) (*Order, *http.Response, error) {
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
		if ordersResponse.Orders != nil && len(*ordersResponse.Orders) > 0 {
			order = &(*ordersResponse.Orders)[0]
		}
	}

	return order, response, nil
}

func (service *OrderService) List(params *ServiceListParams) (*[]Order, *http.Response, error) {
	_url := makeResourceUrl("orders", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	orders := new(ResponseOrder)
	response, err := service.client.Do(req, orders)

	if err != nil {
		// API returns 200 but the response is not a JSON object, return no customers found
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value of type prestashop.ResponseOrder") {
			return nil, response, errors.New("no orders found")
		}

		return nil, response, err
	}

	return orders.Orders, response, nil
}

func (service *OrderService) GetOrdersByCustomerID(customerID int, params *ServiceListParams) (*[]Order, *http.Response, error) {
	searchParams := ServiceListParams{
		Display: &ServiceListDisplay{
			"full",
		},
		Filter: &ServiceListFilter{
			Key:      "id_customer",
			Values:   []string{fmt.Sprintf("%d", customerID)},
			Operator: ListFilterOperatorLiteral,
		},
		// Set defined sort and limit params
		Limit: params.Limit,
		Sort:  params.Sort,
	}

	// Override display params
	if params.Display != nil {
		searchParams.Display = params.Display
	}

	orders, response, err := service.List(&searchParams)

	if err != nil {
		return nil, response, err
	}

	return orders, response, err
}
