package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"layout/pkg/contextValue"
	"time"
)

const (
	callBackBeforeName = "core:before"
	callBackAfterName  = "core:after"
	startTime          = "_start_time"
)

type TracePlugin struct{}

func (op *TracePlugin) Name() string {
	return "sqlPlugin"
}

func (op *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)
	// 结束后
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &TracePlugin{}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	return
}

func after(db *gorm.DB) {
	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}
	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}
	var sqlInfo contextValue.SQL
	sqlInfo.Timestamp = CSTLayoutString()
	sqlInfo.SQL = db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	sqlInfo.Stack = utils.FileWithLineNum()
	sqlInfo.Rows = db.Statement.RowsAffected
	sqlInfo.CostSeconds = time.Since(ts).Seconds()
	if ginCtx, ok := db.Statement.Context.(*gin.Context); ok {
		if sqls, exists := ginCtx.Get("sqls"); exists {
			if sqlArray, ok := sqls.(contextValue.SQLS); ok {
				sqlArray = append(sqlArray, sqlInfo)
				ginCtx.Set("sqls", sqlArray)
			}
		} else {
			ginCtx.Set("sqls", contextValue.SQLS{
				sqlInfo,
			})
		}
	}
	return
}

func CSTLayoutString() string {
	ts := time.Now()
	if cst, err := time.LoadLocation("Asia/Shanghai"); err != nil {
		return ""
	} else {
		return ts.In(cst).Format("2006-01-02 15:04:05")
	}
}
