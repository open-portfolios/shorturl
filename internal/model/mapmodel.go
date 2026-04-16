package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MapModel = (*customMapModel)(nil)

type (
	// MapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMapModel.
	MapModel interface {
		mapModel
		withSession(session sqlx.Session) MapModel
	}

	customMapModel struct {
		*defaultMapModel
	}
)

// NewMapModel returns a model for the database table.
func NewMapModel(conn sqlx.SqlConn) MapModel {
	return &customMapModel{
		defaultMapModel: newMapModel(conn),
	}
}

func (m *customMapModel) withSession(session sqlx.Session) MapModel {
	return NewMapModel(sqlx.NewSqlConnFromSession(session))
}
