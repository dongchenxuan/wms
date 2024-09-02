package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"wms/config"
	"wms/log"
	"wms/models"
)

var Orm *xorm.Engine

func InitMysql() error {
	log.Infof("================[开始初始化Mysql数据库连接]================")
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Instance.Mysql.UserName,
		config.Instance.Mysql.Password,
		config.Instance.Mysql.Host,
		config.Instance.Mysql.Port,
		config.Instance.Mysql.Database,
		config.Instance.Mysql.Settings)
	orm, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		log.Fatalf("[InitMysql] connect mysql => %v\n", err)
		return err
	}
	if err = orm.Ping(); err != nil {
		log.Printf("[InitMysql][sql conn] %s", addr)
		return fmt.Errorf("链接数据库失败")
	}

	orm.ShowSQL(config.Instance.Mysql.ShowSQL)
	orm.SetMaxIdleConns(config.Instance.Mysql.Idle)
	orm.SetMaxOpenConns(config.Instance.Mysql.Max)

	if orm == nil {
		log.Error("[mysql] 连接失败")
	}

	syncTables(orm)

	Orm = orm
	log.Infof("================[结束初始化MySQL数据库连接]================")
	return nil
}

// todo: 同步表结构
func syncTables(orm *xorm.Engine) {
	err := orm.Sync2(&models.WmsArchivesDepotDao{})
	if err != nil {
		panic(err)
	}

	err = orm.Sync2(&models.WmsArchivesDepotDao{})
	if err != nil {
		panic(err)
	}
}
