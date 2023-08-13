package db

import (
	_map "layout/pkg/helper/map"
)

type SQL struct {
	Timestamp   string  `json:"timestamp"`     // 时间，格式：2006-01-02 15:04:05
	Stack       string  `json:"stack"`         // 文件地址和行号
	SQL         string  `json:"sql"`           // SQL 语句
	Rows        int64   `json:"rows_affected"` // 影响行数
	CostSeconds float64 `json:"cost_seconds"`  // 执行时长(单位秒)
}

type SQLS = []SQL

var m *_map.SafeMap

func init() {
	m = _map.NewSafeMap()
}

func ClearSql(traceId string) {
	m.Del(traceId)
}

func AppendSql(sql SQL, traceId string) {
	if sqls, found := m.Get(traceId); found {
		if sqlArray, ok := sqls.(SQLS); ok {
			sqlArray = append(sqlArray, sql)
			m.Set(traceId, sqlArray)
		}
	} else {
		sqlArray := SQLS{
			sql,
		}
		m.Set(traceId, sqlArray)
	}
}

func GetAllSql(traceId string) SQLS {
	if sqlArray, ok := m.Get(traceId); ok {
		if sqls, ok := sqlArray.(SQLS); ok {
			return sqls
		}
	}
	return nil
}
