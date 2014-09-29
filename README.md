# Negotiator

This is a simple golang content negotiation library that was built as a
[Martini](http://martini.codegangsta.io/) middleware/handler but can be used 
standalone as well. The ContentNegotiator object will read the accept header 
from the `net/http` request and encode the given data appropriately.

## Usage

### General 

For general usage, i.e. non-martini use, simply import the package, create a
`ContentNegotiator` object, and call Negotiate:

```go
package main

import (
	"github.com/divideandconquer/negotiator"
	"net/http"
	"log"
)

func main() {
	w := http.ResponseWriter{}
	r := http.Request{}

	output := ... //some struct of data

	// This creates a content negotiator that defaults to json and doesn't pretty print
	cn := negotiator.ContentNegotiator{false, negotiator.JsonEncoder{false}, w}
	log.Println(cn.Negotiate(r, output))
}

```

### Martini

For use with the [Martini](http://martini.codegangsta.io/) framework add the content
negotiator to the list of middlewares to use:

```go
//Martini initialization
m = martini.New()

// add middleware
m.Use(martini.Recovery())
m.Use(martini.Logger())

// setup content negotiation
m.Use(func(c martini.Context, w http.ResponseWriter) {
	cn := negotiator.ContentNegotiator{true, negotiator.JsonEncoder{true}, w}
	c.MapTo(cn, (*negotiator.Negotiator)(nil))
})

// setup router
router := martini.NewRouter()
router.Get("/", func(r *http.Request, neg negotiator.Negotiator) (int, []byte) {
	data := ... //setup whatever data you want to return
	return http.StatusOK, negotiator.Must(neg.Negotiate(r, data)))
}
```

# License
This module is licensed using the Apache-2.0 License:

```
Copyright (c) 2014, Kyle Boorky
```
