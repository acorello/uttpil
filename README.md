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
