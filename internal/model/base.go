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
	Id        uint64 `json:"id" gorm:"primarykey"`            // 主键ID
	CreatedAt int64  `json:"createdAt"`                       // 创建时间
	UpdatedAt int64  `json:"updatedAt"`                       // 更新时间
	DeletedAt int64  `json:"deletedAt" gorm:"index" json:"-"` // 删除时间
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
	u.UpdatedAt = time.Now().Unix()
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
