package stripe

import (
	"errors"
	"fmt"
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

	endpointCustomers              = "/customers"
	endpointCustomersWithID        = "/customers/%s"
	endpointTokens                 = "/tokens"
	endpointSourcesWithID          = "/customers/%s/sources"
	endpointSourcesWithIDAndCardID = "/customers/%s/sources/%s"
	endpointCharges                = "/charges"
	endpointRefunds                = "/refunds"
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

func (c *Client) RemoveCustomer(stripeUserID string) (err error) {
	endpoint := fmt.Sprintf(endpointCustomersWithID, stripeUserID)
	err = c.request("DELETE", endpoint, nil, nil)
	return
}

func (c *Client) AddCreditCard(stripeUserID string, card Card) (created Card, err error) {
	var token Token
	if token, err = c.createCardToken(card); err != nil {
		err = fmt.Errorf("error creating card token: %v", err)
		return
	}

	var req sourceRequest
	req.Source = token.ID

	endpoint := fmt.Sprintf(endpointSourcesWithID, stripeUserID)
	err = c.request("POST", endpoint, &req, &created)
	return
}

func (c *Client) ListCards(stripeUserID string) (cards []Card, err error) {
	var resp listCardsResponse
	endpoint := fmt.Sprintf(endpointSourcesWithID, stripeUserID)
	if err = c.request("GET", endpoint, nil, &resp); err != nil {
		return
	}

	cards = resp.Data
	return
}

func (c *Client) RemoveCreditCard(stripeUserID, cardID string) (err error) {
	endpoint := fmt.Sprintf(endpointSourcesWithIDAndCardID, stripeUserID, cardID)
	err = c.request("DELETE", endpoint, nil, nil)
	return
}

func (c *Client) CreateCharge(stripeUserID string, charge Charge) (created Charge, err error) {
	charge.StripeUserID = stripeUserID
	err = c.request("POST", endpointCharges, &charge, &created)
	return
}

func (c *Client) CreateRefund(request RefundRequest) (refund Refund, err error) {
	err = c.request("POST", endpointRefunds, &request, &refund)
	return
}

func (c *Client) createCardToken(card Card) (created Token, err error) {
	var token Token
	token.Card = card
	err = c.request("POST", endpointTokens, &token, &created)
	return
}

func (c *Client) request(method, endpoint string, request Request, response interface{}) (err error) {
	var req *http.Request
	body := getRequestBody(request)
	url := c.getURL(method, endpoint)
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
	case 400, 404:
		return handleError(resp.Body)
	case 401:
		return ErrUnauthorized

	default:
		return fmt.Errorf("Unexpected status code of: %d (url: <%s>, method: <%s>)", resp.StatusCode, url, method)
	}
}

func (c *Client) getURL(method, endpoint string) string {
	u := *c.u
	u.Path = path.Join(apiVersion, endpoint)
	return u.String()
}
