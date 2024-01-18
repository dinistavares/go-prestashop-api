package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Currency service
type CurrencyService service

type Prestashop struct {
	XMLName xml.Name `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Text    string   `xml:",chardata,omitempty" json:",chardata,omitempty"`
	Xlink   string   `xml:"xlink,attr,omitempty" json:"xlink,attr,omitempty"`
}

type ResponseCurrency struct {
	XMLName       xml.Name       `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink         string         `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	Currency      *Currency      `xml:"currency,omitempty" json:"currency,omitempty"`
	CurrencysData *CurrencysData `xml:"currencies,omitempty" json:"currencies,omitempty"`
}

type CurrencysData struct {
	Currencys *[]Currency `xml:"currency,omitempty" json:"currency,omitempty"`
}

type Currency struct {
	ID             string        `xml:"id,omitempty" json:"id,omitempty"`
	NameData       string        `xml:"name,omitempty" json:"name,omitempty"`
	Names          *LanguageData `xml:"names,omitempty" json:"names,omitempty"`
	Symbol         *LanguageData `xml:"symbol,omitempty" json:"symbol,omitempty"`
	IsoCode        string        `xml:"iso_code,omitempty" json:"iso_code,omitempty"`
	NumericIsoCode string        `xml:"numeric_iso_code,omitempty" json:"numeric_iso_code,omitempty"`
	Precision      string        `xml:"precision,omitempty" json:"precision,omitempty"`
	ConversionRate string        `xml:"conversion_rate,omitempty" json:"conversion_rate,omitempty"`
	Deleted        string        `xml:"deleted,omitempty" json:"deleted,omitempty"`
	Active         string        `xml:"active,omitempty" json:"active,omitempty"`
	Unofficial     string        `xml:"unofficial,omitempty" json:"unofficial,omitempty"`
	Modified       string        `xml:"modified,omitempty" json:"modified,omitempty"`
	Pattern        *LanguageData `xml:"pattern,omitempty" json:"pattern,omitempty"`
}

func (service *CurrencyService) Create(currency *Currency) (*Currency, *http.Response, error) {
	createdCurrency := new(Currency)

	body := ResponseCurrency{
		Xlink:    "http://www.w3.org/1999/xlink",
		Currency: currency,
	}

	_url := "currencies"
	req, _ := service.client.NewRequest("POST", _url, body)

	currenciesResponse := new(ResponseCurrency)
	response, err := service.client.Do(req, currenciesResponse)

	if err != nil {
		return nil, response, err
	}

	if currenciesResponse != nil && currenciesResponse.Currency != nil {
		createdCurrency = currenciesResponse.Currency
	}

	return createdCurrency, response, nil
}

func (service *CurrencyService) Get(currencyID int, params *ServiceListParams) (*Currency, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("currencies/%d", currencyID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	currency := new(Currency)
	currenciesResponse := new(ResponseCurrency)
	response, err := service.client.Do(req, currenciesResponse)

	if err != nil {
		return nil, response, err
	}

	if currenciesResponse != nil {
		if currenciesResponse.Currency != nil {
			currency = currenciesResponse.Currency
		}

		// Use fisrt matching currency
		if currenciesResponse.CurrencysData != nil && currenciesResponse.CurrencysData.Currencys != nil &&
			len(*currenciesResponse.CurrencysData.Currencys) > 0 {
			currency = &(*currenciesResponse.CurrencysData.Currencys)[0]
		}
	}

	return currency, response, nil
}

func (service *CurrencyService) List(params *ServiceListParams) (*[]Currency, *http.Response, error) {
	currencies := new([]Currency)

	_url := makeResourceUrl("currencies", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	currenciesResponse := new(ResponseCurrency)
	response, err := service.client.Do(req, currenciesResponse)

	if err != nil {
		return nil, response, err
	}

	if currenciesResponse != nil && currenciesResponse.CurrencysData != nil &&
		currenciesResponse.CurrencysData.Currencys != nil {
		currencies = currenciesResponse.CurrencysData.Currencys
	}

	return currencies, response, nil
}

func (service *CurrencyService) ListCurrencysByCustomerID(customerID int, params *ServiceListParams) (*[]Currency, *http.Response, error) {
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

	currencies, response, err := service.List(&searchParams)

	if err != nil {
		return nil, response, err
	}

	return currencies, response, err
}
