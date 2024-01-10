package prestashop

func makeResourceUrl(route string, listParams *ServiceListParams) string {
	_url := route

	listParamsString := parseUrlListParameters(listParams)

	if listParamsString != "" {
		_url += "?" + listParamsString
	}

	return _url
}
