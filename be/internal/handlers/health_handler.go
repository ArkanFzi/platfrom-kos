package handlers

import (
	"koskosan-be/internal/config"
	"koskosan-be/internal/database"
	"koskosan-be/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	cfg *config.Config
}

func NewHealthHandler(cfg *config.Config) *HealthHandler {
	return &HealthHandler{cfg: cfg}
}

type HealthResponse struct {
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	Database    string `json:"database"`
	Version     string `json:"version"`
	Uptime      int64  `json:"uptime"`
}

var startTime = time.Now()

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	db := database.GetDB()

	// Check database connection
	dbStatus := "healthy"
	sqlDB, err := db.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "unhealthy"
		utils.GlobalLogger.Error("Database health check failed: %v", err)
	}

	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
		Database:  dbStatus,
		Version:   h.cfg.AppVersion,
		Uptime:    int64(time.Since(startTime).Seconds()),
	}

	if dbStatus == "unhealthy" {
		utils.ErrorResponse(c, 503, "Service unavailable", "Database connection failed")
		return
	}

	utils.OKSuccess(c, response, "Health check passed")
}
