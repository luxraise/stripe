package stripe

import "net/url"

// Token represents a stripe card token
type Token struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Type   string `json:"type"`

	Card Card `json:"card"`

	ClientIP *string `json:"client_ip"`
	Livemode *bool   `json:"livemode"`
	Used     *bool   `json:"used"`

	Created int64 `json:"created"`
}

func (t *Token) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 1)
	t.Card.AppendFormValues(form, "card")
	return
}
