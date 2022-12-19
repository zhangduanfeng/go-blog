package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func MysqlInit() {
	var err error
	DB, err = gorm.Open(
		"mysql",
		"root:zdf112233.@(sh-cynosdbmysql-grp-1vg8w4ba.sql.tencentcdb.com:20182)/blog_db?parseTime=true")
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println("MySQL连接失败")
		panic(err)
	}
	fmt.Println("MySQL连接成功")
	DB.SingularTable(true)

}
