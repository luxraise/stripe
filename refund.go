package stripe

import "net/url"

const (
	RefundReasonDuplicate         = "duplicate"
	RefundReasonFraudulent        = "fraudulent"
	RefundReasonRequestByCustomer = "requested_by_customer"

	RefundStatusPending   = "pending"
	RefundStatusSucceeded = "succeeded"
	RefundStatusFailed    = "failed"
	RefundStatusCanceled  = "canceled"
)

type RefundRequest struct {
	// The identifier of the charge to refund.
	Charge string `json:"charge"`
	// The amount to refund in the smallest currency unit
	// Example: To refund $1.59 USD it would be 159
	Amount int64 `json:"amount"`

	// A set of key-value pairs that you can attach to a Refund object. This can be useful for storing additional information about the refund in a structured format. You can unset individual keys if you POST an empty value for that key. You can clear all keys if you POST an empty value for metadata
	Metadata Dictionary `json:"metadata"`

	// ID of the PaymentIntent to refund.
	PaymentIntent *string `json:"payment_intent"`

	// String indicating the reason for the refund. If set, possible values are duplicate, fraudulent, and requested_by_customer. If you believe the charge to be fraudulent, specifying fraudulent as the reason will add the associated card and email to your block lists, and will also help us improve our fraud detection algorithms.
	Reason *string `json:"reason"`
}

func (r *RefundRequest) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 5)
	setFormString(form, "charge", r.Charge)
	setFormInt64(form, "amount", r.Amount)
	setFormStringPtr(form, "payment_intent", r.PaymentIntent)
	setFormStringPtr(form, "reason", r.Reason)
	r.Metadata.AppendFormValues(form, "metadata")
	return
}

type Refund struct {
	ID     string `json:"id"`
	Object string `json:"object"`

	RefundRequest

	BalanceTransaction *string `json:"balance_transaction"`
	Currency           string  `json:"currency"`

	Status  string `json:"status"`
	Created int64  `json:"created"`

	ReceiptNumber          *string `json:"receipt_number"`
	SourceTransferReversal *string `json:"source_transfer_reversal"`
	TransferReversal       *string `json:"transfer_reversal"`
}
