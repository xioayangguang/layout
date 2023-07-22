package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"layout/pkg/helper/snowflake"
	"time"
)

const (
	UpdateFailedError = "affected rows:0"
)

type Base struct {
	Id        uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key" json:"id"`
	CreatedAt uint   `gorm:"column:created_at;type:int(11) unsigned;default:0;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt uint   `gorm:"column:updated_at;type:int(11) unsigned;default:0;comment:更新时间;NOT NULL" json:"updated_at"`
	DeletedAt uint   `gorm:"column:deleted_at;type:int(11) unsigned;default:0;comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == 0 {
		u.Id = snowflake.GlobalSnowflake.Generate().UInt64()
	}
	tx.Statement.SetColumn("UpdateTime", time.Now().Unix())
	tx.Statement.SetColumn("CreateTime", time.Now().Unix())
	return
}

func (u *Base) BeforeSave(tx *gorm.DB) (err error) {
	u.UpdatedAt = uint(time.Now().Unix())
	return
}
func (u *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdateTime", time.Now().Unix())
	return
}
func (u *Base) CheckUpdateError(tx *gorm.DB) (err error) {
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(UpdateFailedError)
	}
	return nil
}

type BaseUpdateTimePo struct {
	UpdateTime int64 `gorm:"column:update_time"`
}

func (u *BaseUpdateTimePo) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdateTime", time.Now().Unix())
	return
}
