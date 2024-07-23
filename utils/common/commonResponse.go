package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	AdditionalInfo
}

type AdditionalInfo struct {
	TotalData int64  `json:"total_data"`
	TraceId   string `json:"trace_id"`
}

func GenerateSuccessResponse(ctx *gin.Context, message string) {
	ctx.JSON(
		http.StatusOK,
		GenerateSuccessMessage(message),
	)
}

func GenerateSuccessResponseWithData(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		GenerateSuccessMessageWithData(message, data),
	)
}

func GenerateSuccessResponseWithListData(ctx *gin.Context, message string, total int64, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		GenerateSuccessMessageWithListData(message, total, data),
	)
}

func GenerateErrorResponse(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		GenerateErrorMessage(ctx, message),
	)
}

func GenerateSuccessMessage(message string) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    nil,
	}
}

func GenerateSuccessMessageWithData(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func GenerateSuccessMessageWithListData(message string, total int64, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		AdditionalInfo: AdditionalInfo{
			TotalData: total,
		},
	}
}

func GenerateErrorMessage(ctx *gin.Context, message string) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
		AdditionalInfo: AdditionalInfo{
			TraceId: ctx.GetString("trace_id"),
		},
	}
}
