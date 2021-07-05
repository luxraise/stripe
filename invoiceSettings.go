package stripe

import "net/url"

// InvoiceSettings represent Customer invoice settings
type InvoiceSettings struct {
	CustomFields         Dictionary `json:"custom_fields"`
	DefaultPaymentMethod *string    `json:"default_payment_method"`
	Footer               *string    `json:"footer"`
}

func (i *InvoiceSettings) AppendFormValues(values url.Values, key string) {
	if i == nil {
		return
	}

	// TODO: Implement this
}
