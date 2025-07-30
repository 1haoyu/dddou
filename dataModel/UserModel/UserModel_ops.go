package UserModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "users"

func UserInit() {
	dbSql := `CREATE TABLE IF NOT EXISTS "users" (
	 "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 "login_name" TEXT NOT NULL default "",
	 "login_pass" TEXT NOT NULL default "",
	 "account_name" TEXT NOT NULL default "",
	 "cookie" TEXT NOT NULL default "",
	 "login_type" integer NOT NULL default 0,
	 "shop_count" integer NOT NULL default 0,
	 "expires" REAL
	)`
	affected_rows, err := tuuz.Db().Execute(dbSql)
	if err != nil {
		panic(err.Error())
	}
	if affected_rows == 0 {
		return
	}
}

// Api_find_by_username 根据用户名查找用户
func Api_find_by_username(username string) (UserInfo, error) {
	db := tuuz.Db().Table(Table)
	db.Where("login_name", username)

	var user UserInfo
	err := db.Scan(&user)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return UserInfo{}, err
	}
	return user, nil
}

// Api_insert 插入新用户
func Api_insert(user UserInfo) bool {
	db := tuuz.Db().Table(Table)
	// var data = map[string]interface{}{"login_name": user.LoginName, "login_pass": user.LoginPass, "cookie": user.Cookie}
	db.Data(user)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

// Api_update 更新用户信息
func Api_update(username string, data map[string]interface{}) error {
	db := tuuz.Db().Table(Table)
	db.Where("login_name", username)

	_, err := db.Update(data)
	return err
}

func UpdateUserCookie(username, cookie string, exp float64) error {
	// 更新用户表中的cookie字段
	updateData := map[string]interface{}{
		"cookie":  cookie,
		"expires": exp,
	}
	return Api_update(username, updateData)
}

func Api_select_struct[T UserInfo](name any) []T {
	db := tuuz.Db().Table(Table)
	if name != nil {
		db.Where("account_name", name)
	}
	ret := []T{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return []T{}
	} else {
		return ret
	}
}
