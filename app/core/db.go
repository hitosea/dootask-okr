package core

import (
	"dootask-okr/config"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB                *gorm.DB
	Session           gorm.Session
	ErrRecordNotFound = gorm.ErrRecordNotFound
	Expr              = gorm.Expr
)

// InDB 加载数据库
func InDB(str string) (*gorm.DB, error) {
	sp := strings.Split(str, "://")
	dbType := "sqlite"
	dbPath := str
	if len(sp) == 2 {
		dbType = strings.ToLower(sp[0])
		dbPath = sp[1]
	}
	// 配置项
	dbConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.CONF.System.Prefix, // 表名前缀
			SingularTable: false,                     // true:单数 false:复数，统一和原dootask复数表名，此时，`User` 的表名应该是 `t_users`
		},
	}

	// 如果开关打开，则添加日志记录器
	if os.Getenv("MYSQL_EnableSqlLog") == "true" {
		dbConfig.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 是否启用彩色打印
			},
		)
	}

	// 数据库类型
	if dbType == "mysql" {
		return gorm.Open(mysql.Open(dbPath), dbConfig)
	} else {
		return gorm.Open(sqlite.Open(dbPath), dbConfig)
	}
}

// CloseDB 关闭数据库
func CloseDB(db *gorm.DB) {
	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
}

// InitDB 初始化数据库
func InitDB() error {
	db, err := InDB(config.CONF.System.Dsn)
	if err != nil {
		return err
	}
	// defer CloseDB(db)
	DB = db
	return nil
}

// DBTableName 获取表名
func DBTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(model)
	return stmt.Schema.Table
}
