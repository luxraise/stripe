package stripe

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

var (
	// ErrEmptyAPIKey is returned when a Client is initialized with an empty API key
	ErrEmptyAPIKey = errors.New("invalid API key, cannot be empty")
	// ErrUnauthorized is returned when a 401 status code is encountered
	ErrUnauthorized = errors.New("unauthorized, 401 status code encountered")
)

const (
	host       = "https://api.stripe.com"
	apiVersion = "v1"

	endpointCustomers       = "/customers"
	endpointCustomersWithID = "/customers/%s"
	endpointTokens          = "/tokens"
)

// New initializes and returns a new Stripe Client
func New(apiKey string) (client *Client, err error) {
	if len(apiKey) == 0 {
		err = ErrEmptyAPIKey
		return
	}

	var c Client
	if c.u, err = url.Parse(host); err != nil {
		return
	}

	c.apiKey = apiKey
	client = &c
	return
}

type Client struct {
	hc http.Client
	u  *url.URL

	apiKey string
}

func (c *Client) CreateCustomer(customer Customer) (created Customer, err error) {
	err = c.request("POST", endpointCustomers, &customer, &created)
	return
}

func (c *Client) GetCustomer(stripeUserID string) (customer Customer, err error) {
	endpoint := fmt.Sprintf(endpointCustomersWithID, stripeUserID)
	err = c.request("GET", endpoint, nil, &customer)
	return
}

func (c *Client) UpdateCustomer(stripeUserID string, customer Customer) (updated Customer, err error) {
	endpoint := fmt.Sprintf(endpointCustomersWithID, stripeUserID)
	err = c.request("POST", endpoint, &customer, &updated)
	return
}

func (c *Client) AddCreditCard(stripeUserID string) (cardID string, err error) {
	//err = c.request("POST", endpointCustomers, customer, &created)
	return
}

func (c *Client) RemoveCreditCard(stripeUserID, cardID string) (err error) {
	//err = c.request("POST", endpointCustomers, customer, &created)
	return
}

func (c *Client) CreateCardToken(card Card) (created Token, err error) {
	var token Token
	token.Card = card
	err = c.request("POST", endpointTokens, &token, &created)
	return
}

func (c *Client) request(method, endpoint string, request Request, response interface{}) (err error) {
	var body io.Reader
	if body, err = getRequestBody(request); err != nil {
		return
	}

	url := c.getURL(method, endpoint)

	var req *http.Request
	if req, err = http.NewRequest(method, url, body); err != nil {
		err = fmt.Errorf("error creating request: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp *http.Response
	if resp, err = c.hc.Do(req); err != nil {
		err = fmt.Errorf("error performing request: %v", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return handleResponse(resp.Body, response)
	case 400:
		return handleError(resp.Body)
	case 401:
		return ErrUnauthorized

	default:
		return fmt.Errorf("Unexpected status code of: %d", resp.StatusCode)
	}
}

func (c *Client) getURL(method, endpoint string) string {
	u := *c.u
	u.Path = path.Join(apiVersion, endpoint)
	return u.String()
}
