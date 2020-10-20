package error

import (
	"github.com/gin-gonic/gin"
)

type ErrResponseInterceptor interface {
	HandleBadRequest(ctx *gin.Context, bindErr error)
	HandleServiceError(ctx *gin.Context,  serviceErr *APPError)
}
type errResponseInterceptor struct {
}

func NewErrResponseInterceptor() ErrResponseInterceptor {
	return errResponseInterceptor{}
}

func (e errResponseInterceptor) HandleBadRequest(ctx *gin.Context, bindErr error) {
	bdRequestError := BadRequestErrorFunc(bindErr.Error())
	ctx.AbortWithStatusJSON(bdRequestError.StatusCode, bdRequestError.ErrResponse)
}
func (e errResponseInterceptor) HandleServiceError(ctx *gin.Context, serviceErr *APPError) {
	if serviceErr.ErrType == InternalServerErrorType {
		ctx.AbortWithStatus(serviceErr.StatusCode)
	}
	ctx.AbortWithStatusJSON(serviceErr.StatusCode, serviceErr.ErrResponse)
}
