package SkuModel

import (
	"fmt"
	"math"
	"strings"

	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "product"

func ProductInit() {
	dbSql := `CREATE TABLE IF NOT EXISTS "product" (
	 "product_id" TEXT NOT NULL PRIMARY KEY,
	 "shop_id" integer NOT NULL default 0,
	 "name" TEXT NOT NULL default "",
	 "img" TEXT NOT NULL default "",
	 "market_price" integer NOT NULL default 0,
	 "discount_price" integer NOT NULL default 0,
	 "price_lower" integer NOT NULL default 0,
	 "price_higher" integer NOT NULL default 0,
	 "product_material_map" TEXT NOT NULL default "",
	 "pics" TEXT NOT NULL default "",
	 "status_del" integer NOT NULL default 0
	)`
	affected_rows, err := tuuz.Db().Execute(dbSql)
	if err != nil {
		panic(err.Error())
	}
	if affected_rows == 0 {
		return
	}
}

// Api_insert 添加商品
// func Api_insert(item DataItem) (int64, error) {
// 	db := tuuz.Db().Table(Table)
// 	db.Data(item)
// 	insertId, err := db.InsertGetId()
// 	if err != nil {
// 		Log.Dbrr(err, tuuz.FUNCTION_ALL())
// 		return 0, err
// 	}
// 	return insertId, nil
// }

func Api_insert(gl DataItem) bool {
	db := tuuz.Db().Table(Table)
	data := map[string]interface{}{
		"product_id":     gl.ProductID,
		"shop_id":        gl.ShopID,
		"name":           gl.Name,
		"img":            gl.Img,
		"market_price":   gl.MarketPrice,
		"discount_price": gl.DiscountPrice,
		"price_lower":    gl.PriceLower,
		"price_higher":   gl.PriceHigher,
		//"product_material_map": gl.ProductMaterialMap,
		"pics":       strings.Join(gl.Pics, " "), //string([]byte(gl.Pics[0])),
		"status_del": gl.StatusDel,
	}
	db.Data(data)
	//db.Data(gl)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func safeInt64ToInt(i64 int64) (int, error) {
	if i64 < math.MinInt || i64 > math.MaxInt {
		return 0, fmt.Errorf("int64值%d超出int范围", i64)
	}
	return int(i64), nil
}

// Api_find_by_id 通过ID查找商品
func Api_find_by_id(id string) (DataItem, error) {
	db := tuuz.Db().Table(Table)
	db.Where("product_id", id)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return DataItem{}, err
	} else {
		item := DataItem{
			ProductID: ret["product_id"].(string),
			Name:      ret["name"].(string),
			Img:       ret["img"].(string),
		}
		item.Pics = strings.Split(ret["pics"].(string), " ")
		fixnum, err := safeInt64ToInt(ret["shop_id"].(int64))
		if err == nil {
			item.ShopID = fixnum
		}
		numstatus, err := safeInt64ToInt(ret["status_del"].(int64))
		if err == nil {
			item.StatusDel = numstatus
		}
		return item, nil
	}
}

// Api_find_by_sku_id 通过SKU ID查找商品
func Api_find_by_sku_id(skuId string) (DataItem, error) {
	db := tuuz.Db().Table(Table)
	db.Where("product_id", skuId)
	var item DataItem
	err := db.Scan(&item)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return DataItem{}, err
	}
	return item, nil
}

// Api_select_by_shop 获取店铺的所有商品
func Api_select_by_shop(shopId int64, page, pageSize int) ([]DataItem, int64, error) {
	db := tuuz.Db().Table(Table)
	db.Where("shop_id", shopId)

	// 分页处理
	offset, limit := tuuz.Paginate(page, pageSize)
	db.Offset(offset).Limit(limit)

	var items []DataItem
	err := db.Scan(&items)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil, 0, err
	}

	// 获取总数
	total, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return items, 0, err
	}

	return items, total, nil
}

// Api_update 更新商品信息
func Api_update(id string, data gorose.Data) error {
	db := tuuz.Db().Table(Table)
	db.Where("product_id", id)
	_, err := db.Update(data)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
	}
	return err
}

// Api_update_price 更新商品价格
func Api_update_price(id string, price float64) error {
	return Api_update(id, gorose.Data{"price": price})
}

// Api_update_stock 更新商品库存
func Api_update_stock(id string, stock int) error {
	return Api_update(id, gorose.Data{"stock": stock})
}

// Api_toggle_status 切换商品状态
func Api_toggle_status(id string, status int) error {
	return Api_update(id, gorose.Data{"status_del": status})
}

// Api_delete 删除商品
func Api_delete(id int64) error {
	db := tuuz.Db().Table(Table)
	db.Where("product_id", id)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
	}
	return err
}

func Api_find(uid string) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"product_id": uid,
	}
	db.Where(where)
	_, err := db.First()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(uid interface{}, limit, page int) ([]DataItem, int) {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"shop_id": uid,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	// db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil, 0
	}

	var items []DataItem
	for _, data := range ret {
		item := DataItem{
			ProductID: data["product_id"].(string),
			Name:      data["name"].(string),
			Img:       data["img"].(string),
		}
		item.Pics = strings.Split(data["pics"].(string), " ")
		fixnum, err := safeInt64ToInt(data["shop_id"].(int64))
		if err == nil {
			item.ShopID = fixnum
		}
		numstatus, err := safeInt64ToInt(data["status_del"].(int64))
		if err == nil {
			item.StatusDel = numstatus
		}
		items = append(items, item)
	}
	return items, len(items)
}

func Api_select_all(uid interface{}) ([]DataItem, int64, error) {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"shop_id": uid,
	}
	db.Where(where)

	var total int64
	total, err := db.Count()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil, 0, err
	}

	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil, 0, err
	}
	var items []DataItem
	for _, data := range ret {
		item := DataItem{
			ProductID: data["product_id"].(string),
			Name:      data["name"].(string),
			Img:       data["img"].(string),
		}
		item.Pics = strings.Split(data["pics"].(string), " ")
		fixnum, err := safeInt64ToInt(data["shop_id"].(int64))
		if err == nil {
			item.ShopID = fixnum
		}
		numstatus, err := safeInt64ToInt(data["status_del"].(int64))
		if err == nil {
			item.StatusDel = numstatus
		}
		items = append(items, item)
	}

	return items, total, nil
}

func Api_count(uid interface{}) int64 {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"shop_id": uid,
	}
	db.Where(where)
	var count int64
	count, err := db.Count()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	}
	return count
}
