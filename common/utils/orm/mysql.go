package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gowork/common/utils/log"
)

func NewMysqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:mysql@/dmgame?charset=utf8")
	log.Error2Exit(err, "xorm.NewEngine mysql connect error")
	return engine
}
