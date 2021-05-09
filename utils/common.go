package utils

import (
	"reflect"
	"strconv"

	"gorm.io/gorm"
)

func LikeQuery(db *gorm.DB, param interface{}) {
	types := reflect.TypeOf(param)
	values := reflect.ValueOf(param)
	for i := 0; i < types.NumField(); i++ {
		name := types.Field(i).Tag.Get("form")
		value := values.Field(i).String()
		if value != "" {
			db.Where(name + " LIKE " + "'%" + value + "%'")
		}
	}
}

func LikeOrEqualQuery(db *gorm.DB, param interface{}) {
	types := reflect.TypeOf(param)
	values := reflect.ValueOf(param)
	for i := 0; i < types.NumField(); i++ {
		name := types.Field(i).Tag.Get("field")

		switch types.Field(i).Tag.Get("mode") {
		case "like":
			value := values.Field(i).String()
			if value != "" {
				db.Where(name + " LIKE " + "'%" + value + "%'")
			}
		case "equal":
			switch values.Field(i).Type().String() {
			case "int":
				if values.Field(i).Int() != 0 {
					db.Where(name + " = " + strconv.Itoa(int(values.Field(i).Int())))
				}
			case "string":
				db.Where(name + " = " + "'" + values.Field(i).String() + "'")
			}
		}
	}
}

func Pagination(db *gorm.DB, page int, pageSize int) {
	if pageSize > 0 {
		db.Limit(pageSize)
		if page != 0 {
			db.Offset((page - 1) * pageSize)
		}
	}
}
