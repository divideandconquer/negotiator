package negotiator

import (
	"net/http"
	"strings"
)

const MimeJSON = "application/json"
const MimeXML = "application/xml"

// A Negotiator can Negotiate to determine what content type to convert
// a struct into for client consumption
type Negotiator interface {
	Negotiate(req *http.Request, data interface{}) ([]byte, error)
}

// Encoder is an interface for a struct that can encode data into []byte
type Encoder interface {
	Encode(data interface{}) ([]byte, error)
}

// ContentNegotiator is a Neotiator that supports pretty printing
// and a fallback/default encoder
type ContentNegotiator struct {
	PrettyPrint    bool
	DefaultEncoder Encoder
}

// Negotiate inspects the request for the accept header and
// encodes the response appropriately.
func (cn ContentNegotiator) Negotiate(req *http.Request, data interface{}) ([]byte, error) {
	var e = cn.DefaultEncoder

	accept := req.Header.Get("Accept")
	if strings.Contains(accept, MimeXML) {
		//get the xml encoder
		e = XmlEncoder{cn.PrettyPrint}
	} else if strings.Contains(accept, MimeJSON) {
		//default to json
		e = JsonEncoder{cn.PrettyPrint}
	}
	return e.Encode(data)
}

// Check for an error and panic
func Must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}
