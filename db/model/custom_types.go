package model

import "time"

type TlsRecord struct {
	FlowID    uint      `json:"flow_id"`
	UpdatedAt time.Time `json:"updated_at"`
	Ja3       string    `json:"ja3"`
	Ja3s      string    `json:"ja3s"`
}

type CategoryStats struct {
	Name  string `json:"name"`
	Total int32  `json:"total"`
}
