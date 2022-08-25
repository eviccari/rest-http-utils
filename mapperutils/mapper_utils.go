package mapperutils

import (
	"encoding/json"
	"fmt"
)

// StructToMap - To convert a generic golang struct into map[string]interface{}
func StructToMap(st interface{}) (map[string]interface{}, error) {
	var mapped *map[string]interface{}
	marshalled, err := json.Marshal(st)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(marshalled, mapped)
	if err != nil {
		return nil, err
	}

	return *mapped, nil
}

// StructToJSONString - convert the given struct to JSON string
func StructToJSONString(stc interface{}) string {
	b, err := json.Marshal(stc)
	if err != nil {
		fmt.Println(fmt.Errorf("error on stringfy struct :%v", stc))
		return ""
	}

	return string(b)
}
