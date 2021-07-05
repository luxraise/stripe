package stripe

import "net/url"

type sourceRequest struct {
	Source string `json:"source"`
}

func (s *sourceRequest) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 1)
	setFormString(form, "source", s.Source)
	return
}
