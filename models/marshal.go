package models

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

type InterfaceMap map[string]interface{}

func (this InterfaceMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}
	for key, value := range this {
		this.switchValue(&tokens, key, &value)
	}

	tokens = append(tokens, xml.EndElement{start.Name})
	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (this InterfaceMap) switchValue(tokens *[]xml.Token, key string, value *interface{}) {
	t := xml.StartElement{Name: xml.Name{"", key}}
	switch val := (*value).(type) {
	case string:
		*tokens = append(*tokens, t, xml.CharData(val), t.End())
	case bool:
		*tokens = append(*tokens, t, xml.CharData(strconv.FormatBool(val)), t.End())
	case float64:
		*tokens = append(*tokens, t, xml.CharData(strconv.FormatFloat(val, 'f', -1, 64)), t.End())
	case int:
		*tokens = append(*tokens, t, xml.CharData(strconv.Itoa(val)), t.End())
	case uint64:
		*tokens = append(*tokens, t, xml.CharData(strconv.FormatUint(val, 10)), t.End())

	case map[interface{}]interface{}:
		*tokens = append(*tokens, t)
		for k, v := range val {
			this.switchValue(tokens, k.(string), &v)
		}

		*tokens = append(*tokens, t.End())

	case []interface{}:
		for _, v := range val {
			this.switchValue(tokens, key, &v)
		}

	default:
		fmt.Println("Unknown marshal switch type:", reflect.TypeOf(val))
	}

}
