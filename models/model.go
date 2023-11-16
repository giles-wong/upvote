package models

import "time"

type Model struct {
	PkId        int       `gorm:"column:pk_id;primaryKey" json:"pkId"`                                //自增主键
	GmtCreated  time.Time `gorm:"column:gmt_created;default:CURRENT_TIMESTAMP" json:"gmtCreated"`     //创建时间
	GmtModified time.Time `gorm:"column:gmt_modified;default:0000-00-00 00:00:00" json:"gmtModified"` //更新时间
}
