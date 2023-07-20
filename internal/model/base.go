package model

import (
	"github.com/pkg/errors"
	"time"

	"gorm.io/gorm"

	"horse/pkg/snowflake"
	"horse/pkg/soft_delete"
)

const (
	UpdateFailedError = "affected rows:0"
)

type Model struct {
	ID        uint       `json:"id" gorm:"primarykey"`            // 主键ID
	CreatedAt time.Time  `json:"createdAt"`                       // 创建时间
	UpdatedAt time.Time  `json:"updatedAt"`                       // 更新时间
	DeletedAt *time.Time `json:"deletedAt" gorm:"index" json:"-"` // 删除时间
}

type Base struct {
	Id         uint64                `gorm:"column:id;type:bigint(20) unsigned;primary_key" json:"id"`
	CreateTime int64                 `gorm:"column:create_time;type:int(11) unsigned;default:0;NOT NULL" json:"create_time"`
	UpdateTime int64                 `gorm:"column:update_time;type:int(11) unsigned;default:0;NOT NULL" json:"update_time"`
	DeleteTime soft_delete.DeletedAt `gorm:"column:delete_time;type:int(11) unsigned;default:0;NOT NULL" json:"delete_time"`
}
type BaseDeleteTimePo struct {
	DeleteTime soft_delete.DeletedAt `gorm:"column:delete_time;type:int(11);DEFAULT:0;not null;"`
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
	u.UpdateTime = time.Now().Unix()
	return
}
func (u *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	// 如果任意字段有变更
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
