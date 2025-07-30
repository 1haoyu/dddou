package ShopModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "shop"

func ShopInit() {
	dbSql := `CREATE TABLE IF NOT EXISTS "shop" (
	 "subject_id" TEXT NOT NULL PRIMARY KEY,
	 "account_name" TEXT NOT NULL default "",
	 "account_avatar" TEXT NOT NULL default "",
	 "identity_type" integer NOT NULL default 0,
	 "identity_type_desc" TEXT NOT NULL default "",
	 "bus_member_id" TEXT NOT NULL default "",
	 "member_id" TEXT NOT NULL default "",
	 "account_id" TEXT NOT NULL default "",
     "can_login" BOOLEAN NOT NULL DEFAULT 0,
	 "encode_shop_id" TEXT NOT NULL default "",
	 "login_name" TEXT NOT NULL default ""
	)`
	affected_rows, err := tuuz.Db().Execute(dbSql)
	if err != nil {
		panic(err.Error())
	}
	if affected_rows == 0 {
		return
	}
}

func Api_insert(gl Account) bool {
	db := tuuz.Db().Table(Table)
	db.Data(gl)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_insert_more(gls []Account) bool {
	db := tuuz.Db().Table(Table)
	db.Data(gls)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(name any) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]any{
		"account_name": name,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_struct(name any) Account {
	db := tuuz.Db().Table(Table)
	if name != nil {
		db.Where("account_name", name)
	}
	ret := Account{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return Account{}
	} else {
		return ret
	}
}

func Api_find_struct_by_id(name any) Account {
	db := tuuz.Db().Table(Table)
	if name != nil {
		db.Where("account_id", name)
	}
	ret := Account{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return Account{}
	} else {
		return ret
	}
}

func Api_select_struct[T Account](name any) []T {
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

// Api_update 更新用户信息
func Api_update(username string, data map[string]interface{}) error {
	db := tuuz.Db().Table(Table)
	db.Where("subject_id", username)

	_, err := db.Update(data)
	return err
}
