package Log

import (
	"fmt"
	"log"
	"os"

	"github.com/tobycroft/gorose-pro"
)

// 确保日志目录存在
func ensureLogDir() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", 0777)
	}
}

// Write 写入日志到指定文件
func Write(file_name string, logs string, discript string, other string) {
	ensureLogDir()

	file, err := os.OpenFile("log/"+file_name+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	} else {
		logger := log.New(file, "", log.LstdFlags)
		logger.Println(logs, discript, other)
		file.Close()
	}
}

// Error 记录错误到指定日志文件
func Error(file_name string, err error) {
	if err != nil {
		Write(file_name, "", "", err.Error())
	}
}

// Err 记录错误到默认错误日志
func Err(err error) {
	if err != nil {
		Write("Error", "", "", err.Error())
	}
}

// Errs 记录带描述的错误到默认错误日志（同时打印到控制台）
func Errs(err error, log_str string) {
	fmt.Println(log_str, err)
	if err != nil {
		Write("Error", log_str, "", err.Error())
	}
}

// Drr 数据库错误记录
func Drr(err error) {
	if err != nil {
		Write("Database", "", "", err.Error())
	}
}

// Crr 通用错误记录
func Crr(logs error) {
	if logs != nil {
		Write("Common", "", "", logs.Error())
	}
}

// Crrs 带描述的通用错误记录（同时打印到控制台）
func Crrs(logs error, discript string) {
	fmt.Println(logs, discript)
	if logs != nil {
		Write("Common", "", discript, logs.Error())
	}
}

// Dbrr 数据库错误记录（带描述）
func Dbrr(err error, log_str string) {
	fmt.Println(err, log_str)
	if err != nil {
		Write("Dberror", log_str, "", err.Error())
	}
}

// DBrrsql 记录数据库错误及执行的SQL语句
func DBrrsql(err error, db gorose.IOrm, log_str string) {
	fmt.Println(err, "\n", db.LastSql(), "\n", log_str)
	if err != nil {
		Write("Dberror", log_str, db.LastSql(), err.Error())
	}
}
