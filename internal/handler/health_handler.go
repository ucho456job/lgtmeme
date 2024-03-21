package handler

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/labstack/echo/v4"
	"github.com/ucho456job/my_authn_authz/internal/repository"
	"github.com/ucho456job/my_authn_authz/internal/util"
)

type HealthHandler interface {
	CheckHealth(c echo.Context) error
}

type healthHandler struct {
	healthCheckRepository repository.HealthCheckRepository
	sessionManager        util.SessionManager
	logger                *slog.Logger
}

func NewHealthHandler(healthCheckRepository repository.HealthCheckRepository, sessionManager util.SessionManager, logger *slog.Logger) *healthHandler {
	return &healthHandler{
		healthCheckRepository: healthCheckRepository,
		sessionManager:        sessionManager,
		logger:                logger,
	}
}

func (h *healthHandler) CheckHealth(c echo.Context) error {
	key := "healthCheckKey"

	postgresValue, err := h.healthCheckRepository.CheckHealthForPostgres(c, key)
	if err != nil {
		h.logger.Error("Failed to check helath for postgres", "error", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid username or password"})
	}

	redisValue, err := h.sessionManager.CheckHealthForRedis(c, key)
	if err != nil {
		h.logger.Error("Failed to check health for redis", "error", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to check health for redis"})
	}

	h.logger.Info("Server, Postgres and Redis are healthy", "postgresValue", postgresValue, "redisValue", redisValue)
	return c.String(http.StatusOK, "Server is healthy!")
}
