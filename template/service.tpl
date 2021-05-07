
package service

import (
	"errors"
	"{{.AppName}}/models"
	"{{.AppName}}/pkg/logging"
)
type {{.StructName}}Req struct {
    ID   uint   `json:"id" comment:"创建无需填写  修改必填"`
}
func {{.StructName}}Create(req {{.StructName}}Req) (err error) {
	{{.StructLowerName}}s, err := models.{{.StructName}}Find([]string{}, []interface{}{}, nil, "")
	if err != nil {
		logging.Errorf("{{.StructName}}Create {{.StructName}}Find err:%s", err.Error())
		return errors.New("")
	}
	if len({{.StructLowerName}}s) > 0 {
		return errors.New("")
	}
	{{.StructLowerName}} := &models.{{.StructName}}{

	}
	err = models.{{.StructName}}Create({{.StructLowerName}})
	if err != nil {
		logging.Errorf("{{.StructName}}Create err:%s", err.Error())
		return errors.New("")
	}
	return
}

func {{.StructName}}Update(req {{.StructName}}Req) (err error) {
	{{.StructLowerName}}s, err := models.{{.StructName}}Find([]string{}, []interface{}{}, nil, "")
	if err != nil {
		logging.Errorf("{{.StructName}}Create {{.StructName}}Find err:%s", err.Error())
		return errors.New("")
	}
	if len({{.StructLowerName}}s) > 0 && {{.StructLowerName}}s[0].ID != req.ID {
		return errors.New("")
	}
	m := map[string]interface{}{

	}
	err = models.{{.StructName}}Update(req.ID, m)
	if err != nil {
		logging.Errorf("{{.StructName}}Update err:%s", err.Error())
		return errors.New("")
	}
	return
}

type {{.StructName}}UpdateStatusReq struct {
	ID       uint  `json:"id" comment:"数据id" validate:"required"`
	Type     uint  `json:"type" comment:"1变更状态"`
}

func {{.StructName}}UpdateStatus(req {{.StructName}}UpdateStatusReq) (err error) {
	m := map[string]interface{}{}
	switch req.Type {
	case 1:

	}
	err = models.{{.StructName}}Update(req.ID, m)
	if err != nil {
		logging.Errorf("{{.StructName}}UpdateStatus err:%s", err.Error())
		return errors.New("")
	}
	return
}

func {{.StructName}}Delete(ids []uint) (err error) {
	err = models.{{.StructName}}Delete(ids)
	if err != nil {
		logging.Errorf("{{.StructName}}Delete err:%s", err.Error())
		return errors.New("")
	}
	return
}

type {{.StructName}}PageResp struct {
    ID        uint   `json:"id" comment:"数据id"`
}

func {{.StructName}}Page(req models.{{.StructName}}PageReq) ([]{{.StructName}}PageResp, int64, error) {
	var (
		datas []{{.StructName}}PageResp
		count int64
		err   error
		{{.StructLowerName}}s  []models.{{.StructName}}
	)
	{{.StructLowerName}}s, count, err = models.{{.StructName}}Page(req)
	if err != nil {
		logging.Errorf("{{.StructName}}Page err:%s", err.Error())
		return nil, 0, errors.New("")
	}
	for _,v := range {{.StructLowerName}}s {
		datas = append(datas, {{.StructName}}PageResp{
            ID: v.ID,
		})
	}
	return datas, count, nil
}

type {{.StructName}}DetailResp struct {
    ID        uint   `json:"id" comment:"数据id"`
}
type {{.StructName}}DetailReq struct{
    ID uint `json:"id" comment:"数据id" validate:"required,gt=0"`
}
func {{.StructName}}Detail(id uint) ({{.StructName}}DetailResp, error) {
	var data {{.StructName}}DetailResp
	var {{.StructLowerName}} models.{{.StructName}}
	err := models.GetByCondition(nil, []string{"id=?"}, []interface{}{id}, nil, "", &{{.StructLowerName}})
	if err != nil {
		logging.Errorf("{{.StructName}}Detail GetByCondition err:%s", err.Error())
		return data, errors.New("")
	}
	return {{.StructName}}DetailResp{
		ID:           {{.StructLowerName}}.ID,
	}, nil
}

type {{.StructName}}ListResp struct {
	ID uint `json:"id" comment:"数据id"`
}
type {{.StructName}}ListReq struct {
}

func {{.StructName}}List(req {{.StructName}}ListReq) ([]{{.StructName}}ListResp, error) {
	var (
		datas    []{{.StructName}}ListResp
		{{.StructLowerName}}s []models.{{.StructName}}
		err     error
	)
	{{.StructLowerName}}s, err = models.{{.StructName}}Find([]string{}, []interface{}{}, nil, "")
	if err != nil {
		logging.Errorf("{{.StructName}}List {{.StructName}}Find err:%s", err.Error())
		return datas, errors.New("")
	}
	for _,v := range {{.StructLowerName}}s {
		datas = append(datas, {{.StructName}}ListResp{
			ID: v.ID,
		})
	}
	return datas, nil
}
