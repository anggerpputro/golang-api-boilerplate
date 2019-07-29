package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct{}

func NewJsonResponse() *JsonResponse {
	return &JsonResponse{}
}

func response(gc *gin.Context, code int, data map[string]interface{}) {
	responseData := gin.H{
		"statusCode": code,
		"statusText": http.StatusText(code),
		"data":       data,
	}

	gc.JSON(code, responseData)
}

func (jr JsonResponse) ResponseSuccess(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusOK, data)
}

func (jr JsonResponse) ResponseBadRequest(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusBadRequest, data)
}

func (jr JsonResponse) ResponseUnauthorized(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusUnauthorized, data)
}

func (jr JsonResponse) ResponseForbidden(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusForbidden, data)
}

func (jr JsonResponse) ResponseNotFound(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusNotFound, data)
}

func (jr JsonResponse) ResponseError(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusInternalServerError, data)
}
