package models

//keycloak response
type TokenInfoResp struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        uint64 `json:"expires_in"`
	RefreshExpiresIn uint64 `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  uint64 `json:"not_before_policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}
