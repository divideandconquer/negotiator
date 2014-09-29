package negotiator

import (
	"encoding/xml"
)

type JsonEncoder struct {
	PrettyPrint bool
}

func (je JsonEncoder) Encode(data interface{}) ([]byte, error) {
	if je.PrettyPrint {
		return json.MarshalIndent(data, "", "  ")
	} else {
		return json.Marshal(data)
	}
}
