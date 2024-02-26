package api

import (
	"dst-admin-go/service"
	"dst-admin-go/utils/dstConfigUtils"
	"dst-admin-go/utils/fileUtils"
	"dst-admin-go/vo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InitApi struct {
}

const first = "./first"

func (f *InitApi) InitFirst(ctx *gin.Context) {

	exist := fileUtils.Exists(first)
	if exist {
		log.Panicln("非法请求")
	}

	initData := &service.InitDstData{}
	err := ctx.ShouldBind(initData)
	if err != nil {
		log.Panicln(err)
	}

	initEvnService.InitDstEnv(initData, ctx)

	fileUtils.CreateFile(first)
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}

func (f *InitApi) CheckIsFirst(ctx *gin.Context) {

	exist := fileUtils.Exists(first)

	code := 200
	msg := "is first"
	if exist {
		code = 400
		msg = "is not first"
	} else {
		dstConfig := dstConfigUtils.GetDstConfig()
		if dstConfig.Cluster != "" {
			initEvnService.InitBaseLevel(&dstConfig, "默认初始", "pds-g^KU_qE7e8rv1^VVrVXd/01kBDicd7UO5LeL+uYZH1+geZlrutzItvOaw=", true)
		}

	}

	ctx.JSON(http.StatusOK, vo.Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
