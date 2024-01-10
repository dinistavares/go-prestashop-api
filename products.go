package prestashop

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Product service
type ProductService service

type ResponseProduct struct {
	XMLName xml.Name `xml:"prestashop" json:"prestashop,omitempty"`
	Xlink   string   `xml:"xlink,attr" json:"xlink,omitempty"`
	ProductData
}

type ProductData struct {
	Product  *Product   `xml:"product" json:"product,omitempty"`
	Products *[]Product `xml:"products" json:"products,omitempty"`
}

type Product struct {
	ID                      int                  `xml:"id" json:"id,omitempty"`
	IDManufacturer          int                  `xml:"id_manufacturer" json:"id_manufacturer,omitempty"`
	IDSupplier              int                  `xml:"id_supplier" json:"id_supplier,omitempty"`
	IDCategoryDefault       int                  `xml:"id_category_default" json:"id_category_default,omitempty"`
	New                     interface{}          `xml:"new" json:"new,omitempty"`
	CacheDefaultAttribute   int                  `xml:"cache_default_attribute" json:"cache_default_attribute,omitempty"`
	IDDefaultImage          int                  `xml:"id_default_image" json:"id_default_image,omitempty"`
	IDDefaultCombination    int                  `xml:"id_default_combination" json:"id_default_combination,omitempty"`
	IDTaxRulesGroup         int                  `xml:"id_tax_rules_group" json:"id_tax_rules_group,omitempty"`
	PositionInCategory      int                  `xml:"position_in_category" json:"position_in_category,omitempty"`
	ManufacturerName        string               `xml:"manufacturer_name" json:"manufacturer_name,omitempty"`
	Quantity                int                  `xml:"quantity" json:"quantity,omitempty"`
	Type                    string               `xml:"type" json:"type,omitempty"`
	IDShopDefault           int                  `xml:"id_shop_default" json:"id_shop_default,omitempty"`
	Reference               string               `xml:"reference" json:"reference,omitempty"`
	SupplierReference       string               `xml:"supplier_reference" json:"supplier_reference,omitempty"`
	Location                string               `xml:"location" json:"location,omitempty"`
	Width                   string               `xml:"width" json:"width,omitempty"`
	Height                  string               `xml:"height" json:"height,omitempty"`
	Depth                   string               `xml:"depth" json:"depth,omitempty"`
	Weight                  string               `xml:"weight" json:"weight,omitempty"`
	QuantityDiscount        string               `xml:"quantity_discount" json:"quantity_discount,omitempty"`
	Ean13                   string               `xml:"ean13" json:"ean13,omitempty"`
	Isbn                    string               `xml:"isbn" json:"isbn,omitempty"`
	Upc                     string               `xml:"upc" json:"upc,omitempty"`
	Mpn                     string               `xml:"mpn" json:"mpn,omitempty"`
	CacheIsPack             string               `xml:"cache_is_pack" json:"cache_is_pack,omitempty"`
	CacheHasAttachments     string               `xml:"cache_has_attachments" json:"cache_has_attachments,omitempty"`
	IsVirtual               string               `xml:"is_virtual" json:"is_virtual,omitempty"`
	State                   int                  `xml:"state" json:"state,omitempty"`
	AdditionalDeliveryTimes int                  `xml:"additional_delivery_times" json:"additional_delivery_times,omitempty"`
	DeliveryInStock         string               `xml:"delivery_in_stock" json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        string               `xml:"delivery_out_stock" json:"delivery_out_stock,omitempty"`
	ProductType             string               `xml:"product_type" json:"product_type,omitempty"`
	OnSale                  string               `xml:"on_sale" json:"on_sale,omitempty"`
	OnlineOnly              string               `xml:"online_only" json:"online_only,omitempty"`
	Ecotax                  string               `xml:"ecotax" json:"ecotax,omitempty"`
	MinimalQuantity         int                  `xml:"minimal_quantity" json:"minimal_quantity,omitempty"`
	LowStockThreshold       int                  `xml:"low_stock_threshold" json:"low_stock_threshold,omitempty"`
	LowStockAlert           string               `xml:"low_stock_alert" json:"low_stock_alert,omitempty"`
	Price                   string               `xml:"price" json:"price,omitempty"`
	WholesalePrice          string               `xml:"wholesale_price" json:"wholesale_price,omitempty"`
	Unity                   string               `xml:"unity" json:"unity,omitempty"`
	UnitPrice               string               `xml:"unit_price" json:"unit_price,omitempty"`
	UnitPriceRatio          string               `xml:"unit_price_ratio" json:"unit_price_ratio,omitempty"`
	AdditionalShippingCost  string               `xml:"additional_shipping_cost" json:"additional_shipping_cost,omitempty"`
	Customizable            int                  `xml:"customizable" json:"customizable,omitempty"`
	TextFields              int                  `xml:"text_fields" json:"text_fields,omitempty"`
	UploadableFiles         int                  `xml:"uploadable_files" json:"uploadable_files,omitempty"`
	Active                  string               `xml:"active" json:"active,omitempty"`
	RedirectType            string               `xml:"redirect_type" json:"redirect_type,omitempty"`
	IDTypeRedirected        int                  `xml:"id_type_redirected" json:"id_type_redirected,omitempty"`
	AvailableForOrder       string               `xml:"available_for_order" json:"available_for_order,omitempty"`
	AvailableDate           string               `xml:"available_date" json:"available_date,omitempty"`
	ShowCondition           string               `xml:"show_condition" json:"show_condition,omitempty"`
	Condition               string               `xml:"condition" json:"condition,omitempty"`
	ShowPrice               string               `xml:"show_price" json:"show_price,omitempty"`
	Indexed                 string               `xml:"indexed" json:"indexed,omitempty"`
	Visibility              string               `xml:"visibility" json:"visibility,omitempty"`
	AdvancedStockManagement string               `xml:"advanced_stock_management" json:"advanced_stock_management,omitempty"`
	DateAdd                 string               `xml:"date_add" json:"date_add,omitempty"`
	DateUpd                 string               `xml:"date_upd" json:"date_upd,omitempty"`
	PackStockType           int                  `xml:"pack_stock_type" json:"pack_stock_type,omitempty"`
	MetaDescription         string               `xml:"meta_description" json:"meta_description,omitempty"`
	MetaKeywords            string               `xml:"meta_keywords" json:"meta_keywords,omitempty"`
	MetaTitle               string               `xml:"meta_title" json:"meta_title,omitempty"`
	LinkRewrite             string               `xml:"link_rewrite" json:"link_rewrite,omitempty"`
	Name                    string               `xml:"name" json:"name,omitempty"`
	Description             string               `xml:"description" json:"description,omitempty"`
	DescriptionShort        string               `xml:"description_short" json:"description_short,omitempty"`
	AvailableNow            string               `xml:"available_now" json:"available_now,omitempty"`
	AvailableLater          string               `xml:"available_later" json:"available_later,omitempty"`
	Associations            *ProductAssociations `xml:"associations" json:"associations,omitempty"`
}

type Categories struct {
	ID int `xml:"id" json:"id,omitempty"`
}

type Images struct {
	ID int `xml:"id" json:"id,omitempty"`
}

type Combinations struct {
	ID int `xml:"id" json:"id,omitempty"`
}

type ProductOptionValues struct {
	ID int `xml:"id" json:"id,omitempty"`
}

type ProductFeatures struct {
	ID             int `xml:"id" json:"id,omitempty"`
	IDFeatureValue int `xml:"id_feature_value" json:"id_feature_value,omitempty"`
}

type StockAvailables struct {
	ID                 int `xml:"id" json:"id,omitempty"`
	IDProductAttribute int `xml:"id_product_attribute" json:"id_product_attribute,omitempty"`
}

type ProductAssociations struct {
	Categories          *[]Categories          `xml:"categories" json:"categories,omitempty"`
	Images              *[]Images              `xml:"images" json:"images,omitempty"`
	Combinations        *[]Combinations        `xml:"combinations" json:"combinations,omitempty"`
	ProductOptionValues *[]ProductOptionValues `xml:"product_option_values" json:"product_option_values,omitempty"`
	ProductFeatures     *[]ProductFeatures     `xml:"product_features" json:"product_features,omitempty"`
	StockAvailables     *[]StockAvailables     `xml:"stock_availables" json:"stock_availables,omitempty"`
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
