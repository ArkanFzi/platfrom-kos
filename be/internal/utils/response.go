package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StandardResponse adalah format response standard untuk semua endpoint
type StandardResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// SuccessResponse mengirim response sukses
func SuccessResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, StandardResponse{
		Success: true,
		Data:    data,
		Message: message,
		Errors:  []string{},
	})
}

// ErrorResponse mengirim response error
func ErrorResponse(c *gin.Context, statusCode int, message string, errors ...string) {
	if len(errors) == 0 {
		errors = []string{message}
	}
	c.JSON(statusCode, StandardResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// BadRequestError - 400
func BadRequestError(c *gin.Context, message string, errors ...string) {
	ErrorResponse(c, http.StatusBadRequest, message, errors...)
}

// UnauthorizedError - 401
func UnauthorizedError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}

// ForbiddenError - 403
func ForbiddenError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message)
}

// NotFoundError - 404
func NotFoundError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message)
}

// ConflictError - 409
func ConflictError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusConflict, message)
}

// InternalServerError - 500
func InternalServerError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}

// CreatedSuccess - 201
func CreatedSuccess(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusCreated, data, message)
}

// OKSuccess - 200
func OKSuccess(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusOK, data, message)
}

// NoContentSuccess - 204
func NoContentSuccess(c *gin.Context) {
	c.JSON(http.StatusNoContent, StandardResponse{
		Success: true,
		Message: "Success",
	})
}
