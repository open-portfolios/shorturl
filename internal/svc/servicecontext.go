// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"fmt"
	"os"

	"github.com/cylixlee/shorturl/internal/config"
	"github.com/cylixlee/shorturl/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	MapModel      model.MapModel
	SequenceModel model.SequenceModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		c.Mysql.User,
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.DatabaseName,
		c.Mysql.Charset,
		c.Mysql.ParseTime,
		c.Mysql.Locale,
	)
	conn := sqlx.NewMysql(dsn)
	return &ServiceContext{
		Config:        c,
		MapModel:      model.NewMapModel(conn),
		SequenceModel: model.NewSequenceModel(conn),
	}
}
