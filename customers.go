package prestashop

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Customer service
type CustomersService service

type ResponseCustomer struct {
	Customer  *Customer   `json:"customer,omitempty"`
	Customers *[]Customer `json:"customers,omitempty"`
}

type Customer struct {
	ID                       int                   `json:"id,omitempty"`
	IDDefaultGroup           int                   `json:"id_default_group,omitempty"`
	IDLang                   int                   `json:"id_lang,omitempty"`
	NewsletterDateAdd        string                `json:"newsletter_date_add,omitempty"`
	IPRegistrationNewsletter string                `json:"ip_registration_newsletter,omitempty"`
	LastPasswdGen            string                `json:"last_passwd_gen,omitempty"`
	SecureKey                string                `json:"secure_key,omitempty"`
	Deleted                  string                `json:"deleted,omitempty"`
	Passwd                   string                `json:"passwd,omitempty"`
	Lastname                 string                `json:"lastname,omitempty"`
	Firstname                string                `json:"firstname,omitempty"`
	Email                    string                `json:"email,omitempty"`
	IDGender                 int                   `json:"id_gender,omitempty"`
	Birthday                 string                `json:"birthday,omitempty"`
	Newsletter               string                `json:"newsletter,omitempty"`
	Optin                    string                `json:"optin,omitempty"`
	Website                  string                `json:"website,omitempty"`
	Company                  string                `json:"company,omitempty"`
	Siret                    string                `json:"siret,omitempty"`
	Ape                      string                `json:"ape,omitempty"`
	OutstandingAllowAmount   string                `json:"outstanding_allow_amount,omitempty"`
	ShowPublicPrices         string                `json:"show_public_prices,omitempty"`
	IDRisk                   int                   `json:"id_risk,omitempty"`
	MaxPaymentDays           int                   `json:"max_payment_days,omitempty"`
	Active                   string                `json:"active,omitempty"`
	Note                     string                `json:"note,omitempty"`
	IsGuest                  string                `json:"is_guest,omitempty"`
	IDShop                   int                   `json:"id_shop,omitempty"`
	IDShopGroup              int                   `json:"id_shop_group,omitempty"`
	DateAdd                  string                `json:"date_add,omitempty"`
	DateUpd                  string                `json:"date_upd,omitempty"`
	ResetPasswordToken       string                `json:"reset_password_token,omitempty"`
	ResetPasswordValidity    string                `json:"reset_password_validity,omitempty"`
	Associations             *CustomerAssociations `json:"associations,omitempty"`
}

type CustomerAssociations struct {
	Groups []Group `json:"groups,omitempty"`
}

type Group struct {
	ID int `json:"id,omitempty"`
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
