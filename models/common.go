package models

import (
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

//分页
type Page struct {
	Offset int `json:"offset" comment:"跳过多少条 (当前页-1)*limit"`
	Limit  int `json:"limit" comment:"每页多少条"`
}
type Preload struct {
	Column     string
	Conditions []interface{}
}

func Create(obj interface{}, db *gorm.DB) (err error) {
	if db == nil {
		db = GetDB()
	}
	return db.Create(obj).Error
}
func CreateMap(obj interface{}, db *gorm.DB, m map[string]interface{}) (err error) {
	if db == nil {
		db = GetDB()
	}
	return db.Model(obj).Create(m).Error
}

// order: age desc, name
func GetByCondition(db *gorm.DB, Column []string, Value []interface{}, Preload []Preload, Order string, record interface{}) (err error) {
	if db == nil {
		db = GetDB()
	}
	db = SelectByCondition(db, Column, Value, Preload, Order)
	if reflect.ValueOf(record).Elem().Kind() == reflect.Slice {
		err = db.Find(record).Error
	} else {
		err = db.First(record).Error
		if err == gorm.ErrRecordNotFound {
			return nil
		}
	}
	return err
}
func GetByConditionOne(db *gorm.DB, Column []string, Value []interface{}, Preload []Preload, record interface{}) (err error) {
	if db == nil {
		db = GetDB()
	}
	db = SelectByCondition(db, Column, Value, Preload, "")
	err = db.First(record).Error
	return err
}
func GetByConditionPluck(obj interface{}, db *gorm.DB, Column []string, Value []interface{}, Col string, record interface{}) (err error) {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(obj)
	db = SelectByCondition(db, Column, Value, nil, "")
	err = db.Pluck(Col, record).Error
	return err
}
func GetCount(obj interface{}, Column []string, Value []interface{}) (res int, err error) {
	db := GetDB().Model(obj)
	var total int64
	db = SelectByCondition(db, Column, Value, nil, "")
	err = db.Count(&total).Error
	res = cast.ToInt(total)
	return
}
func UpdateByCondition(obj interface{}, db *gorm.DB, Column []string, Value []interface{}, m map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(obj)
	db = SelectByCondition(db, Column, Value, nil, "")
	return db.Updates(m).Error
}
func UpdateStructByCondition(obj interface{}, db *gorm.DB, Column []string, Value []interface{}, data interface{}) error {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(obj)
	db = SelectByCondition(db, Column, Value, nil, "")
	return db.Updates(data).Error
}
func DeleteByCondition(obj interface{}, db *gorm.DB, Column []string, Value []interface{}) error {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(obj)
	db = SelectByCondition(db, Column, Value, nil, "")
	return db.Delete(obj).Error
}

// FindPage 查询分页数据
func FindPage(db *gorm.DB, Offset, Limit int, out interface{}) (count int64, err error) {
	// 如果分页大小小于0，查询全部数据
	if Limit <= 0 {
		err = db.Find(out).Error
		return
	}
	err = db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	}
	err = db.Offset(Offset).Limit(Limit).Find(out).Error
	if err != nil {
		return 0, err
	}
	return
}
func BatchInsert(db *gorm.DB, TableName, Column string, Value []interface{}) error {
	if db == nil {
		db = GetDB()
	}
	sql := "insert into " + TableName + " (" + Column + ") Values "
	str := strings.Split(Column, ",")
	l := len(str)
	vl := len(Value)
	for il := 0; il < vl/l; il++ {
		sql += "("
		for i := 0; i < l; i++ {
			if i != l-1 {
				sql += "?,"
			} else {
				sql += "?"
			}
		}
		if il != (vl/l)-1 {
			sql += "),"
		} else {
			sql += ")"
		}
	}
	return db.Exec(sql, Value...).Error
}

func FindPagePreload(db *gorm.DB, Offset, Limit int, Preload []Preload, out interface{}) (count int64, err error) {
	for _, v := range Preload {
		if v.Conditions == nil {
			db = db.Preload(v.Column)
		} else {
			fmt.Println(v.Column)
			fmt.Println(v.Conditions)
			db = db.Preload(v.Column, v.Conditions...)
		}
	}
	// 如果分页大小小于0，查询全部数据
	if Limit <= 0 {
		err = db.Find(out).Error
		return
	}
	err = db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	}
	err = db.Offset(Offset).Limit(Limit).Find(out).Error
	if err != nil {
		return 0, err
	}
	return
}

func SelectByCondition(db *gorm.DB, Column []string, Value []interface{}, Preload []Preload, Order string) *gorm.DB {
	for i, v := range Value {
		if v == nil {
			db = db.Where(Column[i])
		} else {
			count := strings.Count(Column[i], "?")
			if count == 1 {
				db = db.Where(Column[i], v)
			} else {
				elem := reflect.ValueOf(v)
				if elem.Type().Kind() == reflect.Slice {
					data := []interface{}{}
					for i := 0; i < elem.Len(); i++ {
						ele := elem.Index(i)
						data = append(data, ele.Interface())
					}
					db = db.Where(Column[i], data...)
				}
			}
		}
	}
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
	return db
}
