package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/xlt/shop_web/user_web/forms"
	"github.com/xlt/shop_web/user_web/global"
	"github.com/xlt/shop_web/user_web/global/response"
	"github.com/xlt/shop_web/user_web/middleware"
	"github.com/xlt/shop_web/user_web/model"
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

	pn := ctx.DefaultQuery("pn", "0")
	psize := ctx.DefaultQuery("psize", "2")
	pnInt, _ := strconv.Atoi(pn)
	psizeInt, _ := strconv.Atoi(psize)
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(psizeInt),
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

func PasswordLogin(ctx *gin.Context) {
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	// 第三个参数为true表示验证码验证过后就会删除
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		global.ServerConfig.UserSrvInfo.Host,
		global.ServerConfig.UserSrvInfo.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
		return
	}

	userClient := proto.NewUserClient(conn)
	user, err := userClient.GetUserByMobile(context.Background(),
		&proto.MobileRequest{Mobile: passwordLoginForm.Mobile},
	)
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	result, err := userClient.CheckPassWord(context.Background(), &proto.PassWordCheckInfo{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: user.PassWord,
	})
	if err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	if !result.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	j := middleware.NewJWT()
	claims := model.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(global.ServerConfig.JwtInfo.ExpireSecond*global.ServerConfig.JwtInfo.ExpireCount),
			Issuer:    "xiaolatiao",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.S().Errorw("j.CreateToken failed", "msg", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        user.Id,
		"nick_name": user.NickName,
		"msg":       "登录成功",
		"token":     token,
	})
}

func Register(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	userClient := proto.NewUserClient(userConn)
	user, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		zap.S().Errorw("userClient.CreateUser failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	j := middleware.NewJWT()
	claims := model.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(global.ServerConfig.JwtInfo.ExpireSecond*global.ServerConfig.JwtInfo.ExpireCount),
			Issuer:    "xiaolatiao",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func HandleValidatorError(ctx *gin.Context, err error) {
	e, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(e.Translate(global.Trans)),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": err.Error(),
	})
	return
}

func removeTopStruct(fileds map[string]string) map[string]string {
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
