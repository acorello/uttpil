package uttpil

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/acorello/uttpil/prep"
)

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
func (my UrlValuesHelper) Give(key string, sanitizer prep.Sanitiser, consumer func(string) error) {
	if !my.Has(key) {
		my.joinError(key, fmt.Errorf("not found"))
		return
	}
	value := my.Get(key)
	value = sanitizer(value)
	err := consumer(value)
	my.joinError(key, err)
}

func (my UrlValuesHelper) Get(name string, sanitizers ...prep.Sanitiser) string {
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
