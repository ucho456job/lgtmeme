package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/ucho456job/lgtmeme/config"
	authDto "github.com/ucho456job/lgtmeme/internal/auth/dto"
)

type OwnerAccessTokenService interface {
	CallToken(c echo.Context) (respBody *authDto.AuthzCodeResp, status int, err error)
	CallJWKS(c echo.Context) (keySet jwk.Set, status int, err error)
	CallTokenWithRefreshToken(c echo.Context, refreshToken *string) (respBody *authDto.RefreshTokenResp, status int, err error)
}

type ownerAccessTokenService struct{}

func NewOwnerAccessTokenService() OwnerAccessTokenService {
	return &ownerAccessTokenService{}
}

func (s *ownerAccessTokenService) CallToken(c echo.Context) (respBody *authDto.AuthzCodeResp, status int, err error) {
	baseURL := os.Getenv("BASE_URL")
	url := baseURL + config.TOKEN_ENDPOINT
	clientID := os.Getenv("OWNER_CLIENT_ID")
	clientSecret := os.Getenv("OWNER_CLIENT_SECRET")
	redirectURI := os.Getenv("OWNER_REDIRECT_URI")
	code := c.QueryParam("code")

	reqData := fmt.Sprintf("grant_type=authorization_code&client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
		clientID, clientSecret, redirectURI, code)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(reqData))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.New("failed to get access token")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := json.Unmarshal(body, &respBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return respBody, resp.StatusCode, nil
}

func (s *ownerAccessTokenService) CallJWKS(c echo.Context) (keySet jwk.Set, status int, err error) {
	baseURL := os.Getenv("BASE_URL")
	url := baseURL + config.JWKS_ENDPOINT

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/jwk-set+json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.New("failed to get access token")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	keySet, err = jwk.Parse(body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return keySet, resp.StatusCode, nil
}

func (s *ownerAccessTokenService) CallTokenWithRefreshToken(c echo.Context, refreshToken *string) (respBody *authDto.RefreshTokenResp, status int, err error) {
	baseURL := os.Getenv("BASE_URL")
	url := baseURL + config.TOKEN_ENDPOINT
	clientID := os.Getenv("OWNER_CLIENT_ID")
	clientSecret := os.Getenv("OWNER_CLIENT_SECRET")

	reqData := fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s",
		clientID, clientSecret, *refreshToken)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(reqData))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.New("failed to get access token")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := json.Unmarshal(body, &respBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return respBody, resp.StatusCode, nil
}
