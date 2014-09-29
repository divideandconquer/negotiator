package negotiator

import (
	"net/http"
	"strings"
)

const MimeJSON = "application/json"
const MimeXML = "application/xml"

// Encoder is an interface for a struct that can encode data into []byte
type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	ContentType() string
}

// A Negotiator can Negotiate to determine what content type to convert
// a struct into for client consumption
type Negotiator interface {
	Negotiate(req *http.Request, data interface{}) ([]byte, error)
}

// ContentNegotiator is a Neotiator that supports pretty printing
// and a fallback/default encoder
type ContentNegotiator struct {
	PrettyPrint    bool
	DefaultEncoder Encoder
	ResponseWriter http.ResponseWriter
}

// Negotiate inspects the request for the accept header and
// encodes the response appropriately.
func (cn ContentNegotiator) Negotiate(req *http.Request, data interface{}) ([]byte, error) {
	var e = cn.getEncoder(req)
	cn.ResponseWriter.Header().Set("Content-Type", e.ContentType())
	return e.Encode(data)
}

// getEncoder parses the Accept header an returns the appropriate encoder to use
func (cn ContentNegotiator) getEncoder(req *http.Request) {
	var result = cn.DefaultEncoder
	accept := req.Header.Get("Accept")
	if strings.Contains(accept, MimeXML) {
		//get the xml encoder
		result = XmlEncoder{cn.PrettyPrint}
	} else if strings.Contains(accept, MimeJSON) {
		//default to json
		result = JsonEncoder{cn.PrettyPrint}
	}
	return result
}

// Check for an error and panic
func Must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}
