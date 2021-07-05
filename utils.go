package stripe

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

type Dictionary map[string]string

func (d Dictionary) AppendFormValues(values url.Values, key string) {
	for fieldKey, value := range d {
		headerKey := fmt.Sprintf("%s[%s]", key, fieldKey)
		values.Set(headerKey, value)
	}
}

type Request interface {
	ToFormValues() url.Values
}

// String will return a string pointer
func String(str string) *string {
	return &str
}

type listCardsResponse struct {
	Object  string `json:"object"`
	URL     string `json:"url"`
	HasMore bool   `json:"has_more"`
	Data    []Card `json:"data"`
}

func getRequestBody(request Request) (body io.Reader, err error) {
	if request == nil {
		return
	}

	encoded := request.ToFormValues().Encode()
	body = strings.NewReader(encoded)
	return
}

func handleResponse(r io.Reader, value interface{}) (err error) {
	if value == nil {
		return
	}

	if err = json.NewDecoder(r).Decode(value); err != nil {
		return fmt.Errorf("error encountered while attempting to decode response as JSON: %v", err)
	}

	return
}

func handleError(r io.Reader) (err error) {
	var value ErrorResponse
	if err = handleResponse(r, &value); err != nil {
		return
	}

	fmt.Printf("Value? %+v\n", value)

	return &value.Error
}

func setFormStringSlice(form url.Values, key string, values []string) {
	if len(values) == 0 {
		return
	}

	for i, value := range values {
		fieldKey := fmt.Sprintf("%s[%d]", key, i)
		form.Set(fieldKey, value)
	}
}

func setFormString(form url.Values, key, value string) {
	if len(value) == 0 {
		return
	}

	form.Set(key, value)
}

func setFormStringPtr(form url.Values, key string, value *string) {
	if value == nil {
		return
	}

	form.Set(key, *value)
}

func setFormInt64(form url.Values, key string, value int64) {
	form.Set(key, strconv.FormatInt(value, 10))
}

func setFormInt64Ptr(form url.Values, key string, value *int64) {
	if value == nil {
		return
	}

	setFormInt64(form, key, *value)
}

func getFieldKey(key, field string) string {
	if len(key) == 0 {
		return field
	}

	return fmt.Sprintf("%s[%s]", key, field)
}
