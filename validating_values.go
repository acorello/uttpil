package uttpil

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

type UrlValues struct {
	url.Values
}

func NewUrlValues(r *http.Request) (values UrlValues, err error) {
	err = r.ParseForm()
	if err != nil {
		return values, err
	}
	return UrlValues{
		Values: r.Form,
	}, nil
}

func (my UrlValues) Trim(name string) string {
	v := my.Get(name)
	v = strings.TrimSpace(v)
	return v
}

func (my UrlValues) IntOrPanic(name string, defaultIfBlank int) int {
	v := my.Trim(name)
	if v == "" {
		return defaultIfBlank
	}
	if i, err := strconv.Atoi(v); err != nil {
		panic(fmt.Errorf("failed to parse %q: %v", name, err))
	} else {
		return i
	}
}

func (my UrlValues) Trim_NotBlank(name string) (string, error) {
	v := my.Trim(name)
	if v == "" {
		return "", fmt.Errorf("%q is blank or empty", name)
	}
	return v, nil
}

type UrlValuesHelper struct {
	url.Values
	errors map[string]error
}

func NewUrlValuesHelper(r *http.Request) (values UrlValuesHelper, err error) {
	err = r.ParseForm()
	if err != nil {
		return values, err
	}
	return UrlValuesHelper{
		Values: r.Form,
		errors: make(map[string]error),
	}, nil
}

// Give provides value for `key` to the parser. Error returned is collected and all errors
// should be retrieved with Errors()
//
// Example:
//
//	    var person Person
//
//		form.Give("FirstName", func(name string) error {
//			name = strings.TrimSpace(name) // sanitise
//			if name == "" {				   // validate
//				return fmt.Errorf("blank")
//			}
//			person.FirstName = name
//			return nil
//		})
//
//		form.Errors()
func (my UrlValuesHelper) Give(key string, parser func(string) error) {
	if !my.Has(key) {
		my.joinError(key, fmt.Errorf("not found"))
		return
	}
	err := parser(my.Get(key))
	my.joinError(key, err)
}

func (my UrlValuesHelper) Get(name string, sanitizers ...func(string) string) string {
	v := my.Values.Get(name)
	for s := range slices.Values(sanitizers) {
		v = s(v)
	}
	return v
}

func (my UrlValuesHelper) Errors() map[string]error {
	return my.errors
}

func (my UrlValuesHelper) joinError(key string, err error) {
	my.errors[key] = errors.Join(my.errors[key], err)
}
