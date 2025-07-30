package tuuz

import (
	"runtime"
	"strings"

	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz/database"
)

// Db 获取数据库实例
func Db() gorose.IOrm {
	return database.Database.NewOrm()
}

// FUNCTION_ALL 获取当前函数名
func FUNCTION_ALL() string {
	pc, _, _, _ := runtime.Caller(1)
	fullName := runtime.FuncForPC(pc).Name()

	// 简化函数名：只保留包名和函数名
	parts := strings.Split(fullName, ".")
	if len(parts) > 1 {
		return parts[len(parts)-2] + "." + parts[len(parts)-1]
	}
	return fullName
}

// Paginate 分页计算
func Paginate(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // 默认分页大小
	}
	return (page - 1) * pageSize, pageSize
}
