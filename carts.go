package prestashop

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Cart service
type CartService service

type ResponseCart struct {
	XMLName xml.Name `xml:"prestashop" json:"prestashop,omitempty"`
	Xlink   string   `xml:"xlink,attr" json:"xlink,omitempty"`
	CartData
}

type CartData struct {
	Cart  *Cart   `xml:"cart" json:"cart,omitempty"`
	Carts *[]Cart `xml:"carts" json:"carts,omitempty"`
}

type Cart struct {
	ID                    int               `xml:"id" json:"id,omitempty"`
	IDAddressDelivery     int               `xml:"id_address_delivery" json:"id_address_delivery,omitempty"`
	IDAddressInvoice      int               `xml:"id_address_invoice" json:"id_address_invoice,omitempty"`
	IDCurrency            int               `xml:"id_currency" json:"id_currency,omitempty"`
	IDCustomer            int               `xml:"id_customer" json:"id_customer,omitempty"`
	IDGuest               int               `xml:"id_guest" json:"id_guest,omitempty"`
	IDLang                int               `xml:"id_lang" json:"id_lang,omitempty"`
	IDShopGroup           int               `xml:"id_shop_group" json:"id_shop_group,omitempty"`
	IDShop                int               `xml:"id_shop" json:"id_shop,omitempty"`
	IDCarrier             int               `xml:"id_carrier" json:"id_carrier,omitempty"`
	Recyclable            string            `xml:"recyclable" json:"recyclable,omitempty"`
	Gift                  string            `xml:"gift" json:"gift,omitempty"`
	GiftMessage           string            `xml:"gift_message" json:"gift_message,omitempty"`
	MobileTheme           string            `xml:"mobile_theme" json:"mobile_theme,omitempty"`
	DeliveryOption        string            `xml:"delivery_option" json:"delivery_option,omitempty"`
	SecureKey             string            `xml:"secure_key" json:"secure_key,omitempty"`
	AllowSeperatedPackage string            `xml:"allow_seperated_package" json:"allow_seperated_package,omitempty"`
	DateAdd               string            `xml:"date_add" json:"date_add,omitempty"`
	DateUpd               string            `xml:"date_upd" json:"date_upd,omitempty"`
	Associations          *CartAssociations `xml:"associations" json:"associations,omitempty"`
}

type CartRows struct {
	IDProduct          int `xml:"id_product" json:"id_product,omitempty"`
	IDProductAttribute int `xml:"id_product_attribute" json:"id_product_attribute,omitempty"`
	IDAddressDelivery  int `xml:"id_address_delivery" json:"id_address_delivery,omitempty"`
	IDCustomization    int `xml:"id_customization" json:"id_customization,omitempty"`
	Quantity           int `xml:"quantity" json:"quantity,omitempty"`
}

type CartAssociations struct {
	CartRows *[]CartRows `xml:"cart_rows" json:"cart_rows,omitempty"`
}

func (service *CartService) Get(cartID int, params *ServiceListParams) (*Cart, *http.Response, error) {
	resourceRoute := fmt.Sprintf("carts/%d", cartID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	cart := new(Cart)
	cartsResponse := new(ResponseCart)
	response, err := service.client.Do(req, cartsResponse)

	if err != nil {
		return nil, response, err
	}

	if cartsResponse != nil {
		if cartsResponse.Cart != nil {
			cart = cartsResponse.Cart
		}

		// Use fisrt matching cart
		if cartsResponse.Carts != nil && len(*cartsResponse.Carts) > 0 {
			cart = &(*cartsResponse.Carts)[0]
		}
	}

	return cart, response, nil
}

func (service *CartService) List(params *ServiceListParams) (*[]Cart, *http.Response, error) {
	_url := makeResourceUrl("carts", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	carts := new(ResponseCart)
	response, err := service.client.Do(req, carts)

	if err != nil {
		// API returns 200 but the response is not a JSON object, return no customers found
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value of type prestashop.ResponseCart") {
			return nil, response, errors.New("no carts found")
		}

		return nil, response, err
	}

	return carts.Carts, response, nil
}

func (service *CartService) GetCartsByCustomerID(customerID int, params *ServiceListParams) (*[]Cart, *http.Response, error) {
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

	carts, response, err := service.List(&searchParams)

	if err != nil {
		return nil, response, err
	}

	return carts, response, err
}
