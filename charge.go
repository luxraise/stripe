package stripe

import "net/url"

type Charge struct {
	// System fields
	// ID of tokenized card
	ID string `json:"id"`
	// Object type (will be set as "card")
	Object string `json:"object"`
	// Balance transaction ID
	BalanceTransaction string `json:"balance_transaction"`
	// Captured state
	Captured bool `json:"captured"`
	// Disputed state
	Disputed bool `json:"disputed"`
	// Paid state
	Paid bool `json:"paid"`

	// Required fields
	// The amount to charge in the smallest currency unit
	// Example: To charge $1.59 USD it would be 159
	Amount int64 `json:"amount"`
	// The ISO country code for the currency to be used
	// Example: usd
	Currency string `json:"currency"`
	// Stripe customer to charge
	StripeUserID string `json:"stripeUserID"`
	// The token source to charge
	Source Source `json:"source"`

	// Optional fields
	// Description of charge
	Description *string `json:"description"`
	// Custom metadata for the charge
	Metadata Dictionary `json:"metadata"`
}

func (c *Charge) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 4)
	setFormInt64(form, "amount", c.Amount)
	setFormString(form, "currency", c.Currency)
	setFormString(form, "customer", c.StripeUserID)
	setFormString(form, "source", string(c.Source))
	setFormStringPtr(form, "description", c.Description)
	c.Metadata.AppendFormValues(form, "metadata")
	return
}
