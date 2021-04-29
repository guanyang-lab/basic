/**
* @Auther:gy
* @Date:2021/3/6 13:36
 */

package models

import (
	"gorm.io/gorm"
)

var tdb *gorm.DB
var err error

func GetDB() *gorm.DB {
	return tdb
}
