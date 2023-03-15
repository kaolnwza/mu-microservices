package helper

import (
	"encoding/json"
)

func StructCopy(base interface{}, dest interface{}) error {
	baseJson, err := json.Marshal(base)
	if err != nil {
		return err
	}

	return json.Unmarshal(baseJson, &dest)
}
