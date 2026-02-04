package handlers

import (
	"koskosan-be/internal/database"
	"koskosan-be/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	Database    string `json:"database"`
	Version     string `json:"version"`
	Uptime      int64  `json:"uptime"`
}

var startTime = time.Now()

func HealthCheck(c *gin.Context) {
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
		Version:   "1.0.0", // TODO: Read dari config
		Uptime:    int64(time.Since(startTime).Seconds()),
	}

	if dbStatus == "unhealthy" {
		utils.ErrorResponse(c, 503, "Service unavailable", "Database connection failed")
		return
	}

	utils.OKSuccess(c, response, "Health check passed")
}
