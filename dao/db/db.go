package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

// Init 数据库初始化
func Init(dirverName,dsn string) (err error) {
	DB, err = sqlx.Open(dirverName,dsn)
	if err != nil {
		return err
	}

	// 查看连接是否成功
	err = DB.Ping()

	// 设置最大连接
	DB.SetMaxOpenConns(100)

	// 设置最大空闲
	DB.SetMaxIdleConns(16)
	return
}