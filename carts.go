package prestashop

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Cart service
type CartService service

type ResponseCart struct {
	Cart  *Cart   `json:"cart,omitempty"`
	Carts *[]Cart `json:"carts,omitempty"`
}

type Cart struct {
	ID                    int               `json:"id,omitempty"`
	IDAddressDelivery     int               `json:"id_address_delivery,omitempty"`
	IDAddressInvoice      int               `json:"id_address_invoice,omitempty"`
	IDCurrency            int               `json:"id_currency,omitempty"`
	IDCustomer            int               `json:"id_customer,omitempty"`
	IDGuest               int               `json:"id_guest,omitempty"`
	IDLang                int               `json:"id_lang,omitempty"`
	IDShopGroup           int               `json:"id_shop_group,omitempty"`
	IDShop                int               `json:"id_shop,omitempty"`
	IDCarrier             int               `json:"id_carrier,omitempty"`
	Recyclable            string            `json:"recyclable,omitempty"`
	Gift                  string            `json:"gift,omitempty"`
	GiftMessage           string            `json:"gift_message,omitempty"`
	MobileTheme           string            `json:"mobile_theme,omitempty"`
	DeliveryOption        string            `json:"delivery_option,omitempty"`
	SecureKey             string            `json:"secure_key,omitempty"`
	AllowSeperatedPackage string            `json:"allow_seperated_package,omitempty"`
	DateAdd               string            `json:"date_add,omitempty"`
	DateUpd               string            `json:"date_upd,omitempty"`
	Associations          *CartAssociations `json:"associations,omitempty"`
}

type CartRows struct {
	IDProduct          int `json:"id_product,omitempty"`
	IDProductAttribute int `json:"id_product_attribute,omitempty"`
	IDAddressDelivery  int `json:"id_address_delivery,omitempty"`
	IDCustomization    int `json:"id_customization,omitempty"`
	Quantity           int `json:"quantity,omitempty"`
}

type CartAssociations struct {
	CartRows *[]CartRows `json:"cart_rows,omitempty"`
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
