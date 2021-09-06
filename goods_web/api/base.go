package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/xlt/shop_web/goods_web/global"
)

func HandleValidatorError(ctx *gin.Context, err error) {
	e, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": RemoveTopStruct(e.Translate(global.Trans)),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": err.Error(),
	})
	return
}

func RemoveTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(ctx *gin.Context, err error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "服务不可用",
				})
			default:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "其他错误",
				})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "未知错误",
			})
		}
	}
}
