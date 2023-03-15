package helper

import "encoding/json"

func StringToJSON(v interface{}, str string) error {
	return json.Unmarshal([]byte(str), &v)
}
