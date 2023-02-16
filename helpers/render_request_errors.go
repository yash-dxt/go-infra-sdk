package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestError struct {
	StatusCode int   `json:"statusCode"`
	Err        error `json:"error"`
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func renderRequestError(ctx gin.Context, reqErr RequestError) {
	ctx.JSON(reqErr.StatusCode, gin.H{"error": reqErr.Error()})
}

// Use this function for service related errors.
func HandleError(ctx gin.Context, err error) {
	requestErr, ok := err.(*RequestError)

	if ok {
		renderRequestError(ctx, *requestErr)
	} else {
		renderError(ctx, err)
	}

}

func renderError(ctx gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func RenderBadRequest(ctx gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "msg": message})
}

func RenderInternalServerError(ctx gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "msg": message})
}

func RenderUnauthorizedError(ctx gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
}
