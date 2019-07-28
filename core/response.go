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
	data["statusCode"] = code
	data["statusText"] = http.StatusText(code)

	gc.JSON(code, data)
}

func (jr JsonResponse) ResponseSuccess(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusOK, data)
}

func (jr JsonResponse) ResponseError(gc *gin.Context, data map[string]interface{}) {
	response(gc, http.StatusInternalServerError, data)
}
