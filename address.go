package stripe

import "net/url"

type Address struct {
	// City, district, suburb, town, or village. (Optional)
	City string `json:"city"`
	// City, district, suburb, town, or village. (Optional)
	Country string `json:"country"`
	// City, district, suburb, town, or village. (Optional)
	Line1 string `json:"line1"`
	// City, district, suburb, town, or village. (Optional)
	Line2 string `json:"line2"`
	// City, district, suburb, town, or village. (Optional)
	PostalCode string `json:"postal_code"`
	// City, district, suburb, town, or village. (Optional)
	State string `json:"state"`
}

func (a *Address) AppendFormValues(values url.Values, key string) {
	if a == nil {
		return
	}

	// TODO: Implement this
}
