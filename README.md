# Negotiator

This is a simple golang content negotiation library that was built as a
[Martini](http://martini.codegangsta.io/) middleware/handler but can be used 
standalone as well. The ContentNegotiator object will read the `Accepts` header 
from the `net/http` request and encode the given data appropriately.

## Supported Content-Types
* application/json
* application/xml

**NOTE**: The above encoders are included in this repo and will automatically be
setup if you use the `NewJsonXmlContentNegotiator` function but support for any
mimetype can be added dynamically.

## Usage

To use the built in JSON and XML encoders use the `NewJsonXmlContentNegotiator`
function to create the `ContentNegotiator`.  

If you dont want JSON and XML support use the `NewContentNegotiator` function
to create a base `ContentNegotiator`.  This negotiator will need at least one Encoder
to function. You can add encoders to this using the `AddEncoder` function as seen
in the example below.

```go
// Don't want to support XML? Use the following lines:
cn := negotiator.NewContentNegotiator(defaultEncoder, responseWriter)
cn.AddEncoder("application/json", negotiator.JsonEncoder{true})
```

### Creating Encoders

If you want to add support for additional mimetypes simple create a struct
that implements the `Encoder` interface.  Use the skeleton below as a starting
point.

```go
package main

type FooEncoder struct {}

func (foo FooEncoder) Encode(data interface{}) ([]byte, error) {
	// encode the data and return a byte array
}

// ContentType returns the string that will be used
// for the Content-Type header on the response
func (js JsonEncoder) ContentType() string {
	//return the appropriate Content-Type header string
	//return "application/foo; charset=utf-8"
}
```
Once you have an encoder add it to the content negotiator with the `AddEncoder`
function:

```go
// Pass in the Accepts header string to respond to and the encoder itself
cn.AddEncoder("application/foo", FooEncoder{})
```
Now if the client sends an `Accepts` header of `application/foo` the `FooEncoder`
will be used to encode the response.

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
	...
	
	output := ... //some struct of data

	// This creates a content negotiator can handle JSON and XML, defaults to json, and doesn't pretty print
	cn := negotiator.NewJsonXmlContentNegotiator(negotiator.JsonEncoder{false}, responseWriter, false)
	// To add your own mimetypes and encoders use the AddEncoder function:
	//cn.AddEncoder("text/html", htmlEncoder)
	log.Println(cn.Negotiate(request, output))
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
	// This creates a content negotiator can handle JSON and XML, defaults to json, and doesn't pretty print
	cn := negotiator.NewJsonXmlContentNegotiator(negotiator.JsonEncoder{false}, w, false)

	// To add your own mimetypes and encoders use the AddEncoder function:
	//cn.AddEncoder("text/html", htmlEncoder)
	
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
