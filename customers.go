package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Customer service
type CustomersService service

type ResponseCustomer struct {
	XMLName       xml.Name       `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink         string         `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	Customer      *Customer      `xml:"customer,omitempty" json:"customer,omitempty"`
	CustomersData *CustomersData `xml:"customers,omitempty" json:"customers,omitempty"`
}

type CustomersData struct {
	Customers *[]Customer `xml:"customer,omitempty" json:"customers,omitempty"`
}

type Customer struct {
	ID                       *int                  `xml:"id,omitempty" json:"id,omitempty"`
	IDDefaultGroup           int                   `xml:"id_default_group,omitempty" json:"id_default_group,omitempty"`
	IDLang                   int                   `xml:"id_lang,omitempty" json:"id_lang,omitempty"`
	NewsletterDateAdd        string                `xml:"newsletter_date_add,omitempty" json:"newsletter_date_add,omitempty"`
	IPRegistrationNewsletter string                `xml:"ip_registration_newsletter,omitempty" json:"ip_registration_newsletter,omitempty"`
	LastPasswdGen            string                `xml:"last_passwd_gen,omitempty" json:"last_passwd_gen,omitempty"`
	SecureKey                string                `xml:"secure_key,omitempty" json:"secure_key,omitempty"`
	Deleted                  string                `xml:"deleted,omitempty" json:"deleted,omitempty"`
	Passwd                   string                `xml:"passwd,omitempty" json:"passwd,omitempty"`
	Lastname                 string                `xml:"lastname,omitempty" json:"lastname,omitempty"`
	Firstname                string                `xml:"firstname,omitempty" json:"firstname,omitempty"`
	Email                    string                `xml:"email,omitempty" json:"email,omitempty"`
	IDGender                 int                   `xml:"id_gender,omitempty" json:"id_gender,omitempty"`
	Birthday                 string                `xml:"birthday,omitempty" json:"birthday,omitempty"`
	Newsletter               string                `xml:"newsletter,omitempty" json:"newsletter,omitempty"`
	Optin                    string                `xml:"optin,omitempty" json:"optin,omitempty"`
	Website                  string                `xml:"website,omitempty" json:"website,omitempty"`
	Company                  string                `xml:"company,omitempty" json:"company,omitempty"`
	Siret                    string                `xml:"siret,omitempty" json:"siret,omitempty"`
	Ape                      string                `xml:"ape,omitempty" json:"ape,omitempty"`
	OutstandingAllowAmount   string                `xml:"outstanding_allow_amount,omitempty" json:"outstanding_allow_amount,omitempty"`
	ShowPublicPrices         string                `xml:"show_public_prices,omitempty" json:"show_public_prices,omitempty"`
	IDRisk                   int                   `xml:"id_risk,omitempty" json:"id_risk,omitempty"`
	MaxPaymentDays           int                   `xml:"max_payment_days,omitempty" json:"max_payment_days,omitempty"`
	Active                   string                `xml:"active,omitempty" json:"active,omitempty"`
	Note                     string                `xml:"note,omitempty" json:"note,omitempty"`
	IsGuest                  string                `xml:"is_guest,,omitempty" json:"is_guest,omitempty"`
	IDShop                   int                   `xml:"id_shop,omitempty" json:"id_shop,omitempty"`
	IDShopGroup              int                   `xml:"id_shop_group,omitempty" json:"id_shop_group,omitempty"`
	DateAdd                  string                `xml:"date_add,omitempty" json:"date_add,omitempty"`
	DateUpd                  string                `xml:"date_upd,omitempty" json:"date_upd,omitempty"`
	ResetPasswordToken       string                `xml:"reset_password_token,omitempty" json:"reset_password_token,omitempty"`
	ResetPasswordValidity    string                `xml:"reset_password_validity,omitempty" json:"reset_password_validity,omitempty"`
	Associations             *CustomerAssociations `xml:"associations,omitempty" json:"associations,omitempty"`
}

type CustomerAssociations struct {
	Groups CustomerAssociationsGroups `xml:"groups" json:"groups,omitempty"`
}

type CustomerAssociationsGroups struct {
	NodeType string  `xml:"nodeType,attr" json:"node_type,omitempty"`
	Api      string  `xml:"api,attr" json:"api,omitempty"`
	Group    []Group `xml:"group" json:"group,omitempty"`
}

type Group struct {
	Href string `xml:"href,attr" json:"href,omitempty"`
	ID   string `xml:"id" json:"id,omitempty"`
}

func (service *CustomersService) Create(customer *Customer) (*Customer, *http.Response, error) {
	createdCustomer := new(Customer)

	body := ResponseCustomer{
		Xlink:    "http://www.w3.org/1999/xlink",
		Customer: customer,
	}

	_url := makeResourceUrl("customers", nil)
	req, _ := service.client.NewRequest("POST", _url, body)

	customerResponse := new(ResponseCustomer)
	response, err := service.client.Do(req, customerResponse)

	if err != nil {
		return nil, response, err
	}

	if customerResponse != nil && customerResponse.Customer != nil {
		createdCustomer = customerResponse.Customer
	}

	return createdCustomer, response, nil
}

func (service *CustomersService) Get(customerID int, params *ServiceListParams) (*Customer, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("customers/%d", customerID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	customer := new(Customer)
	customerResponse := new(ResponseCustomer)
	response, err := service.client.Do(req, customerResponse)

	if err != nil {
		return nil, response, err
	}

	if customerResponse != nil {
		if customerResponse.Customer != nil {
			customer = customerResponse.Customer
		}

		// Use fisrt matching customer
		if customerResponse.CustomersData != nil && customerResponse.CustomersData.Customers != nil &&
			len(*customerResponse.CustomersData.Customers) > 0 {
			customer = &(*customerResponse.CustomersData.Customers)[0]
		}
	}

	return customer, response, nil
}

func (service *CustomersService) List(params *ServiceListParams) (*[]Customer, *http.Response, error) {
	customers := new([]Customer)

	_url := makeResourceUrl("customers", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	customersResponse := new(ResponseCustomer)
	response, err := service.client.Do(req, customersResponse)

	if err != nil {
		return nil, response, err
	}

	if customersResponse != nil && customersResponse.CustomersData != nil &&
		customersResponse.CustomersData.Customers != nil {
		customers = customersResponse.CustomersData.Customers
	}

	return customers, response, nil
}

func (service *CustomersService) GetCustomersByEmail(customerEmail string, params *ServiceListParams) (*Customer, *http.Response, error) {
	searchLimit := 1

	searchParams := ServiceListParams{
		Display: &ServiceListDisplay{
			"full",
		},
		Filter: &ServiceListFilter{
			Key:      "email",
			Values:   []string{customerEmail},
			Operator: ListFilterOperatorLiteral,
		},
		Limit: &ServiceListLimit{
			Limit: &searchLimit,
		},
		// Set defined sort params
		Sort: params.Sort,
	}

	// Override display params
	if params.Display != nil {
		searchParams.Display = params.Display
	}

	customer := new(Customer)
	customers, response, err := service.List(&searchParams)

	if err != nil {
		return nil, response, err
	}

	if customers != nil && len(*customers) > 0 {
		customer = &(*customers)[0]
	}

	return customer, response, err
}
