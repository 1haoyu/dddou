package CookieModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "cookie_info"

// type CookieInfo struct {
// 	LoginName    string   `json:"login_name"`
// 	LoginPass    string   `json:"login_pass"`
// 	AccountName  string   `json:"account_name"`
// 	LoginType    int      `json:"login_type"`
// 	Cookie       string   `json:"cookie"`
// 	Token        string   `json:"token"`
// 	VerifyFp     string   `json:"verifyFp"`
// 	MemberID     string   `json:"member_id"`
// 	SubjectID    string   `json:"subject_id"`
// 	CanLogin     bool     `json:"can_login"`
// 	EncodeShopID string   `json:"encode_shop_id"`
// 	IdentityType int      `json:"identity_type"`
// 	Expires      *float64 `json:"expires"`
// }

type CookieInfo struct {
	SubjectID    string  `gorose:"subject_id" json:"subject_id"`
	LoginName    string  `gorose:"login_name" json:"login_name"`
	AccountID    string  `gorose:"account_id" json:"account_id"`
	AccountName  string  `gorose:"account_name" json:"account_name"`
	LoginType    int     `gorose:"login_type" json:"login_type"`
	Cookie       string  `gorose:"cookie" json:"cookie"`
	Token        string  `gorose:"token" json:"token"`
	VerifyFp     string  `gorose:"verifyFp" json:"verifyFp" gorm:"column:verify_fp"`
	MemberID     string  `gorose:"member_id" json:"member_id"`
	CanLogin     bool    `gorose:"can_login" json:"can_login"`
	EncodeShopID string  `gorose:"encode_shop_id" json:"encode_shop_id"`
	IdentityType int     `gorose:"identity_type" json:"identity_type"`
	Expires      float64 `gorose:"expires" json:"expires"`
}

func CreateCookieInfoTable() {
	dbSql := `
        CREATE TABLE IF NOT EXISTS cookie_info (
        	subject_id TEXT NOT NULL PRIMARY KEY,
            login_name TEXT NOT NULL DEFAULT '',
            account_id TEXT NOT NULL DEFAULT '',
            account_name TEXT NOT NULL DEFAULT '',
            login_type INTEGER NOT NULL DEFAULT 0,
            cookie TEXT NOT NULL DEFAULT '',
            token TEXT NOT NULL DEFAULT '',
            verifyFp TEXT NOT NULL DEFAULT '',
            member_id TEXT NOT NULL DEFAULT '',
            can_login BOOLEAN NOT NULL DEFAULT 0,
            encode_shop_id TEXT NOT NULL DEFAULT '',
            identity_type INTEGER NOT NULL DEFAULT 0,
            expires REAL
        );
    `
	_, err := tuuz.Db().Execute(dbSql)
	if err != nil {
		panic(err.Error())
	}
}

func Api_insert(gl CookieInfo) bool {
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

func Api_insert_more(gls []CookieInfo) bool {
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

func Api_find_base() gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]any{
		"identity_type": 0,
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

func Api_update_token(name, token interface{}) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"subject_id": name,
	}
	db.Where(where)
	data := map[string]interface{}{
		"token": token,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update_verify(name, verifyFp interface{}) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"subject_id": name,
	}
	db.Where(where)
	data := map[string]interface{}{
		"verifyFp": verifyFp,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find_struct(name any) CookieInfo {
	db := tuuz.Db().Table(Table)
	db.Where("subject_id", name)
	ret := CookieInfo{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return CookieInfo{}
	} else {
		return ret
	}
}

func Api_find_struct_by_id(name any) CookieInfo {
	db := tuuz.Db().Table(Table)
	db.Where("account_id", name)
	ret := CookieInfo{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return CookieInfo{}
	} else {
		return ret
	}
}

func Api_select_struct[T CookieInfo](name any) []T {
	db := tuuz.Db().Table(Table)
	if name != nil {
		db.Where("login_name", name)
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
func Api_update(subjectid string, data map[string]interface{}) error {
	db := tuuz.Db().Table(Table)
	db.Where("subject_id", subjectid)

	_, err := db.Update(data)
	return err
}
