package executor

import "github.com/ainilili/bitdb/core/model"

type Executor interface {
	Query(sql string, args... interface{}) model.Result
}