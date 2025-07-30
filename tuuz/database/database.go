package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	gorose "github.com/tobycroft/gorose-pro"
	"main.go/tuuz/Log"
)

// 全局数据库实例
var Database *gorose.Engin

// 初始化数据库连接
func Init() {
	var err error
	Database, err = gorose.Open(&gorose.Config{
		Driver:          "sqlite3",   // 使用SQLite驱动
		Dsn:             "./data.db", // 数据库文件路径
		Prefix:          "",          // 表前缀
		SetMaxOpenConns: 10,          // 最大打开连接数
		SetMaxIdleConns: 5,           // 最大空闲连接数
	})

	if err != nil {
		Log.Crrs(err, "数据库连接失败")
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	Database.TagName("orm")
	Database.IgnoreName("ignore")
}

// GetDB 获取数据库实例
func GetDB() *gorose.Engin {
	return Database
}
