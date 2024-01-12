package prestashop

type LanguageData struct {
	ID       string `xml:"id" json:"id,omitempty"`
	Href     string `xml:"href" json:"href,omitempty"`
	Language string `xml:"language,omitempty" json:"language,omitempty"`
}

func makeResourceUrl(route string, listParams *ServiceListParams) string {
	_url := route

	listParamsString := parseUrlListParameters(listParams)

	if listParamsString != "" {
		_url += "?" + listParamsString
	}

	return _url
}

func setDefaultResourceByIDDisplayParams(params *ServiceListParams) *ServiceListParams {
	if params == nil {
		params = &ServiceListParams{}
	}

	if params.Display == nil {
		params.Display = &ServiceListDisplay{"full"}
	}

	return params
}
