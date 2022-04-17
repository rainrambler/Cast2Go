package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type JsonContent struct {
	Data interface{}
}

func parseFile(jsonFile string) (*JsonContent, error) {
	raw, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	var data = new(interface{})
	err = json.Unmarshal(raw, data)
	if err != nil {
		return nil, err
	}

	return &JsonContent{*data}, nil
}

// https://github.com/ChimeraCoder/gojson
func iterContent(content interface{}) {
	switch content.(type) {
	case []interface{}:
		val := content.([]interface{})
		fmt.Printf("DBG: [%d] item\n", len(val))
		for _, child := range val {
			iterContent(child)
		}
	case map[interface{}]interface{}:
		val := content.(map[interface{}]interface{})
		for k, v := range val {
			fmt.Printf("Key: [%v]\n", k)
			iterContent(v)
		}
	case map[string]interface{}:
		val := content.(map[string]interface{})
		for k, v := range val {
			fmt.Printf("Key: [%v]\n", k)
			iterContent(v)
		}
	case string:
		fmt.Printf("sTr:[%s]\n", content.(string))
	case int:
		fmt.Printf("iNt:[%s]\n", content.(string))
	case float64:
		fmt.Printf("fLt:[%v]\n", content.(float64))
	default:
		// https://stackoverflow.com/questions/6372474/how-to-determine-an-interface-values-real-type
		otherType := reflect.TypeOf(content)
		fmt.Printf("Unknown: %+v\n", otherType)
	}
}

// https://github.com/ChimeraCoder/gojson
func iterContentLayer(content interface{}, layers int) {
	if layers <= 0 {
		return
	}
	otherType := reflect.TypeOf(content)
	fmt.Printf("DBG: Type: %v\n", otherType)

	switch content.(type) {

	case []interface{}:
		val := content.([]interface{})
		//fmt.Printf("DBG: [%d] item\n", len(val))
		fmt.Println("[")
		for _, child := range val {
			iterContentLayer(child, layers-1)
		}
		fmt.Println("]")

	case map[interface{}]interface{}:
		val := content.(map[interface{}]interface{})
		fmt.Println("{")
		for k, v := range val {
			fmt.Printf("Key: [%v]\n", k)
			iterContentLayer(v, layers-1)
		}
		fmt.Println("}")

	case map[string]interface{}:
		val := content.(map[string]interface{})
		fmt.Println("{")
		for k, v := range val {
			fmt.Printf("Key: [%v]\n", k)
			iterContentLayer(v, layers-1)
		}
		fmt.Println("}")

	case string:
		fmt.Printf("sTr:[%s]\n", content.(string))
	case int:
		fmt.Printf("iNt:[%s]\n", content.(string))
	case float64:
		fmt.Printf("fLt:[%v]\n", content.(float64))
	default:
		// https://stackoverflow.com/questions/6372474/how-to-determine-an-interface-values-real-type
		otherType := reflect.TypeOf(content)
		fmt.Printf("Unknown: %+v\n", otherType)
	}
}
