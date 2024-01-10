package prestashop

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Product service
type ProductService service

type ResponseProduct struct {
	Product  *Product   `json:"product,omitempty"`
	Products *[]Product `json:"products,omitempty"`
}

type Product struct {
	ID                      int          `json:"id,omitempty"`
	IDManufacturer          int          `json:"id_manufacturer,omitempty"`
	IDSupplier              int          `json:"id_supplier,omitempty"`
	IDCategoryDefault       int          `json:"id_category_default,omitempty"`
	New                     interface{}  `json:"new,omitempty"`
	CacheDefaultAttribute   int          `json:"cache_default_attribute,omitempty"`
	IDDefaultImage          int          `json:"id_default_image,omitempty"`
	IDDefaultCombination    int          `json:"id_default_combination,omitempty"`
	IDTaxRulesGroup         int          `json:"id_tax_rules_group,omitempty"`
	PositionInCategory      int          `json:"position_in_category,omitempty"`
	ManufacturerName        string       `json:"manufacturer_name,omitempty"`
	Quantity                int          `json:"quantity,omitempty"`
	Type                    string       `json:"type,omitempty"`
	IDShopDefault           int          `json:"id_shop_default,omitempty"`
	Reference               string       `json:"reference,omitempty"`
	SupplierReference       string       `json:"supplier_reference,omitempty"`
	Location                string       `json:"location,omitempty"`
	Width                   string       `json:"width,omitempty"`
	Height                  string       `json:"height,omitempty"`
	Depth                   string       `json:"depth,omitempty"`
	Weight                  string       `json:"weight,omitempty"`
	QuantityDiscount        string       `json:"quantity_discount,omitempty"`
	Ean13                   string       `json:"ean13,omitempty"`
	Isbn                    string       `json:"isbn,omitempty"`
	Upc                     string       `json:"upc,omitempty"`
	Mpn                     string       `json:"mpn,omitempty"`
	CacheIsPack             string       `json:"cache_is_pack,omitempty"`
	CacheHasAttachments     string       `json:"cache_has_attachments,omitempty"`
	IsVirtual               string       `json:"is_virtual,omitempty"`
	State                   int          `json:"state,omitempty"`
	AdditionalDeliveryTimes int          `json:"additional_delivery_times,omitempty"`
	DeliveryInStock         string       `json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        string       `json:"delivery_out_stock,omitempty"`
	ProductType             string       `json:"product_type,omitempty"`
	OnSale                  string       `json:"on_sale,omitempty"`
	OnlineOnly              string       `json:"online_only,omitempty"`
	Ecotax                  string       `json:"ecotax,omitempty"`
	MinimalQuantity         int          `json:"minimal_quantity,omitempty"`
	LowStockThreshold       int          `json:"low_stock_threshold,omitempty"`
	LowStockAlert           string       `json:"low_stock_alert,omitempty"`
	Price                   string       `json:"price,omitempty"`
	WholesalePrice          string       `json:"wholesale_price,omitempty"`
	Unity                   string       `json:"unity,omitempty"`
	UnitPrice               string       `json:"unit_price,omitempty"`
	UnitPriceRatio          string       `json:"unit_price_ratio,omitempty"`
	AdditionalShippingCost  string       `json:"additional_shipping_cost,omitempty"`
	Customizable            int          `json:"customizable,omitempty"`
	TextFields              int          `json:"text_fields,omitempty"`
	UploadableFiles         int          `json:"uploadable_files,omitempty"`
	Active                  string       `json:"active,omitempty"`
	RedirectType            string       `json:"redirect_type,omitempty"`
	IDTypeRedirected        int          `json:"id_type_redirected,omitempty"`
	AvailableForOrder       string       `json:"available_for_order,omitempty"`
	AvailableDate           string       `json:"available_date,omitempty"`
	ShowCondition           string       `json:"show_condition,omitempty"`
	Condition               string       `json:"condition,omitempty"`
	ShowPrice               string       `json:"show_price,omitempty"`
	Indexed                 string       `json:"indexed,omitempty"`
	Visibility              string       `json:"visibility,omitempty"`
	AdvancedStockManagement string       `json:"advanced_stock_management,omitempty"`
	DateAdd                 string       `json:"date_add,omitempty"`
	DateUpd                 string       `json:"date_upd,omitempty"`
	PackStockType           int          `json:"pack_stock_type,omitempty"`
	MetaDescription         string       `json:"meta_description,omitempty"`
	MetaKeywords            string       `json:"meta_keywords,omitempty"`
	MetaTitle               string       `json:"meta_title,omitempty"`
	LinkRewrite             string       `json:"link_rewrite,omitempty"`
	Name                    string       `json:"name,omitempty"`
	Description             string       `json:"description,omitempty"`
	DescriptionShort        string       `json:"description_short,omitempty"`
	AvailableNow            string       `json:"available_now,omitempty"`
	AvailableLater          string       `json:"available_later,omitempty"`
	Associations            *ProductAssociations `json:"associations,omitempty"`
}

type Categories struct {
	ID int `json:"id,omitempty"`
}

type Images struct {
	ID int `json:"id,omitempty"`
}

type Combinations struct {
	ID int `json:"id,omitempty"`
}

type ProductOptionValues struct {
	ID int `json:"id,omitempty"`
}

type ProductFeatures struct {
	ID             int `json:"id,omitempty"`
	IDFeatureValue int `json:"id_feature_value,omitempty"`
}

type StockAvailables struct {
	ID                 int `json:"id,omitempty"`
	IDProductAttribute int `json:"id_product_attribute,omitempty"`
}

type ProductAssociations struct {
	Categories          *[]Categories          `json:"categories,omitempty"`
	Images              *[]Images              `json:"images,omitempty"`
	Combinations        *[]Combinations        `json:"combinations,omitempty"`
	ProductOptionValues *[]ProductOptionValues `json:"product_option_values,omitempty"`
	ProductFeatures     *[]ProductFeatures     `json:"product_features,omitempty"`
	StockAvailables     *[]StockAvailables     `json:"stock_availables,omitempty"`
}

func (service *ProductService) Get(productID int, params *ServiceListParams) (*Product, *http.Response, error) {
	resourceRoute := fmt.Sprintf("products/%d", productID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	product := new(Product)
	productsResponse := new(ResponseProduct)
	response, err := service.client.Do(req, productsResponse)

	if err != nil {
		return nil, response, err
	}

	if productsResponse != nil {
		if productsResponse.Product != nil {
			product = productsResponse.Product
		}

		// Use fisrt matching product
		if productsResponse.Products != nil && len(*productsResponse.Products) > 0 {
			product = &(*productsResponse.Products)[0]
		}
	}

	return product, response, nil
}

func (service *ProductService) List(params *ServiceListParams) (*[]Product, *http.Response, error) {
	_url := makeResourceUrl("products", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	products := new(ResponseProduct)
	response, err := service.client.Do(req, products)

	if err != nil {
		// API returns 200 but the response is not a JSON object, return no customers found
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value of type prestashop.ResponseProduct") {
			return nil, response, errors.New("no products found")
		}

		return nil, response, err
	}

	return products.Products, response, nil
}

// func (service *ProductService) GetProductssByCustomerID(customerID int, params *ServiceListParams) (*[]Product, *http.Response, error) {
// 	searchParams := ServiceListParams{
// 		Display: &ServiceListDisplay{
// 			"full",
// 		},
// 		Filter: &ServiceListFilter{
// 			Key:      "id_customer",
// 			Values:   []string{fmt.Sprintf("%d", customerID)},
// 			Operator: ListFilterOperatorLiteral,
// 		},
// 		// Set defined sort and limit params
// 		Limit: params.Limit,
// 		Sort:  params.Sort,
// 	}

// 	// Override display params
// 	if params.Display != nil {
// 		searchParams.Display = params.Display
// 	}

// 	products, response, err := service.List(&searchParams)

// 	if err != nil {
// 		return nil, response, err
// 	}

// 	return products, response, err
// }
