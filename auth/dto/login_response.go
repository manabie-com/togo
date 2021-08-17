package dto

import tokenprovider "github.com/manabie-com/togo/token_provider"

type LoginResponse struct {
	AccessToken tokenprovider.Token `json:"accessToken"`
}
