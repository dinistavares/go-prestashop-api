package prestashop

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Customer service
type CustomersService service

type ResponseCustomer struct {
	XMLName xml.Name `xml:"prestashop" json:"prestashop,omitempty"`
	Xlink   string   `xml:"xlink,attr" json:"xlink,omitempty"`
	CustomerData
}

type CustomerData struct {
	Customer  *Customer   `xml:"customer" json:"customer,omitempty"`
	Customers *[]Customer `xml:"customers" json:"customers,omitempty"`
}

type Customer struct {
	ID                       int                   `xml:"id" json:"id,omitempty"`
	IDDefaultGroup           int                   `xml:"id_default_group" json:"id_default_group,omitempty"`
	IDLang                   int                   `xml:"id_lang" json:"id_lang,omitempty"`
	NewsletterDateAdd        string                `xml:"newsletter_date_add" json:"newsletter_date_add,omitempty"`
	IPRegistrationNewsletter string                `xml:"ip_registration_newsletter" json:"ip_registration_newsletter,omitempty"`
	LastPasswdGen            string                `xml:"last_passwd_gen" json:"last_passwd_gen,omitempty"`
	SecureKey                string                `xml:"secure_key" json:"secure_key,omitempty"`
	Deleted                  string                `xml:"deleted" json:"deleted,omitempty"`
	Passwd                   string                `xml:"passwd" json:"passwd,omitempty"`
	Lastname                 string                `xml:"lastname" json:"lastname,omitempty"`
	Firstname                string                `xml:"firstname" json:"firstname,omitempty"`
	Email                    string                `xml:"email" json:"email,omitempty"`
	IDGender                 int                   `xml:"id_gender" json:"id_gender,omitempty"`
	Birthday                 string                `xml:"birthday" json:"birthday,omitempty"`
	Newsletter               string                `xml:"newsletter" json:"newsletter,omitempty"`
	Optin                    string                `xml:"optin" json:"optin,omitempty"`
	Website                  string                `xml:"website" json:"website,omitempty"`
	Company                  string                `xml:"company" json:"company,omitempty"`
	Siret                    string                `xml:"siret" json:"siret,omitempty"`
	Ape                      string                `xml:"ape" json:"ape,omitempty"`
	OutstandingAllowAmount   string                `xml:"outstanding_allow_amount" json:"outstanding_allow_amount,omitempty"`
	ShowPublicPrices         string                `xml:"show_public_prices" json:"show_public_prices,omitempty"`
	IDRisk                   int                   `xml:"id_risk" json:"id_risk,omitempty"`
	MaxPaymentDays           int                   `xml:"max_payment_days" json:"max_payment_days,omitempty"`
	Active                   string                `xml:"active" json:"active,omitempty"`
	Note                     string                `xml:"note" json:"note,omitempty"`
	IsGuest                  string                `xml:"is_guest," json:"is_guest,omitempty"`
	IDShop                   int                   `xml:"id_shop" json:"id_shop,omitempty"`
	IDShopGroup              int                   `xml:"id_shop_group" json:"id_shop_group,omitempty"`
	DateAdd                  string                `xml:"date_add" json:"date_add,omitempty"`
	DateUpd                  string                `xml:"date_upd" json:"date_upd,omitempty"`
	ResetPasswordToken       string                `xml:"reset_password_token" json:"reset_password_token,omitempty"`
	ResetPasswordValidity    string                `xml:"reset_password_validity" json:"reset_password_validity,omitempty"`
	Associations             *CustomerAssociations `xml:"associations" json:"associations,omitempty"`
}

type CustomerAssociations struct {
	Groups []Group `xml:"groups" json:"groups,omitempty"`
}

type Group struct {
	ID int `xml:"groups" json:"id,omitempty"`
}

func (service *CustomersService) Get(customerID int, params *ServiceListParams) (*Customer, *http.Response, error) {
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
		if customerResponse.Customers != nil && len(*customerResponse.Customers) > 0 {
			customer = &(*customerResponse.Customers)[0]
		}
	}

	return customer, response, nil
}

func (service *CustomersService) List(params *ServiceListParams) (*[]Customer, *http.Response, error) {
	_url := makeResourceUrl("customers", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	customers := new(ResponseCustomer)
	response, err := service.client.Do(req, customers)

	if err != nil {
		// API returns 200 but the response is not a JSON object, return no customers found
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value of type prestashop.ResponseCustomer") {
			return nil, response, errors.New("no customers found")
		}

		return nil, response, err
	}

	return customers.Customers, response, nil
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
