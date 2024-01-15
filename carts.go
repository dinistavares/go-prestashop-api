package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Cart service
type CartService service

type ResponseCart struct {
	XMLName   xml.Name  `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink     string    `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	Cart      *Cart      `xml:"cart,omitempty" json:"cart,omitempty"`
	CartsData *CartsData `xml:"carts,omitempty" json:"carts,omitempty"`
}

type CartsData struct {
	Carts *[]Cart `xml:"cart,omitempty" json:"cart,omitempty"`
}

type Cart struct {
	ID                    *int              `xml:"id,omitempty" json:"id,omitempty"`
	IDAddressDelivery     int              `xml:"id_address_delivery,omitempty" json:"id_address_delivery,omitempty"`
	IDAddressInvoice      int              `xml:"id_address_invoice,omitempty" json:"id_address_invoice,omitempty"`
	IDCurrency            int              `xml:"id_currency,omitempty" json:"id_currency,omitempty"`
	IDCustomer            int              `xml:"id_customer,omitempty" json:"id_customer,omitempty"`
	IDGuest               int              `xml:"id_guest,omitempty" json:"id_guest,omitempty"`
	IDLang                int              `xml:"id_lang,omitempty" json:"id_lang,omitempty"`
	IDShopGroup           int              `xml:"id_shop_group,omitempty" json:"id_shop_group,omitempty"`
	IDShop                int              `xml:"id_shop,omitempty" json:"id_shop,omitempty"`
	IDCarrier             int              `xml:"id_carrier,omitempty" json:"id_carrier,omitempty"`
	Recyclable            string           `xml:"recyclable,omitempty" json:"recyclable,omitempty"`
	Gift                  string           `xml:"gift,omitempty" json:"gift,omitempty"`
	GiftMessage           string           `xml:"gift_message,omitempty" json:"gift_message,omitempty"`
	MobileTheme           string           `xml:"mobile_theme,omitempty" json:"mobile_theme,omitempty"`
	DeliveryOption        string           `xml:"delivery_option,omitempty" json:"delivery_option,omitempty"`
	SecureKey             string           `xml:"secure_key,omitempty" json:"secure_key,omitempty"`
	AllowSeperatedPackage string           `xml:"allow_seperated_package,omitempty" json:"allow_seperated_package,omitempty"`
	DateAdd               string           `xml:"date_add,omitempty" json:"date_add,omitempty"`
	DateUpd               string           `xml:"date_upd,omitempty" json:"date_upd,omitempty"`
	Associations          *CartAssociations `xml:"associations,omitempty" json:"associations,omitempty"`
}

type CartRow struct {
	IDProduct          int `xml:"id_product,omitempty" json:"id_product,omitempty"`
	IDProductAttribute int `xml:"id_product_attribute,omitempty" json:"id_product_attribute,omitempty"`
	IDAddressDelivery  int `xml:"id_address_delivery,omitempty" json:"id_address_delivery,omitempty"`
	IDCustomization    int `xml:"id_customization,omitempty" json:"id_customization,omitempty"`
	Quantity           int `xml:"quantity,omitempty" json:"quantity,omitempty"`
}

type CartAssociations struct {
	CartRows *[]CartRows `xml:"cart_rows,omitempty" json:"cart_rows,omitempty"`
}

type CartRows struct {
	CartRow []CartRow `xml:"cart_row" json:"cart_row,omitempty"`
}

func (service *CartService) Create(cart *Cart) (*Cart, *http.Response, error) {
	createdCart := new(Cart)

	body := ResponseCart{
		Xlink: "http://www.w3.org/1999/xlink",
		Cart: cart,
	}

	_url := "carts"
	req, _ := service.client.NewRequest("POST", _url, body)

	cartsResponse := new(ResponseCart)
	response, err := service.client.Do(req, cartsResponse)

	if err != nil {
		return nil, response, err
	}

	if cartsResponse != nil && cartsResponse.Cart != nil {
		createdCart = cartsResponse.Cart
	}

	return createdCart, response, nil
}

func (service *CartService) Get(cartID int, params *ServiceListParams) (*Cart, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
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
		if cartsResponse.CartsData != nil && cartsResponse.CartsData.Carts != nil &&
			len(*cartsResponse.CartsData.Carts) > 0 {
			cart = &(*cartsResponse.CartsData.Carts)[0]
		}
	}

	return cart, response, nil
}

func (service *CartService) List(params *ServiceListParams) (*[]Cart, *http.Response, error) {
	carts := new([]Cart)

	_url := makeResourceUrl("carts", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	cartsResponse := new(ResponseCart)
	response, err := service.client.Do(req, cartsResponse)

	if err != nil {
		return nil, response, err
	}

	if cartsResponse != nil && cartsResponse.CartsData != nil &&
		cartsResponse.CartsData.Carts != nil {
		carts = cartsResponse.CartsData.Carts
	}

	return carts, response, nil
}

func (service *CartService) ListCartsByCustomerID(customerID int, params *ServiceListParams) (*[]Cart, *http.Response, error) {
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
