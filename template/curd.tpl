package models

import(
	"gorm.io/gorm"
)

func {{.StructName}}Create({{.StructLowerName}} *{{.StructName}}) (err error) {
	return GetDB().Create(&{{.StructLowerName}}).Error
}
func {{.StructName}}CreateTx({{.StructLowerName}} *{{.StructName}}, tx *gorm.DB) (err error) {
	return tx.Create(&{{.StructLowerName}}).Error
}
func {{.StructName}}Update(id uint, m map[string]interface{}) (err error) {
	return GetDB().Model(&{{.StructName}}{}).Where("id=?", id).Updates(m).Error
}
func {{.StructName}}UpdateTx(id uint, m map[string]interface{}, tx *gorm.DB) (err error) {
	return tx.Model(&{{.StructName}}{}).Where("id=?", id).Updates(m).Error
}
func {{.StructName}}Delete(ids []uint) (err error) {
	return GetDB().Where("id in (?)", ids).Delete(&{{.StructName}}{}).Error
}
func {{.StructName}}DeleteTx(ids []uint, tx *gorm.DB) (err error) {
	return GetDB().Where("id in (?)", ids).Delete(&{{.StructName}}{}).Error
}
func {{.StructName}}Find(Column []string, Value []interface{}, Preload []Preload, Order string) ({{.StructLowerName}}s []{{.StructName}}, err error) {
	err = GetByCondition(GetDB(), Column, Value, Preload, Order, &{{.StructLowerName}}s)
	return
}
func {{.StructName}}FindByStruct({{.StructLowerName}} {{.StructName}}, Preload []Preload, Order string) ({{.StructLowerName}}s []{{.StructName}}, err error) {
	db := GetDB().Where({{.StructLowerName}})
	for _, v := range Preload {
		if v.Conditions == nil {
			db = db.Preload(v.Column)
		} else {
			db = db.Preload(v.Column, v.Conditions...)
		}
	}
	if Order != "" {
		db = db.Order(Order)
	} else {
		db = db.Order("id desc")
	}
	err = db.Find(&{{.StructLowerName}}s).Error
	return
}
func {{.StructName}}First(Column []string, Value []interface{}, Preload []Preload, Order string) ({{.StructLowerName}} {{.StructName}}, err error) {
	err = GetByCondition(GetDB(), Column, Value, Preload, Order, &{{.StructLowerName}})
	return
}
func {{.StructName}}Pluck(Column []string, Value []interface{}, Col string, record interface{}) (err error) {
	db := GetDB().Model(&{{.StructName}}{})
	db = SelectByCondition(db, Column, Value, nil, "")
	return db.Pluck(Col, record).Error
}

type {{.StructName}}PageReq struct {
	Page
}

func {{.StructName}}Page(req {{.StructName}}PageReq) ({{.StructLowerName}}s []{{.StructName}}, count int64, err error) {
	db := GetDB().Model(&{{.StructName}}{})
	count, err = FindPage(db, req.Offset, req.Limit, &{{.StructLowerName}}s)
	return
}
