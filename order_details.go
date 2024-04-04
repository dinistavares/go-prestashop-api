package prestashop

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// OrderDetail service
type OrderDetailService service

type ResponseOrderDetail struct {
	XMLName          xml.Name         `xml:"prestashop,omitempty" json:"prestashop,omitempty"`
	Xlink            string           `xml:"xlink,attr,omitempty" json:"xlink,omitempty"`
	OrderDetail      *OrderDetail     `xml:"order_detail,omitempty" json:"order_detail,omitempty"`
	OrderDetailsData *OrderDetailData `xml:"order_details,omitempty" json:"order_details,omitempty"`
}

type OrderDetailData struct {
	OrderDetail *[]OrderDetail `xml:"order_detail,omitempty" json:"order_details,omitempty"`
}

type OrderDetail struct {
	ID                        int                      `xml:"id,omitempty" json:"id,omitempty"`
	IDOrder                   int                      `xml:"id_order,omitempty" json:"id_order,omitempty"`
	ProductID                 int                      `xml:"product_id,omitempty" json:"product_id,omitempty"`
	ProductAttributeID        int                      `xml:"product_attribute_id,omitempty" json:"product_attribute_id,omitempty"`
	ProductQuantityReinjected int                      `xml:"product_quantity_reinjected,omitempty" json:"product_quantity_reinjected,omitempty"`
	GroupReduction            string                   `xml:"group_reduction,omitempty" json:"group_reduction,omitempty"`
	DiscountQuantityApplied   int                      `xml:"discount_quantity_applied,omitempty" json:"discount_quantity_applied,omitempty"`
	DownloadHash              string                   `xml:"download_hash,omitempty" json:"download_hash,omitempty"`
	DownloadDeadline          string                   `xml:"download_deadline,omitempty" json:"download_deadline,omitempty"`
	IDOrderInvoice            int                      `xml:"id_order_invoice,omitempty" json:"id_order_invoice,omitempty"`
	IDWarehouse               int                      `xml:"id_warehouse,omitempty" json:"id_warehouse,omitempty"`
	IDShop                    int                      `xml:"id_shop,omitempty" json:"id_shop,omitempty"`
	IDCustomization           int                      `xml:"id_customization,omitempty" json:"id_customization,omitempty"`
	ProductName               string                   `xml:"product_name,omitempty" json:"product_name,omitempty"`
	ProductQuantity           int                      `xml:"product_quantity,omitempty" json:"product_quantity,omitempty"`
	ProductQuantityInStock    int                      `xml:"product_quantity_in_stock,omitempty" json:"product_quantity_in_stock,omitempty"`
	ProductQuantityReturn     int                      `xml:"product_quantity_return,omitempty" json:"product_quantity_return,omitempty"`
	ProductQuantityRefunded   int                      `xml:"product_quantity_refunded,omitempty" json:"product_quantity_refunded,omitempty"`
	ProductPrice              string                   `xml:"product_price,omitempty" json:"product_price,omitempty"`
	ReductionPercent          string                   `xml:"reduction_percent,omitempty" json:"reduction_percent,omitempty"`
	ReductionAmount           string                   `xml:"reduction_amount,omitempty" json:"reduction_amount,omitempty"`
	ReductionAmountTaxIncl    string                   `xml:"reduction_amount_tax_incl,omitempty" json:"reduction_amount_tax_incl,omitempty"`
	ReductionAmountTaxExcl    string                   `xml:"reduction_amount_tax_excl,omitempty" json:"reduction_amount_tax_excl,omitempty"`
	ProductQuantityDiscount   string                   `xml:"product_quantity_discount,omitempty" json:"product_quantity_discount,omitempty"`
	ProductEan13              string                   `xml:"product_ean13,omitempty" json:"product_ean13,omitempty"`
	ProductISBN               string                   `xml:"product_isbn,omitempty" json:"product_isbn,omitempty"`
	ProductUpc                string                   `xml:"product_upc,omitempty" json:"product_upc,omitempty"`
	ProductMpn                string                   `xml:"product_mpn,omitempty" json:"product_mpn,omitempty"`
	ProductReference          string                   `xml:"product_reference,omitempty" json:"product_reference,omitempty"`
	ProductSupplierReference  string                   `xml:"product_supplier_reference,omitempty" json:"product_supplier_reference,omitempty"`
	ProductWeight             string                   `xml:"product_weight,omitempty" json:"product_weight,omitempty"`
	TaxComputationMethod      int                      `xml:"tax_computation_method,omitempty" json:"tax_computation_method,omitempty"`
	IDTaxRulesGroup           int                      `xml:"id_tax_rules_group,omitempty" json:"id_tax_rules_group,omitempty"`
	Ecotax                    string                   `xml:"ecotax,omitempty" json:"ecotax,omitempty"`
	EcotaxTaxRate             string                   `xml:"ecotax_tax_rate,omitempty" json:"ecotax_tax_rate,omitempty"`
	DownloadNb                int                      `xml:"download_nb,omitempty" json:"download_nb,omitempty"`
	UnitPriceTaxIncl          string                   `xml:"unit_price_tax_incl,omitempty" json:"unit_price_tax_incl,omitempty"`
	UnitPriceTaxExcl          string                   `xml:"unit_price_tax_excl,omitempty" json:"unit_price_tax_excl,omitempty"`
	TotalPriceTaxIncl         string                   `xml:"total_price_tax_incl,omitempty" json:"total_price_tax_incl,omitempty"`
	TotalPriceTaxExcl         string                   `xml:"total_price_tax_excl,omitempty" json:"total_price_tax_excl,omitempty"`
	TotalShippingPriceTaxExcl string                   `xml:"total_shipping_price_tax_excl,omitempty" json:"total_shipping_price_tax_excl,omitempty"`
	TotalShippingPriceTaxIncl string                   `xml:"total_shipping_price_tax_incl,omitempty" json:"total_shipping_price_tax_incl,omitempty"`
	PurchaseSupplierPrice     string                   `xml:"purchase_supplier_price,omitempty" json:"purchase_supplier_price,omitempty"`
	OriginalProductPrice      string                   `xml:"original_product_price,omitempty" json:"original_product_price,omitempty"`
	OriginalWholesalePrice    string                   `xml:"original_wholesale_price,omitempty" json:"original_wholesale_price,omitempty"`
	TotalRefundedTaxExcl      string                   `xml:"total_refunded_tax_excl,omitempty" json:"total_refunded_tax_excl,omitempty"`
	TotalRefundedTaxIncl      string                   `xml:"total_refunded_tax_incl,omitempty" json:"total_refunded_tax_incl,omitempty"`
	Associations              *OrderDetailAssociations `xml:"associations,omitempty" json:"associations,omitempty"`
}

type OrderDetailAssociations struct {
	Taxes *[]OrderDetailTax `xml:"taxes>tax,omitempty" json:"taxes,omitempty"`
}

type OrderDetailTax struct {
	Href string `xml:"href,attr,omitempty" json:"href,attr,omitempty"`
	ID   int    `xml:"id,omitempty" json:"id,omitempty"`
}

func (service *OrderDetailService) Create(order_details *OrderDetail) (*OrderDetail, *http.Response, error) {
	createdOrderDetail := new(OrderDetail)

	body := ResponseOrderDetail{
		Xlink:       "http://www.w3.org/1999/xlink",
		OrderDetail: order_details,
	}

	_url := "order_details"
	req, _ := service.client.NewRequest("POST", _url, body)

	order_detailssResponse := new(ResponseOrderDetail)
	response, err := service.client.Do(req, order_detailssResponse)

	if err != nil {
		return nil, response, err
	}

	if order_detailssResponse != nil && order_detailssResponse.OrderDetail != nil {
		createdOrderDetail = order_detailssResponse.OrderDetail
	}

	return createdOrderDetail, response, nil
}

func (service *OrderDetailService) Get(order_detailsID int, params *ServiceListParams) (*OrderDetail, *http.Response, error) {
	params = setDefaultResourceByIDDisplayParams(params)
	resourceRoute := fmt.Sprintf("order_details/%d", order_detailsID)

	_url := makeResourceUrl(resourceRoute, params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	order_details := new(OrderDetail)
	order_detailssResponse := new(ResponseOrderDetail)
	response, err := service.client.Do(req, order_detailssResponse)

	if err != nil {
		return nil, response, err
	}

	if order_detailssResponse != nil {
		if order_detailssResponse.OrderDetail != nil {
			order_details = order_detailssResponse.OrderDetail
		}

		// Use fisrt matching order_details
		if order_detailssResponse.OrderDetailsData != nil && order_detailssResponse.OrderDetailsData.OrderDetail != nil &&
			len(*order_detailssResponse.OrderDetailsData.OrderDetail) > 0 {
			order_details = &(*order_detailssResponse.OrderDetailsData.OrderDetail)[0]
		}
	}

	return order_details, response, nil
}

func (service *OrderDetailService) List(params *ServiceListParams) (*[]OrderDetail, *http.Response, error) {
	order_details := new([]OrderDetail)
	_url := makeResourceUrl("order_details", params)
	req, _ := service.client.NewRequest("GET", _url, nil)

	order_detailssResponse := new(ResponseOrderDetail)
	response, err := service.client.Do(req, order_detailssResponse)

	if err != nil {
		return nil, response, err
	}

	if order_detailssResponse != nil && order_detailssResponse.OrderDetailsData != nil &&
		order_detailssResponse.OrderDetailsData.OrderDetail != nil {
		order_details = order_detailssResponse.OrderDetailsData.OrderDetail
	}

	return order_details, response, nil
}
