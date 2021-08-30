package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/xlt/shop_web/user_web/global"
	"github.com/xlt/shop_web/user_web/global/response"
	"github.com/xlt/shop_web/user_web/proto"
)

func GetUserList(ctx *gin.Context) {
	ip := global.ServerConfig.UserSrvInfo.Host
	port := global.ServerConfig.UserSrvInfo.Port

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
		return
	}

	userClient := proto.NewUserClient(conn)
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		zap.S().Errorw("userClient.GetUserList failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	userList := make([]response.User, len(rsp.Data))
	for index, value := range rsp.Data {
		user := response.User{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		userList[index] = user
		index++
	}

	ctx.JSON(http.StatusOK, userList)
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
			return
		}
	}
}
