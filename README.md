# UTTPIL

A collection of minimalist helpers for implementing HTTP servers using Go standard library.

Named as a pronounceable blend of 'HTTP' and 'Utilities'

⚠️ EXPERIMENTAL ⚠️

## The `ForMethod` http-handler

There is a one-to-many relationship between a URL path and the HTTP methods it could serve.

`uttpil.ForMethod` is a struct implementing `http.Handler` (i.e. `ServeHTTP(http.ResponseWriter, *http.Request)`) having one field for each HTTP method.

An `http.HandlerFunc` should be assigned to each field/method you want to support for the path. `ForMethod` returns `501 NotImplemented` when no handler is available for the requested method.

```go
mux.Handle("/resource/path", uttpil.ForMethod{
    GET: func(w http.ResponseWriter, r *http.Request) {
		// handle it inline
    },
    POST: h.Post, // define your handler elsewhere
})
```

## The `LoggingHandler`

Minimal middleware to log which path and method has been served.

Example:

```go
mux.Handle("/resource/path", uttipl.LoggingHandler(uttpil.ForMethod{
    POST: h.Post,
    // et al.
}))
```

Will log `serving '/resource/path <METHOD>'` on stdout (for any method, not just `POST`)

## The `UrlValuesHelper`

- provides the method `Give(key string, consumer func(value string) error)`, which is convenient when parsing many form fields if we want to collect every error. Collected errors will be prefixed with the field name, so we can reuse consumer and avoid repeating the field name.  
  If the field is absent, the consumer will NOT be called and a ErrFieldNotFound is recorded.

```go
form, _err := uttpil.NewUrlValuesHelper(r) // confusingly, the stdlib decodes a form into url.Values,
form.Give("formField", func(val string) error {
    val = strings.TrimSpace(val) // sanitise
    if val == "" {
		return fmt.Errorf("blank")
    }
    if id, err := resource.Parse(val); err != nil {
        return err
    } else {
        c.Id = id
        return nil
    }
})
// …
var errors map[string]error // map is the field for which the consumer returned an error (eg. `formField`)
errors = form.Errors()
// may include {"formField": `field "formField" has errors: blank`}
// decide what to do
```

- Overrides `http.UrlValues.Get(name string)` with `Get(name string, sanitizers ...func(string) string) string`; allowing to sanitize or normalize the value format before returning it.  
  Example:
  ```go
  name := form.Get("Name", strings.TrimSpace, strings.ToTitle)
  ``` 
