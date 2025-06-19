// Package prep provides predefined sanitisers
package prep

type Sanitiser func(string) string

func None(v string) string {
	return v
}
