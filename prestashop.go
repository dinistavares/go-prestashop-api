package prestashop

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultAuthHeaderName        = "Authorization"
	acceptedContentType          = "application/xml"
	userAgent                    = "go-prestashop-api/1.1"
	clientRequestRetryAttempts   = 2
	clientRequestRetryHoldMillis = 1000

	defaultData
)

var errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
var errorDoAttemptNilRequest = errors.New("request could not be constructed")

type ClientConfig struct {
	HttpClient      *http.Client
	RestEndpointURL string
}

type auth struct {
	HeaderName string
	ApiKey     string
}

type Client struct {
	config  *ClientConfig
	client  *http.Client
	auth    *auth
	baseURL *url.URL

	Carrier    *CarrierService
	Currency   *CurrencyService
	Customer   *CustomersService
	Cart       *CartService
	Order      *OrderService
	OrderState *OrderStateService
	Product    *ProductService
}

type service struct {
	client *Client
}

type errorResponse struct {
	XMLName xml.Name `xml:"prestashop"`
	Xlink   string   `xml:"xlink,attr"`

	Response *http.Response
	RawError string

	Errors []ErrorData `xml:"errors" json:"errors"`
}

type ErrorData struct {
	Errors []Error `xml:"error"`
}

type Error struct {
	Code    int    `xml:"code" json:"code"`
	Message string `xml:"message" json:"message"`
}

func (response *errorResponse) Error() string {
	if response.Errors == nil && response.RawError != "" {
		errorNew := Error{
			Message: response.RawError,
			Code:    -1,
		}

		response.Errors = append(response.Errors, ErrorData{Errors: []Error{errorNew}})
	}

	return fmt.Sprintf("%v %v: %d %+v",
		response.Response.Request.Method, response.Response.Request.URL,
		response.Response.StatusCode, response.Errors)
}

func New(shopURL string) (*Client, error) {
	if shopURL == "" {
		return nil, errors.New("store url is required")
	}

	config := ClientConfig{
		HttpClient:      http.DefaultClient,
		RestEndpointURL: shopURL,
	}

	return NewWithConfig(&config)
}

func NewWithConfig(config *ClientConfig) (*Client, error) {
	if config == nil {
		return nil, errors.New("configuration is empty")
	}

	// Create client
	baseURL, err := url.Parse(config.RestEndpointURL + "/api/")

	if err != nil {
		return nil, err
	}

	client := &Client{config: config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL}

	// Map services
	client.Carrier = &CarrierService{client: client}
	client.Currency = &CurrencyService{client: client}
	client.Customer = &CustomersService{client: client}
	client.Cart = &CartService{client: client}
	client.Order = &OrderService{client: client}
	client.OrderState = &OrderStateService{client: client}
	client.Product = &ProductService{client: client}

	return client, nil
}

// Authenticate saves authenitcation parameters for user
func (client *Client) Authenticate(webserviceKey string) {
	client.auth.HeaderName = defaultAuthHeaderName

	client.auth.ApiKey = "Basic " + base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{webserviceKey, ""}, ":")))
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := client.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := xml.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.auth.HeaderName, client.auth.ApiKey)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("Content-type", acceptedContentType)

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	var lastErr error

	attempts := 0

	for attempts < clientRequestRetryAttempts {
		// Hold before this attempt? (ie. not first attempt)
		if attempts > 0 {
			time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
		}

		// Dispatch request attempt
		attempts++
		resp, shouldRetry, err := client.doAttempt(req, v)

		// Return response straight away? (we are done)
		if !shouldRetry {
			return resp, err
		}

		// Should retry: store last error (we are not done)
		lastErr = err
	}

	// Set default error? (all attempts failed, but no error is set)
	if lastErr == nil {
		lastErr = errorDoAllAttemptsExhausted
	}

	// All attempts failed, return last attempt error
	return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*http.Response, bool, error) {
	if req == nil {
		return nil, false, errorDoAttemptNilRequest
	}

	resp, err := client.client.Do(req)

	if checkRequestRetry(resp, err) {
		return nil, true, err
	}

	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, false, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = xml.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return resp, false, err
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
	// Low-level error
	if err != nil {
		return true
	}

	// No low-level error (should not retry)
	return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &errorResponse{Response: response}

	data, err := io.ReadAll(response.Body)

	if err == nil && data != nil {
		xml.Unmarshal(data, errorResponse)
	}

	errorResponse.RawError = string(data)

	return errorResponse
}
