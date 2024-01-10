package prestashop

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type ListFilterOperator uint8

type ListSortOrder uint8

const (
	// ListFilterOperatorOr specifies an 'or' filter type eg. 'filter[id]=[1|5]' = (returns values with id's 1 OR 5)
	ListFilterOperatorOr       ListFilterOperator = 0

	// ListFilterOperatorInterval specifies an 'interval' filter type eg. 'filter[id]=[1,5]' = (returns values id's 1 through 5)
	ListFilterOperatorInterval ListFilterOperator = 1

	// ListFilterOperatorLiteral specifies a 'literal' filter type eg. 'filter[field]=[value]'
	ListFilterOperatorLiteral  ListFilterOperator = 2

	// ListFilterOperatorBegin specifies a 'begin' filter type eg. 'filter[field]=[val]%'
	ListFilterOperatorBegin    ListFilterOperator = 3

	// ListFilterOperatorEnd specifies an 'end' filter type eg. 'filter[field]=%[ue]'
	ListFilterOperatorEnd      ListFilterOperator = 4

	// ListFilterOperatorContains specifies a 'contains' filter type eg. 'filter[field]=%[alu]%'
	ListFilterOperatorContains ListFilterOperator = 5

	// ListSortOrderAscending specifies a 'ascending ' sort type eg. 'sort=[lastname_ASC]'
	ListSortOrderAscending     ListSortOrder = 0 
	
	// ListSortOrderAscending specifies a 'descending ' sort type eg. 'sort=[lastname_DESC]'
	ListSortOrderDescending    ListSortOrder = 1 
)

type ServiceListParams struct {
	Display *ServiceListDisplay
	Filter  *ServiceListFilter
	Limit   *ServiceListLimit
	Sort    *[]ServiceListSort
}

type ServiceListDisplay []string

type ServiceListFilter struct {
	Key      string
	Values   []string
	Operator ListFilterOperator
}

type ServiceListSort struct {
	Key   string
	Order ListSortOrder
}

type ServiceListLimit struct {
	StartAt *int
	Limit   *int
}

func parseServiceListDisplay(display ServiceListDisplay) string {
	displayValues := strings.Join(display, ",")

	if displayValues == "" {
		return displayValues
	}

	if displayValues == "full" {
		return "display=" + displayValues
	}

	return "display=[" + url.QueryEscape(displayValues) + "]"
}

func parseServiceListFilter(filter ServiceListFilter) string {
	var (
		filterString string
		filterValue  string
	)

	if filter.Key == "" || len(filter.Values) == 0 {
		return filterString
	}

	switch filter.Operator {
	case ListFilterOperatorLiteral:
		filterValue = fmt.Sprintf("[%v]", filter.Values[0])

		break
	case ListFilterOperatorBegin:
		filterValue = fmt.Sprintf("[%v]%%", filter.Values[0])

		break
	case ListFilterOperatorEnd:
		filterValue = fmt.Sprintf("%%[%v]", filter.Values[0])

		break
	case ListFilterOperatorContains:
		filterValue = fmt.Sprintf("%%[%v]%%", filter.Values[0])

		break
	case ListFilterOperatorOr:
		filterValue = fmt.Sprintf("[%v]", url.QueryEscape(strings.Join(filter.Values, "|")))

		break
	case ListFilterOperatorInterval:
		filterValue = fmt.Sprintf("[%v]", url.QueryEscape(strings.Join(filter.Values, ",")))

		break
	default:
		return filterString
	}

	// Form filter string
	filterString = "filter[" + filter.Key + "]=" + filterValue

	return filterString
}

func parseServiceListSort(sort []ServiceListSort) string {
	var sortValues []string

	for _, sort := range sort {
		if sort.Key == "" {
			continue
		}

		switch sort.Order {
		case ListSortOrderAscending:
			sortValues = append(sortValues, sort.Key+"_ASC")

			break
		case ListSortOrderDescending:
			sortValues = append(sortValues, sort.Key+"_DESC")

			break
		}
	}

	return "sort=[" + url.QueryEscape(strings.Join(sortValues, ",")) + "]"
}

func parseServiceListLimit(limit ServiceListLimit) string {
	var limitString string

	if limit.Limit == nil {
		return limitString
	}

	if limit.StartAt != nil {
		limitString = strconv.Itoa(*limit.StartAt) + ","
	}

	limitString += strconv.Itoa(*limit.Limit)

	return "limit=" + url.QueryEscape(limitString)
}

func parseUrlListParameters(params *ServiceListParams) string {
	var listParamString []string

	if params == nil {
		return ""
	}

	if params.Display != nil {
		listParamString = append(listParamString, parseServiceListDisplay(*params.Display))
	}

	if params.Filter != nil {
		listParamString = append(listParamString, parseServiceListFilter(*params.Filter))
	}

	if params.Sort != nil {
		listParamString = append(listParamString, parseServiceListSort(*params.Sort))
	}

	if params.Limit != nil {
		listParamString = append(listParamString, parseServiceListLimit(*params.Limit))
	}

	return strings.Join(listParamString, "&")
}
