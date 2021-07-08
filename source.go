package stripe

import (
	"encoding/json"
	"net/url"
)

type Source string

func (s *Source) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err == nil {
		*s = Source(str)
		return
	}

	var c Card
	// TODO: Add universal obj which switches based on `Object` type
	if err = json.Unmarshal(bs, &c); err != nil {
		return
	}

	*s = Source(c.ID)
	return
}

type sourceRequest struct {
	Source string `json:"source"`
}

func (s *sourceRequest) ToFormValues() (form url.Values) {
	// Pre-allocate values with enough space for the basic user information rows
	form = make(url.Values, 1)
	setFormString(form, "source", s.Source)
	return
}
