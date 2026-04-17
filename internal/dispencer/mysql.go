package dispencer

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	query = "REPLACE INTO sequence (stub) VALUES ('a')"
)

type mysqlDispencer struct {
	conn sqlx.SqlConn
}

func NewMysql(conn sqlx.SqlConn) Interface {
	return &mysqlDispencer{conn: conn}
}

func (m *mysqlDispencer) Dispence(ctx context.Context) (uint64, error) {
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecCtx(ctx)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
