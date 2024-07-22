package system

import (
	"time"
)

type SysCustom struct {
	CreatedAt    time.Time  // 创建时间
	UpdatedAt    time.Time  // 更新时间
	DeletedAt    *time.Time `sql:"index"`
	CustomId     string     `gorm:"column:customId" json:"customId"`         //自定义布局Id
	CustomName   string     `gorm:"column:customName" json:"customName"`     //自定义布局名称
	CustomLayout string     `gorm:"column:customLayout" json:"customLayout"` //自定义布局数组
	PermissionId string     `gorm:"column:permissionId" json:"permissionId"` //权限Id
}

type CustomLayout struct {
	ConstainDetail string `gorm:"column:constainDetail" json:"constainDetail"`
	ContainId      string `gorm:"column:containId" json:"containId"`
	CoverType      string `gorm:"column:coverType" json:"coverType"`
	Draggable      string `gorm:"column:draggable" json:"draggable"`
	H              string `gorm:"column:h" json:"h"`
	I              string `gorm:"column:i" json:"i"`
	IsResizable    string `gorm:"column:isResizable" json:"isResizable"`
	LayoutGroup    string `gorm:"column:layoutGroup" json:"layoutGroup"`
	LayoutStatic   string `gorm:"column:layoutStatic" json:"layoutStatic"`
	Moved          string `gorm:"column:moved" json:"moved"`
	Text           string `gorm:"column:text" json:"text"`
	W              string `gorm:"column:w" json:"w"`
	X              string `gorm:"column:x" json:"x"`
	Y              string `gorm:"column:y" json:"y"`
	CustomId       string `gorm:"column:customId" json:"customId"`
}
