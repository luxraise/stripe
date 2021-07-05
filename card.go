package stripe

import (
	"net/url"
)

type Card struct {
	// System fields
	// ID of tokenized card
	ID string `json:"id"`
	// Object type (will be set as "card")
	Object string `json:"object"`

	// Required fields
	// Two-digit number representing the card's expiration month.
	ExpirationMonth int64 `json:"exp_month"`
	// Two or four-digit number representing the card's expiration year.
	ExpirationYear int64 `json:"exp_year"`
	// The card number, as a string without any separators.
	CardNumber string `json:"number"`

	// Usually required fields
	// Card security code. Highly recommended to always include this value, but it's required only for accounts based in European countries.
	CVC *string `json:"cvc"`

	// Optional fields
	// Cardholder's full name.
	CardholderName *string `json:"name"`
	// Address line 1 (Street address / PO Box / Company name).
	AddressLine1 *string `json:"address_line1"`
	// Address line 2 (Apartment / Suite / Unit / Building).
	AddressLine2 *string `json:"address_line2"`
	// City / District / Suburb / Town / Village.
	City *string `json:"address_city"`
	// State / County / Province / Region.
	State *string `json:"address_state"`
	// ZIP or postal code.
	Zipcode *string `json:"address_zip"`
	// Billing address country, if provided.
	Country *string `json:"address_country"`

	// Required in order to add the card to an account; in all other cases, this parameter is not used. When added to an account, the card (which must be a debit card) can be used as a transfer destination for funds in this currency.
	// Note: This is utilized for connect only
	Currency *string `json:"currency"`

	// Returned by system, not needed for creation/update
	Fingerprint string `json:"fingerprint"`
	LastFour    string `json:"last4"`
	Brand       string `json:"brand"`
	CVCCheck    string `json:"cvc_check"`
}

func (c *Card) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 4)
	c.AppendFormValues(form, "")
	return
}

func (c *Card) AppendFormValues(form url.Values, key string) {
	setFormInt64(form, getFieldKey(key, "exp_month"), c.ExpirationMonth)
	setFormInt64(form, getFieldKey(key, "exp_year"), c.ExpirationYear)
	setFormString(form, getFieldKey(key, "number"), c.CardNumber)
	setFormStringPtr(form, getFieldKey(key, "cvc"), c.CVC)
	setFormStringPtr(form, getFieldKey(key, "name"), c.CardholderName)
	setFormStringPtr(form, getFieldKey(key, "address_line_1"), c.AddressLine1)
	setFormStringPtr(form, getFieldKey(key, "address_line_2"), c.AddressLine2)
	setFormStringPtr(form, getFieldKey(key, "city"), c.City)
	setFormStringPtr(form, getFieldKey(key, "state"), c.State)
	setFormStringPtr(form, getFieldKey(key, "zip"), c.Zipcode)
	setFormStringPtr(form, getFieldKey(key, "country"), c.Country)
	setFormStringPtr(form, getFieldKey(key, "currency"), c.Currency)
}
