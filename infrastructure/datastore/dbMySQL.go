package datastore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMySQL() *gorm.DB {
	connectString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",
		"docker",
		"docker",
		"0.0.0.0",
		"3306",
		"my_testdb",
		"?charset=utf8&parseTime=True&loc=Local",
	)
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	db.Set("gorm:table_options", "ENGIN=InnoDB")

	return db
}