package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Carrier service
type CarrierService service

type ResponseCarrier struct {
	XMLName      xml.Name     `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink        string       `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	Carrier      *Carrier     `xml:"carrier,omitempty" json:"carrier,omitempty"`
	CarriersData *CarrierData `xml:"carriers,omitempty" json:"carriers,omitempty"`
}

type CarrierData struct {
	Carriers *[]Carrier `xml:"carrier,omitempty" json:"carrier,omitempty"`
}

type TaxRulesGroup struct {
	ID            int    `xml:",chardata" json:"id,omitempty"`
	Href          string `xml:"href,attr,omitempty" json:"href,attr,omitempty"`
	NotFilterable string `xml:"notFilterable,attr,omitempty" json:"notFilterable,attr,omitempty"`
}

type Carrier struct {
	ID                 string         `xml:"id,omitempty" json:"id,omitempty"`
	Deleted            string         `xml:"deleted,omitempty" json:"deleted,omitempty"`
	IsModule           string         `xml:"is_module,omitempty" json:"is_module,omitempty"`
	TaxRulesGroup      *TaxRulesGroup `xml:"id_tax_rules_group,omitempty" json:"id_tax_rules_group,omitempty"`
	IDReference        string         `xml:"id_reference,omitempty" json:"id_reference,omitempty"`
	Name               string         `xml:"name,omitempty" json:"name,omitempty"`
	Active             string         `xml:"active,omitempty" json:"active,omitempty"`
	IsFree             string         `xml:"is_free,omitempty" json:"is_free,omitempty"`
	URL                string         `xml:"url,omitempty" json:"url,omitempty"`
	ShippingHandling   string         `xml:"shipping_handling,omitempty" json:"shipping_handling,omitempty"`
	ShippingExternal   string         `xml:"shipping_external,omitempty" json:"shipping_external,omitempty"`
	RangeBehavior      string         `xml:"range_behavior,omitempty" json:"range_behavior,omitempty"`
	ShippingMethod     string         `xml:"shipping_method,omitempty" json:"shipping_method,omitempty"`
	MaxWidth           float64        `xml:"max_width,omitempty" json:"max_width,omitempty"`
	MaxHeight          float64        `xml:"max_height,omitempty" json:"max_height,omitempty"`
	MaxDepth           float64        `xml:"max_depth,omitempty" json:"max_depth,omitempty"`
	MaxWeight          float64        `xml:"max_weight,omitempty" json:"max_weight,omitempty"`
	Grade              string         `xml:"grade,omitempty" json:"grade,omitempty"`
	ExternalModuleName string         `xml:"external_module_name,omitempty" json:"external_module_name,omitempty"`
	NeedRange          string         `xml:"need_range,omitempty" json:"need_range,omitempty"`
	Position           string         `xml:"position,omitempty" json:"position,omitempty"`
	Delay              *LanguageData  `xml:"delay,omitempty" json:"delay,omitempty"`
}

func (service *CarrierService) Create(carrier *Carrier) (*Carrier, *http.Response, error) {
	createdCarrier := new(Carrier)

	body := ResponseCarrier{
		Xlink:   "http://www.w3.org/1999/xlink",
		Carrier: carrier,
	}

	_url := "carriers"
	req, _ := service.client.NewRequest("POST", _url, body)

	carriersResponse := new(ResponseCarrier)
	response, err := service.client.Do(req, carriersResponse)

	if err != nil {
		return nil, response, err
	}

	if carriersResponse != nil && carriersResponse.Carrier != nil {
		createdCarrier = carriersResponse.Carrier
	}

	return createdCarrier, response, nil
}

func (service *CarrierService) Get(carrierID int, params *ServiceListParams) (*Carrier, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("carriers/%d", carrierID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	carrier := new(Carrier)
	carriersResponse := new(ResponseCarrier)
	response, err := service.client.Do(req, carriersResponse)

	if err != nil {
		return nil, response, err
	}

	if carriersResponse != nil {
		if carriersResponse.Carrier != nil {
			carrier = carriersResponse.Carrier
		}

		// Use fisrt matching carrier
		if carriersResponse.CarriersData != nil && carriersResponse.CarriersData.Carriers != nil &&
			len(*carriersResponse.CarriersData.Carriers) > 0 {
			carrier = &(*carriersResponse.CarriersData.Carriers)[0]
		}
	}

	return carrier, response, nil
}

func (service *CarrierService) List(params *ServiceListParams) (*[]Carrier, *http.Response, error) {
	carriers := new([]Carrier)
	_url := makeResourceUrl("carriers", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	carriersResponse := new(ResponseCarrier)
	response, err := service.client.Do(req, carriersResponse)

	if err != nil {
		return nil, response, err
	}

	if carriersResponse != nil && carriersResponse.CarriersData != nil &&
		carriersResponse.CarriersData.Carriers != nil {
		carriers = carriersResponse.CarriersData.Carriers
	}

	return carriers, response, nil
}
