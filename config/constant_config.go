package config

import "time"

const HEALTH_ENDPOINT = "/api/health"

// auth server
const (
	// api endpoint
	AUTHZ_ENDPOINT  = "/auth-api/authorize"
	JWKS_ENDPOINT   = "/auth-api/jwks"
	LOGIN_ENDPOINT  = "/auth-api/login"
	LOGOUT_ENDPOINT = "/auth-api/logout"
	TOKEN_ENDPOINT  = "/auth-api/token"

	// view endpoint
	LOGIN_VIEW_ENDPOINT = "/login"

	// file path
	LOGIN_VIEW_FILEPATH = "view/out/login.html"
)

// client server
const (
	// api endpoint
	CLIENT_ADMIN_ENDPOINT          = "/client-api/admin"
	CLIENT_ADMIN_CALLBACK_ENDPOINT = "/client-api/admin/callback"
	CLIENT_IMAGES_ENDPOINT         = "/client-api/images"

	// view endpoint
	STATIC_ENDPOINT           = "/"
	HOME_VIEW_ENDPOINT        = "/"
	IMAGE_NEW_VIEW_ENDPOINT   = "/new"
	ADMIN_VIEW_ENDPOINT       = "/admin"
	PRIVACY_POLICY_ENDPOINT   = "/privacy-policy"
	TERMS_OF_SERVICE_ENDPOINT = "/terms-of-service"

	// file path
	STATIC_FILEPATH           = "view/out"
	HOME_VIEW_FILEPATH        = "view/out/index.html"
	IMAGE_NEW_VIEW_FILEPATH   = "view/out/new.html"
	ADMIN_VIEW_FILEPATH       = "view/out/admin.html"
	PRIVACY_POLICY_FILEPATH   = "view/out/privacy-policy.html"
	TERMS_OF_SERVICE_FILEPATH = "view/out/terms-of-service.html"
)

// resoruce server
const (
	// api endpoint
	RESOURCE_IMAGES_ENDPOINT = "/resource-api/images"
)

// session name
const (
	LOGIN_SESSION_NAME                = "login_session"
	PRE_AUTHN_SESSION_NAME            = "pre_authn_session"
	GENERAL_ACCESS_TOKEN_SESSION_NAME = "general_access_token"
	ADMIN_ACCESS_TOKEN_SESSION_NAME   = "admin_access_token"
	STATE_AND_NONCE_SESSION_NAME      = "state_and_nonce"
	REFRESH_TOKEN_SESSION_NAME        = "refresh_token"
)

// expire
const (
	DEFAULT_SESSION_EXPIRE_SEC       = 60 * 60 * 23
	AUTHZ_CODE_EXPIRE_SEC            = 60
	REFRESH_TOKEN_SESSION_EXPIRE_SEC = 60 * 60 * 24 * 30

	ACCESS_TOKEN_EXPIRES_IN = time.Hour * 24
	ID_TOKEN_EXPIRES_IN     = time.Minute * 10
)

// number
const (
	REFRESH_TOKEN_SIZE = 64
	GET_IMAGES_LIMIT   = 9
)

// scope
const (
	IMAGES_READ_SCOPE   = "images.read"
	IMAGES_CREATE_SCOPE = "images.create"
	IMAGES_UPDATE_SCOPE = "images.update"
	IMAGES_DELETE_SCOPE = "images.delete"
)
