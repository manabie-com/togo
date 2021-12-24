package model

import (
	"encoding/base64"
	"encoding/json"
)

type Paging struct {
	Offset int `json:"offset"`
}

func TokenToPaging(token string) (*Paging, error) {
	paging := Paging{}

	if len(token) == 0 {
		return &paging, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(decoded, &paging)
	if err != nil {
		return nil, err
	}
	return &paging, nil
}

func PagingToToken(paging *Paging) string {
	bytes, _ := json.Marshal(paging)
	return base64.StdEncoding.EncodeToString(bytes)
}
