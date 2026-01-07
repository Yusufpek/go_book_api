package api

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Status int 	   `json:"status"`
	Message string   `json:"message"`
	Data   any      `json:"data,omitempty"`
}

func ResponseJson(c *gin.Context, status int, message string, data any) {
	var response = JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, response)
}