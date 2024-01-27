package controller

import (
	"IM/internel/db"
	"IM/internel/db/ent"
	"IM/internel/redis"
	"IM/internel/types"
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// SendCode 发送验证码
func SendCode(ctx *gin.Context) {
	var req types.SendCodeReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(ctx, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	// 生成随机验证码
	randomCode := utils.GenerateRegisterCode(types.RegisterCodeLength)

	// 保存到 redis
	// TODO 校验并保存最后一个验证码
	rKey := fmt.Sprintf(types.RedisEmailCode, req.Action, req.Email)
	err := redis.Client.Set(ctx, rKey, randomCode, types.RedisEmailCodeExpireTime*time.Minute).Err()
	if err != nil {
		logrus.Errorf("failed to set email(%s) code(%s) to redis, err: %v", rKey, randomCode, err)
		response.RespError(ctx, code.ServerErrCache)
		return
	}

	// 发送验证码
	err = utils.SendMail(req.Email, randomCode)
	if err != nil {
		response.RespError(ctx, code.ServerErrThirdPartyAPI)
		return
	}

	response.RespSuccess(ctx, "发送验证码成功")
}

// Register 注册
func Register(ctx *gin.Context) {
	_, err := db.DB.User.Query().First(ctx)
	if ent.IsNotFound(err) {
		response.RespSuccess(ctx, "chenggong")
		return
	} else if err != nil {
		logrus.Error(err)
		response.RespErrorWithMsg(ctx, code.ServerErr, err.Error())
		return
	}

}
