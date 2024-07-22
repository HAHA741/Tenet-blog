package system

import "time"

type Article struct {
	Id             int       `gorm:"column:article_id" json:"id"`
	ArticleTitle   string    `gorm:"column:article_title" json:"title"`
	ArticleContent string    `gorm:"column:article_content" json:"content"`
	Author         string    `gorm:"column:author" json:"author"`
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
	ArticleType    int       `gorm:"column:article_type" json:"articleType"`
}

type ArticleType struct {
	TypeId   int    `gorm:"column:type_id" json:"typeId"`
	TypeName string `gorm:"column:type_name" json:"typeName"`
}

type Page struct {
	Current int ` json:"current"`
	Size    int ` json:"size"`
}
