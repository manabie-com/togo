package utils

import "encoding/json"

func MarshalDto(data, res any) error {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshaled, res)
	if err != nil {
		return err
	}
	return nil
}
