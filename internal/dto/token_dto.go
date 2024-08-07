package dto

import "github.com/google/uuid"

type TokenForm struct {
	GrantType    string    `form:"grant_type" validate:"required,grantType"`
	Code         string    `form:"code" validate:"required_with=GrantType=authorization_code"`
	RedirectURI  string    `form:"redirect_uri" validate:"required_with=GrantType=authorization_code"`
	ClientID     uuid.UUID `form:"client_id" validate:"required"`
	ClientSecret string    `form:"client_secret" validate:"required"`
	RefreshToken string    `form:"refresh_token" validate:"required_with=GrantType=refresh_token"`
}

type AuthzCodeResp struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
	IDToken      string `json:"idToken"`
}

type RefreshTokenResp struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

type ClientCredentialsResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
}
