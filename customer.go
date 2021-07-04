package stripe

import (
	"net/url"
)

// Customer represents a customer
type Customer struct {
	ID     string `json:"id,omitempty"`
	Object string `json:"object,omitempty"`

	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	Discount      *string `json:"discount,omitempty"`
	Email         *string `json:"email,omitempty"`
	DefaultSource *string `json:"default_source,omitempty"`
	Phone         *string `json:"phone,omitempty"`

	Metadata Dictionary `json:"metadata,omitempty"`
	Address  *Address   `json:"address,omitempty"`

	Balance  int64  `json:"balance,omitempty"`
	Currency string `json:"currency,omitempty"`

	InvoicePrefix       *string          `json:"invoice_prefix,omitempty"`
	InvoiceSettings     *InvoiceSettings `json:"invoice_settings,omitempty"`
	NextInvoiceSequence *int64           `json:"next_invoice_sequence,omitempty"`

	TaxExempt *string `json:"tax_exempt,omitempty"`

	Livemode   *bool `json:"livemode,omitempty"`
	Delinquent *bool `json:"delinquent,omitempty"`

	PreferredLocales []string   `json:"preferred_locales,omitempty"`
	Shipping         Dictionary `json:"shipping,omitempty"`

	Created int64 `json:"created,omitempty"`
}

func (c *Customer) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 6)
	setFormStringPtr(form, "name", c.Name)
	setFormStringPtr(form, "description", c.Description)
	setFormStringPtr(form, "discount", c.Discount)
	setFormStringPtr(form, "email", c.Email)
	setFormStringPtr(form, "default_source", c.DefaultSource)
	setFormStringPtr(form, "phone", c.Phone)
	setFormInt64(form, "balance", c.Balance)
	setFormString(form, "currency", c.Currency)
	setFormStringPtr(form, "invoice_prefix", c.InvoicePrefix)
	setFormInt64Ptr(form, "next_invoice_sequence", c.NextInvoiceSequence)
	setFormStringPtr(form, "tax_exempt", c.TaxExempt)
	setFormStringSlice(form, "preferred_locales", c.PreferredLocales)

	c.Metadata.AppendFormValues(form, "metadata")
	c.Address.AppendFormValues(form, "address")
	c.InvoiceSettings.AppendFormValues(form, "invoice_settings")
	c.Shipping.AppendFormValues(form, "shipping")
	return
}
