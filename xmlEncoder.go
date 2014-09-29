package negotiator

import (
	"encoding/xml"
)

type XmlEncoder struct {
	PrettyPrint bool
}

func (xe XmlEncoder) Encode(data interface{}) ([]byte, error) {
	if xe.PrettyPrint {
		return xml.MarshalIndent(data, "", "  ")
	} else {
		return xml.Marshal(data)
	}
}

func (js JsonEncoder) ContentType() string {
	return "application/xml; charset=utf-8"
}
