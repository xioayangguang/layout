package model

import (
	"gorm.io/gorm"
	"layout/pkg/helper/snowflake"
	"time"
)

type Base struct {
	Id        uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key" json:"id"`
	CreatedAt uint   `gorm:"column:created_at;type:int(11) unsigned;default:0;comment:创建时间;NOT NULL" json:"created_at"` //
	UpdatedAt uint   `gorm:"column:updated_at;type:int(11) unsigned;default:0;comment:更新时间;NOT NULL" json:"updated_at"`
	DeletedAt uint   `gorm:"column:deleted_at;type:int(11) unsigned;default:0;comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == 0 {
		u.Id = snowflake.GlobalSnowflake.Generate().UInt64()
	}
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	return
}

func (u *Base) BeforeSave(tx *gorm.DB) (err error) {
	u.UpdatedAt = uint(time.Now().Unix())
	return
}
func (u *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return
}
