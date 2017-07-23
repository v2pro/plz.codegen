package helper

import "encoding/json"

func ToJson(obj interface{}) string {
	output, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(output)
}

func FromJson(obj interface{}, encodedAsJson string) {
	err := json.Unmarshal([]byte(encodedAsJson), obj)
	if err != nil {
		panic(err.Error())
	}
}