

package {{.StructLowerName}}

import (
	"gitee.com/yanggit123/tool"
	"github.com/gin-gonic/gin"
	"{{.AppName}}/models"
	"{{.AppName}}/service"
)

func {{.StructName}}Create(c *gin.Context) {
	var (
		rsp tool.Rsp
		req service.{{.StructName}}Req
		err error
	)
	err = tool.BindJSON(c, &req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	err = service.{{.StructName}}Create(req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	rsp.ReplySuccess(c, "")
	return
}
func {{.StructName}}Update(c *gin.Context) {
	var (
		rsp tool.Rsp
		req service.{{.StructName}}Req
		err error
	)
	err = tool.BindJSON(c, &req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	err = service.{{.StructName}}Update(req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	rsp.ReplySuccess(c, "")
	return
}
func {{.StructName}}UpdateStatus(c *gin.Context) {
	var (
		rsp tool.Rsp
		req service.{{.StructName}}UpdateStatusReq
		err error
	)
	err = tool.BindJSON(c, &req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	err = service.{{.StructName}}UpdateStatus(req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	rsp.ReplySuccess(c, "")
	return
}
func {{.StructName}}Delete(c *gin.Context) {
	var (
        rsp tool.Rsp
        req []uint
        err error
    )
    err = tool.BindJSONNotWithValidate(c, &req)
    if err != nil {
        rsp.ReplyFailOperation(c, err.Error())
        return
    }
    err = service.{{.StructName}}Delete(req)
    if err != nil {
        rsp.ReplyFailOperation(c, err.Error())
        return
    }
    rsp.ReplySuccess(c, "")
	return
}
func {{.StructName}}Page(c *gin.Context) {
	var (
		rsp tool.Rsp
		req models.{{.StructName}}PageReq
	)
	err := tool.BindJSON(c, &req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	data, count, err := service.{{.StructName}}Page(req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	rsp.ReplySuccess(c, tool.ListRsp{
		data,
		int(count),
	})
	return
}
func {{.StructName}}Detail(c *gin.Context) {
	var (
		rsp tool.Rsp
		req service.{{.StructName}}DetailReq
	)
	err := tool.BindJSON(c, &req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	data, err := service.{{.StructName}}Detail(req.ID)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	rsp.ReplySuccess(c, data)
	return
}
func {{.StructName}}List(c *gin.Context) {
	var (
		rsp tool.Rsp
		req service.{{.StructName}}ListReq
	)
	err := tool.BindJSON(c, &req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	data, err := service.{{.StructName}}List(req)
	if err != nil {
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	rsp.ReplySuccess(c, data)
	return
}
