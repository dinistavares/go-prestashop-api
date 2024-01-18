package prestashop

import "errors"

type LanguageData struct {
	Language *[]Language `xml:"language,omitempty" json:"language,omitempty"`
}

type Language struct {
	ID    int    `xml:"id,attr" json:"id,omitempty"`
	Href  string `xml:"href,attr" json:"href,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
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

func GetDefaultLanguageValue(languageData *LanguageData) (string, error) {
	if languageData == nil || languageData.Language == nil || len(*languageData.Language) < 1 {
		return "", errors.New("language data missing")
	}

	return (*languageData.Language)[0].Value, nil
}

func GetLanguageValueByID(languageID int, languageData *LanguageData) (string, error) {
	var (
		languageValue string
		languageFound bool
	)

	if languageData == nil || languageData.Language == nil || len(*languageData.Language) == 0 {
		return languageValue, errors.New("language data missing")
	}

	for _, language := range *languageData.Language {
		if language.ID == languageID {
			languageValue = language.Value
			languageFound = true

			break
		}
	}

	if languageFound == false {
		return languageValue, errors.New("no langauge found")
	}

	return languageValue, nil
}
