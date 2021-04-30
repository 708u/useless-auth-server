package valueobject

type AccessToken struct {
	Value     string `json:"access_token"`
	TokenType string `json:"token_type"`
}
