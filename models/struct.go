package models

import (
	"gitee.com/yanggit123/tool"
)

//管理后台-用户管理
type User struct {
	tool.Model2
	Name           string `gorm:"type:varchar(30);comment:'用户名称'" json:"name"`
	Account        string `gorm:"type:varchar(30);comment:'账号'" json:"account"`
	Password       string `gorm:"comment:'账号密码'" json:"password"`
	CreatorID      uint   `gorm:"comment:'创建人id'" json:"creator_id"`
	OrganizationID uint   `gorm:"comment:'机构id'" json:"organization_id"`
	IsAdmin        bool   `gorm:"default:false;comment:'是否管理员 true是'" json:"is_admin"`
}
type Article struct {
	tool.Model2
	Title          string `json:"title" gorm:"comment:'文章标题'"`
	Content        string `json:"content" gorm:"comment:'文章内容';type:mediumtext;"`
	OrganizationId uint   `json:"organization_id" gorm:"comment:'创建人机构id'"`
	PublicTime     string `json:"public_time" gorm:"comment:'发布时间';default:''"`
	Status         int8   `json:"status" gorm:"comment:'文章状态（1待审核2审核通过3审核不通过）';default:1"`
	IsPublic       uint8  `json:"is_public" gorm:"comment:'是否发布（0，未发布，1已发布）';default:0"`
	IsCommit       uint8  `json:"is_commit" gorm:"comment:'是否提交（0，未提交，1已提交）';default:1"`
	PublishTime    string `json:"publish_time" gorm:"comment:'发布时间';default:''"`
}
