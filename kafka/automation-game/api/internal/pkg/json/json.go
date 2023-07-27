package json

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
