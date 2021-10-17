package koneksi

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDatabases() *gorm.DB{
	User := "root"
	Pass := ""
	Host := "localhost"
	Port := "3306"
	DBName := "db_golang"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", User, Pass, Host, Port, DBName)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		fmt.Println("koneksi Ke databases gagal", err)
	}
	return db
}