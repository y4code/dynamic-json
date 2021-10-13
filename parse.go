package dynamic_json

import (
	"encoding/json"
)

func Parse(raw string) map[string]string {
	var (
		dynamic = make(map[string]interface{})
		err     error
	)
	if err = json.Unmarshal([]byte(raw), &dynamic); err != nil {
		panic(err)
	}
	return Unmarshal(dynamic, "")
}

func Unmarshal(dynamic map[string]interface{}, parentKey string) map[string]string {
	var (
		KV = make(map[string]string)
	)
	if parentKey != "" {
		parentKey = parentKey + "."
	}
	for key, value := range dynamic {
		if valueMaybeObject, ok := value.(map[string]interface{}); ok {
			childKV := Unmarshal(valueMaybeObject, parentKey+key)
			for childKey, childValue := range childKV {
				KV[childKey] = childValue
			}
		} else {
			KV[parentKey+key] = value.(string)
		}
	}
	return KV
}
