package prestashop

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Order service
type OrderService service

type ResponseOrder struct {
	Order  *Order   `json:"order,omitempty"`
	Orders *[]Order `json:"orders,omitempty"`
}

type Order struct {
	ID                    int                `json:"id,omitempty"`
	IDAddressDelivery     int                `json:"id_address_delivery,omitempty"`
	IDAddressInvoice      int                `json:"id_address_invoice,omitempty"`
	IDCart                int                `json:"id_cart,omitempty"`
	IDCurrency            int                `json:"id_currency,omitempty"`
	IDLang                int                `json:"id_lang,omitempty"`
	IDCustomer            int                `json:"id_customer,omitempty"`
	IDCarrier             int                `json:"id_carrier,omitempty"`
	CurrentState          int                `json:"current_state,omitempty"`
	Module                string             `json:"module,omitempty"`
	InvoiceNumber         int                `json:"invoice_number,omitempty"`
	InvoiceDate           string             `json:"invoice_date,omitempty"`
	DeliveryNumber        int                `json:"delivery_number,omitempty"`
	DeliveryDate          string             `json:"delivery_date,omitempty"`
	Valid                 string             `json:"valid,omitempty"`
	DateAdd               string             `json:"date_add,omitempty"`
	DateUpd               string             `json:"date_upd,omitempty"`
	ShippingNumber        string             `json:"shipping_number,omitempty"`
	Note                  string             `json:"note,omitempty"`
	IDShopGroup           int                `json:"id_shop_group,omitempty"`
	IDShop                int                `json:"id_shop,omitempty"`
	SecureKey             string             `json:"secure_key,omitempty"`
	Payment               string             `json:"payment,omitempty"`
	Recyclable            string             `json:"recyclable,omitempty"`
	Gift                  string             `json:"gift,omitempty"`
	GiftMessage           string             `json:"gift_message,omitempty"`
	MobileTheme           string             `json:"mobile_theme,omitempty"`
	TotalDiscounts        string             `json:"total_discounts,omitempty"`
	TotalDiscountsTaxIncl string             `json:"total_discounts_tax_incl,omitempty"`
	TotalDiscountsTaxExcl string             `json:"total_discounts_tax_excl,omitempty"`
	TotalPaid             string             `json:"total_paid,omitempty"`
	TotalPaidTaxIncl      string             `json:"total_paid_tax_incl,omitempty"`
	TotalPaidTaxExcl      string             `json:"total_paid_tax_excl,omitempty"`
	TotalPaidReal         string             `json:"total_paid_real,omitempty"`
	TotalProducts         string             `json:"total_products,omitempty"`
	TotalProductsWt       string             `json:"total_products_wt,omitempty"`
	TotalShipping         string             `json:"total_shipping,omitempty"`
	TotalShippingTaxIncl  string             `json:"total_shipping_tax_incl,omitempty"`
	TotalShippingTaxExcl  string             `json:"total_shipping_tax_excl,omitempty"`
	CarrierTaxRate        string             `json:"carrier_tax_rate,omitempty"`
	TotalWrapping         string             `json:"total_wrapping,omitempty"`
	TotalWrappingTaxIncl  string             `json:"total_wrapping_tax_incl,omitempty"`
	TotalWrappingTaxExcl  string             `json:"total_wrapping_tax_excl,omitempty"`
	RoundMode             int                `json:"round_mode,omitempty"`
	RoundType             int                `json:"round_type,omitempty"`
	ConversionRate        string             `json:"conversion_rate,omitempty"`
	Reference             string             `json:"reference,omitempty"`
	Associations          *OrderAssociations `json:"associations,omitempty"`
}

type OrderRows struct {
	ID                 int    `json:"id,omitempty"`
	ProductID          int    `json:"product_id,omitempty"`
	ProductAttributeID int    `json:"product_attribute_id,omitempty"`
	ProductQuantity    int    `json:"product_quantity,omitempty"`
	ProductName        string `json:"product_name,omitempty"`
	ProductReference   string `json:"product_reference,omitempty"`
	ProductEan13       string `json:"product_ean13,omitempty"`
	ProductIsbn        string `json:"product_isbn,omitempty"`
	ProductUpc         string `json:"product_upc,omitempty"`
	ProductPrice       string `json:"product_price,omitempty"`
	IDCustomization    int    `json:"id_customization,omitempty"`
	UnitPriceTaxIncl   string `json:"unit_price_tax_incl,omitempty"`
	UnitPriceTaxExcl   string `json:"unit_price_tax_excl,omitempty"`
}

type OrderAssociations struct {
	OrderRows *[]OrderRows `json:"order_rows,omitempty"`
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
