package util

import (
	"encoding/json"

	"github.com/boj/redistore"
	"github.com/gomodule/redigo/redis"
	"github.com/ucho456job/my_authn_authz/internal/dto"

	"github.com/labstack/echo/v4"
)

type SessionManager interface {
	CacheLoginSession(c echo.Context, userID string) error
	LoadLoginSession(c echo.Context) (userID string, isLogin bool, err error)
	CachePreAuthnSession(c echo.Context, q dto.AuthorizationQuery) error
	CacheAuthzCodeWithCtx(c echo.Context, q dto.AuthorizationQuery, authzCode, userID string) error
	CheckHealthForRedis(c echo.Context, key string) (string, error)
}

type sessionManager struct {
	store *redistore.RediStore
	pool  *redis.Pool
}

func NewSessionManager(store *redistore.RediStore, pool *redis.Pool) SessionManager {
	return &sessionManager{
		store: store,
		pool:  pool,
	}
}

func (sm *sessionManager) CacheLoginSession(c echo.Context, userID string) error {
	sess, err := sm.store.Get(c.Request(), LOGIN_SESSION_NAME)
	if err != nil {
		return err
	}

	sess.Values["userId"] = userID
	sess.Values["isLogin"] = true

	return sess.Save(c.Request(), c.Response())
}

func (sm *sessionManager) LoadLoginSession(c echo.Context) (userID string, isLogin bool, err error) {
	sess, err := sm.store.Get(c.Request(), LOGIN_SESSION_NAME)
	if err != nil {
		return "", false, err
	}

	val, ok := sess.Values["userId"]
	if !ok {
		return "", false, nil
	}

	userID, ok = val.(string)
	if !ok {
		return "", false, nil
	}

	isLoginVal, ok := sess.Values["isLogin"]
	if !ok {
		return userID, false, nil
	}

	isLogin, ok = isLoginVal.(bool)
	if !ok {
		return userID, false, nil
	}

	return userID, isLogin, nil
}

func (sm *sessionManager) CachePreAuthnSession(c echo.Context, q dto.AuthorizationQuery) error {
	sess, err := sm.store.Get(c.Request(), PRE_AUTHN_SESSION_NAME)
	if err != nil {
		return err
	}

	sess.Values["responseType"] = q.ResponseType
	sess.Values["clientId"] = q.ClientID
	sess.Values["redirectUri"] = q.RedirectURI
	sess.Values["scope"] = q.Scope
	sess.Values["state"] = q.State
	sess.Values["nonce"] = q.Nonce

	return sess.Save(c.Request(), c.Response())
}

type AuthzCodeContext struct {
	UserID      string `json:"userId"`
	ClientID    string `json:"clientId"`
	Scope       string `json:"scope"`
	RedirectURI string `json:"redirectUri"`
	Nonce       string `json:"nonce"`
}

func (sm *sessionManager) CacheAuthzCodeWithCtx(c echo.Context, q dto.AuthorizationQuery, authzCode, userID string) error {
	saveData := AuthzCodeContext{
		UserID:      userID,
		ClientID:    q.ClientID,
		Scope:       q.Scope,
		RedirectURI: q.RedirectURI,
		Nonce:       q.Nonce,
	}

	encodedData, err := json.Marshal(saveData)
	if err != nil {
		return err
	}

	conn := sm.pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", authzCode, encodedData, "EX", 60)
	if err != nil {
		return err
	}

	return nil
}

func (sm *sessionManager) CheckHealthForRedis(c echo.Context, key string) (value string, err error) {
	conn := sm.pool.Get()
	defer conn.Close()

	value, err = redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return value, nil
}