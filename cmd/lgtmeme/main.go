package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ucho456job/lgtmeme/config"
	"github.com/ucho456job/lgtmeme/internal/handler"
	"github.com/ucho456job/lgtmeme/internal/middleware"
	"github.com/ucho456job/lgtmeme/internal/repository"
	"github.com/ucho456job/lgtmeme/internal/service"
	"github.com/ucho456job/lgtmeme/internal/util/clock"
	"github.com/ucho456job/lgtmeme/internal/util/uuidgen"
)

func main() {
	config.NewEnv()
	config.NewDB()
	config.NewSessionStore()
	config.NewLogger()
	validator := config.NewValidator()

	e := echo.New()
	e.Validator = validator
	e.Use(middleware.SessionMiddleware(), middleware.LoggerMiddleware)
	e.Static(config.STATIC_ENDPOINT, config.STATIC_FILEPATH)

	newAuthServer(e)
	newClientServer(e)
	newResourceServer(e)

	// Graceful shutdown
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	port := ":" + os.Getenv("PORT")
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func newAuthServer(e *echo.Echo) {
	healthRepo := repository.NewHealthRepository(config.DB)
	oauthClientRepo := repository.NewOauthClientRepository(config.DB)
	refreshTokenRepo := repository.NewRefreshTokenRepository(config.DB)
	authSessManaRepo := repository.NewSessionManagerRepository(config.Store, config.Pool)
	userRepo := repository.NewUserRepository(config.DB)

	jwtServ := service.NewJWTService()

	authzHandler := handler.NewAuthzHandler(oauthClientRepo, userRepo, authSessManaRepo)
	healthHandler := handler.NewHealthHandler(healthRepo, authSessManaRepo)
	jwksHandler := handler.NewJwksHandler(jwtServ)
	loginHandler := handler.NewLoginHandler(userRepo, authSessManaRepo)
	logoutHandler := handler.NewLogoutHandler(authSessManaRepo)
	tokenHandler := handler.NewTokenHandler(oauthClientRepo, refreshTokenRepo, userRepo, authSessManaRepo, jwtServ)

	e.GET(config.AUTHZ_ENDPOINT, authzHandler.Authorize)
	e.HEAD(config.HEALTH_ENDPOINT, healthHandler.Check)
	e.GET(config.JWKS_ENDPOINT, jwksHandler.Get)
	e.POST(config.LOGIN_ENDPOINT, loginHandler.Login)
	e.GET(config.LOGIN_VIEW_ENDPOINT, loginHandler.GetView)
	e.GET(config.LOGOUT_ENDPOINT, logoutHandler.Logout)
	e.POST(config.TOKEN_ENDPOINT, tokenHandler.Generate)
}

func newClientServer(e *echo.Echo) {
	sessManaRepo := repository.NewSessionManagerRepository(config.Store, config.Pool)

	imgServ := service.NewImageService()
	accessTokenServ := service.NewAccessTokenService()

	accessTokenMiddle := middleware.NewAccessTokenMiddleware(sessManaRepo, accessTokenServ)

	viewHandler := handler.NewViewHandler()
	e.GET(config.ADMIN_VIEW_ENDPOINT, viewHandler.GetAdminView, accessTokenMiddle.SetAdminAccessToken())
	e.GET(config.HOME_VIEW_ENDPOINT, viewHandler.GetHomeView, accessTokenMiddle.SetGeneralAccessToken())
	e.GET(config.IMAGE_NEW_VIEW_ENDPOINT, viewHandler.GetImageView, accessTokenMiddle.SetGeneralAccessToken())
	e.GET(config.ERROR_VIEW_ENDPOINT, viewHandler.GetErrView)
	e.GET(config.PRIVACY_POLICY_ENDPOINT, viewHandler.GetPrivacyPolicyView)
	e.GET(config.TERMS_OF_SERVICE_ENDPOINT, viewHandler.GetTermsOfServiceView)

	adminHandler := handler.NewAdminHandler(sessManaRepo, accessTokenServ)
	e.GET(config.CLIENT_ADMIN_CALLBACK_ENDPOINT, adminHandler.Callback)

	imgHandler := handler.NewClientImageHandler(sessManaRepo, accessTokenServ, imgServ)
	e.POST(config.CLIENT_IMAGES_ENDPOINT, imgHandler.Post)
	e.GET(config.CLIENT_IMAGES_ENDPOINT, imgHandler.BulkGet)
	e.PATCH(config.CLIENT_IMAGES_ENDPOINT+"/:image_id", imgHandler.Patch)
	e.DELETE(config.CLIENT_IMAGES_ENDPOINT+"/:image_id", imgHandler.Delete)
}

func newResourceServer(e *echo.Echo) {
	imgRepo := repository.NewImageRepository(config.DB, &clock.RealClocker{})
	sessManaRepo := repository.NewSessionManagerRepository(config.Store, config.Pool)
	userRepo := repository.NewUserRepository(config.DB)

	accessTokenServ := service.NewAccessTokenService()
	storageServ := service.NewStorageService()

	accessTokenMiddle := middleware.NewAccessTokenMiddleware(sessManaRepo, accessTokenServ)

	imgHandler := handler.NewResourceImageHandler(imgRepo, userRepo, storageServ, &uuidgen.RealUUIDGenerator{})

	e.POST(config.RESOURCE_IMAGES_ENDPOINT, imgHandler.Post, accessTokenMiddle.VerifyAccessToken(config.IMAGES_CREATE_SCOPE))
	e.GET(config.RESOURCE_IMAGES_ENDPOINT, imgHandler.BulkGet, accessTokenMiddle.VerifyAccessToken(config.IMAGES_READ_SCOPE))
	e.PATCH(config.RESOURCE_IMAGES_ENDPOINT+"/:image_id", imgHandler.Patch, accessTokenMiddle.VerifyAccessToken(config.IMAGES_UPDATE_SCOPE))
	e.DELETE(config.RESOURCE_IMAGES_ENDPOINT+"/:image_id", imgHandler.Delete, accessTokenMiddle.VerifyAccessToken(config.IMAGES_DELETE_SCOPE))
}
