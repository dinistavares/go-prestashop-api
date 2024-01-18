package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Product service
type ProductService service

type ResponseProduct struct {
	XMLName      xml.Name     `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink        string       `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	Product      *Product     `xml:"product,omitempty" json:"product,omitempty"`
	ProductsData *ProductData `xml:"products,omitempty" json:"products,omitempty"`
}

type ProductData struct {
	Products *[]Product `xml:"product,omitempty" json:"product,omitempty"`
}

type DefaultImage struct {
	ID   int    `xml:",chardata" json:"id,omitempty"`
	Href string `xml:"href,attr" json:"image_url,omitempty"`
}

type Product struct {
	ID                      int                  `xml:"id,omitempty" json:"id,omitempty"`
	IDManufacturer          int                  `xml:"id_manufacturer,omitempty" json:"id_manufacturer,omitempty"`
	IDSupplier              int                  `xml:"id_supplier,omitempty" json:"id_supplier,omitempty"`
	IDCategoryDefault       int                  `xml:"id_category_default,omitempty" json:"id_category_default,omitempty"`
	New                     interface{}          `xml:"new,omitempty" json:"new,omitempty"`
	CacheDefaultAttribute   int                  `xml:"cache_default_attribute,omitempty" json:"cache_default_attribute,omitempty"`
	DefaultImage            *DefaultImage        `xml:"id_default_image,omitempty" json:"id_default_image,omitempty"`
	IDDefaultCombination    int                  `xml:"id_default_combination,omitempty" json:"id_default_combination,omitempty"`
	IDTaxRulesGroup         int                  `xml:"id_tax_rules_group,omitempty" json:"id_tax_rules_group,omitempty"`
	PositionInCategory      int                  `xml:"position_in_category,omitempty" json:"position_in_category,omitempty"`
	ManufacturerName        string               `xml:"manufacturer_name,omitempty" json:"manufacturer_name,omitempty"`
	Quantity                int                  `xml:"quantity,omitempty" json:"quantity,omitempty"`
	Type                    string               `xml:"type,omitempty" json:"type,omitempty"`
	IDShopDefault           int                  `xml:"id_shop_default,omitempty" json:"id_shop_default,omitempty"`
	Reference               string               `xml:"reference,omitempty" json:"reference,omitempty"`
	SupplierReference       string               `xml:"supplier_reference,omitempty" json:"supplier_reference,omitempty"`
	Location                string               `xml:"location,omitempty" json:"location,omitempty"`
	Width                   float64              `xml:"width,omitempty" json:"width,omitempty"`
	Height                  float64              `xml:"height,omitempty" json:"height,omitempty"`
	Depth                   float64              `xml:"depth,omitempty" json:"depth,omitempty"`
	Weight                  float64              `xml:"weight,omitempty" json:"weight,omitempty"`
	QuantityDiscount        string               `xml:"quantity_discount,omitempty" json:"quantity_discount,omitempty"`
	Ean13                   string               `xml:"ean13,omitempty" json:"ean13,omitempty"`
	Isbn                    string               `xml:"isbn,omitempty" json:"isbn,omitempty"`
	Upc                     string               `xml:"upc,omitempty" json:"upc,omitempty"`
	Mpn                     string               `xml:"mpn,omitempty" json:"mpn,omitempty"`
	CacheIsPack             string               `xml:"cache_is_pack,omitempty" json:"cache_is_pack,omitempty"`
	CacheHasAttachments     string               `xml:"cache_has_attachments,omitempty" json:"cache_has_attachments,omitempty"`
	IsVirtual               string               `xml:"is_virtual,omitempty" json:"is_virtual,omitempty"`
	State                   int                  `xml:"state,omitempty" json:"state,omitempty"`
	AdditionalDeliveryTimes int                  `xml:"additional_delivery_times,omitempty" json:"additional_delivery_times,omitempty"`
	DeliveryInStock         *LanguageData        `xml:"delivery_in_stock,omitempty" json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        *LanguageData        `xml:"delivery_out_stock,omitempty" json:"delivery_out_stock,omitempty"`
	ProductType             string               `xml:"product_type,omitempty" json:"product_type,omitempty"`
	OnSale                  string               `xml:"on_sale,omitempty" json:"on_sale,omitempty"`
	OnlineOnly              string               `xml:"online_only,omitempty" json:"online_only,omitempty"`
	Ecotax                  float64              `xml:"ecotax,omitempty" json:"ecotax,omitempty"`
	MinimalQuantity         int                  `xml:"minimal_quantity,omitempty" json:"minimal_quantity,omitempty"`
	LowStockThreshold       int                  `xml:"low_stock_threshold,omitempty" json:"low_stock_threshold,omitempty"`
	LowStockAlert           string               `xml:"low_stock_alert,omitempty" json:"low_stock_alert,omitempty"`
	Price                   float64              `xml:"price,omitempty" json:"price,omitempty"`
	WholesalePrice          float64              `xml:"wholesale_price,omitempty" json:"wholesale_price,omitempty"`
	Unity                   string               `xml:"unity,omitempty" json:"unity,omitempty"`
	UnitPrice               float64              `xml:"unit_price,omitempty" json:"unit_price,omitempty"`
	UnitPriceRatio          float64              `xml:"unit_price_ratio,omitempty" json:"unit_price_ratio,omitempty"`
	AdditionalShippingCost  float64              `xml:"additional_shipping_cost,omitempty" json:"additional_shipping_cost,omitempty"`
	Customizable            int                  `xml:"customizable,omitempty" json:"customizable,omitempty"`
	TextFields              int                  `xml:"text_fields,omitempty" json:"text_fields,omitempty"`
	UploadableFiles         int                  `xml:"uploadable_files,omitempty" json:"uploadable_files,omitempty"`
	Active                  string               `xml:"active,omitempty" json:"active,omitempty"`
	RedirectType            string               `xml:"redirect_type,omitempty" json:"redirect_type,omitempty"`
	IDTypeRedirected        int                  `xml:"id_type_redirected,omitempty" json:"id_type_redirected,omitempty"`
	AvailableForOrder       string               `xml:"available_for_order,omitempty" json:"available_for_order,omitempty"`
	AvailableDate           string               `xml:"available_date,omitempty" json:"available_date,omitempty"`
	ShowCondition           string               `xml:"show_condition,omitempty" json:"show_condition,omitempty"`
	Condition               string               `xml:"condition,omitempty" json:"condition,omitempty"`
	ShowPrice               string               `xml:"show_price,omitempty" json:"show_price,omitempty"`
	Indexed                 string               `xml:"indexed,omitempty" json:"indexed,omitempty"`
	Visibility              string               `xml:"visibility,omitempty" json:"visibility,omitempty"`
	AdvancedStockManagement string               `xml:"advanced_stock_management,omitempty" json:"advanced_stock_management,omitempty"`
	DateAdd                 string               `xml:"date_add,omitempty" json:"date_add,omitempty"`
	DateUpd                 string               `xml:"date_upd,omitempty" json:"date_upd,omitempty"`
	PackStockType           int                  `xml:"pack_stock_type,omitempty" json:"pack_stock_type,omitempty"`
	MetaDescription         *LanguageData        `xml:"meta_description,omitempty" json:"meta_description,omitempty"`
	MetaKeywords            *LanguageData        `xml:"meta_keywords,omitempty" json:"meta_keywords,omitempty"`
	MetaTitle               *LanguageData        `xml:"meta_title,omitempty" json:"meta_title,omitempty"`
	LinkRewrite             *LanguageData        `xml:"link_rewrite,omitempty" json:"link_rewrite,omitempty"`
	Names                   *LanguageData        `xml:"name,omitempty" json:"name,omitempty"`
	Description             *LanguageData        `xml:"description,omitempty" json:"description,omitempty"`
	DescriptionShort        *LanguageData        `xml:"description_short,omitempty" json:"description_short,omitempty"`
	AvailableNow            *LanguageData        `xml:"available_now,omitempty" json:"available_now,omitempty"`
	AvailableLater          *LanguageData        `xml:"available_later,omitempty" json:"available_later,omitempty"`
	Associations            *ProductAssociations `xml:"associations,omitempty" json:"associations,omitempty"`
}

type Categories struct {
	ID int `xml:"id,omitempty" json:"id,omitempty"`
}

type Images struct {
	ID int `xml:"id,omitempty" json:"id,omitempty"`
}

type Combinations struct {
	ID int `xml:"id,omitempty" json:"id,omitempty"`
}

type ProductOptionValues struct {
	ID int `xml:"id,omitempty" json:"id,omitempty"`
}

type ProductFeatures struct {
	ID             int `xml:"id,omitempty" json:"id,omitempty"`
	IDFeatureValue int `xml:"id_feature_value,omitempty" json:"id_feature_value,omitempty"`
}

type StockAvailables struct {
	ID                 int `xml:"id,omitempty" json:"id,omitempty"`
	IDProductAttribute int `xml:"id_product_attribute,omitempty" json:"id_product_attribute,omitempty"`
}

type ProductRow struct {
	Categories          *[]Categories          `xml:"categories,omitempty" json:"categories,omitempty"`
	Images              *[]Images              `xml:"images,omitempty" json:"images,omitempty"`
	Combinations        *[]Combinations        `xml:"combinations,omitempty" json:"combinations,omitempty"`
	ProductOptionValues *[]ProductOptionValues `xml:"product_option_values,omitempty" json:"product_option_values,omitempty"`
	ProductFeatures     *[]ProductFeatures     `xml:"product_features,omitempty" json:"product_features,omitempty"`
	StockAvailables     *[]StockAvailables     `xml:"stock_availables,omitempty" json:"stock_availables,omitempty"`
}

type ProductAssociations struct {
	ProductRows *[]ProductRow `xml:"product_rows>cart_row,omitempty" json:"cart_rows,omitempty"`
}

func (service *ProductService) Create(product *Product) (*Product, *http.Response, error) {
	createdProduct := new(Product)

	body := ResponseProduct{
		Xlink:   "http://www.w3.org/1999/xlink",
		Product: product,
	}

	_url := "products"
	req, _ := service.client.NewRequest("POST", _url, body)

	productsResponse := new(ResponseProduct)
	response, err := service.client.Do(req, productsResponse)

	if err != nil {
		return nil, response, err
	}

	if productsResponse != nil && productsResponse.Product != nil {
		createdProduct = productsResponse.Product
	}

	return createdProduct, response, nil
}

func (service *ProductService) Get(productID int, params *ServiceListParams) (*Product, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
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
		if productsResponse.ProductsData != nil && productsResponse.ProductsData.Products != nil &&
			len(*productsResponse.ProductsData.Products) > 0 {
			product = &(*productsResponse.ProductsData.Products)[0]
		}
	}

	return product, response, nil
}

func (service *ProductService) List(params *ServiceListParams) (*[]Product, *http.Response, error) {
	products := new([]Product)

	_url := makeResourceUrl("products", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	productsResponse := new(ResponseProduct)
	response, err := service.client.Do(req, productsResponse)

	if err != nil {
		return nil, response, err
	}

	if productsResponse != nil && productsResponse.ProductsData != nil &&
		productsResponse.ProductsData.Products != nil {
		products = productsResponse.ProductsData.Products
	}

	return products, response, nil
}
