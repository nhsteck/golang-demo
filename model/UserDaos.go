package model

import (
	database "g.ghn.vn/scte-common/godal"
)

const SQL_CUSTOM_GET = "" +
	"SELECT * " +
	"FROM users " +
	"WHERE LENGTH(id) > $1 "

func CustomGet(dbType database.IDatabase, lenId int) (interface{}, error) {
	params := []interface{}{lenId}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET, params)
}
