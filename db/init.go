package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uncleyd/core/config"
	"github.com/uncleyd/core/logger"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// key:groupId,value:{key:id,value:*gorm.Db}
var ormMap = make(map[string]map[string]*gorm.DB)

var isInit = false

// Init
func Init() {
	if isInit {
		logger.Sugar.Warnf("db is init")
		return
	}

	var mysqls = config.Get().MySql
	for _, v := range mysqls {
		if v == nil {
			continue
		}
		if !v.Enable {
			continue
		}

		logger.Sugar.Infow("init mysql db.",
			"groupId", v.GroupId,
			"id", v.Id,
			"dbName", v.DbName,
		)

		db, err := createORM(v)
		if err != nil {
			panic(err)
		}

		dbs := ormMap[v.GroupId]
		if dbs == nil {
			dbs = make(map[string]*gorm.DB)
			ormMap[v.GroupId] = dbs
		}
		dbs[v.Id] = db
	}

	isInit = true
}

func GetDb(id string) *gorm.DB {
	for _, group := range ormMap {
		for k, v := range group {
			if k == id {
				return v
			}
		}
	}
	return nil
}

func DBWithGroupId(groupId string) map[string]*gorm.DB {
	return ormMap[groupId]
}

func AdminDB() *gorm.DB {
	return GetDb("admin")
}

const (
	connectFormat = "%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

func createORM(cfg *config.MySqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(connectFormat, cfg.UserName, cfg.Password, cfg.Host, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnect)
	sqlDB.SetMaxOpenConns(cfg.MaXOpenConnect)
	sqlDB.SetConnMaxLifetime(28799)
	//sqlDB.(cfg.LogMode)

	return db, nil
}

func TableName(name string) string {
	return "tb_" + name
}
