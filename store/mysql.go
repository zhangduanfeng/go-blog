package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

var DB *gorm.DB

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logrus.WithFields(
			logrus.Fields{
				"rows_returned": v[5],
			},
		).Info(v[3])
	case "log":
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "log"})
	}
}

func MysqlInit() {
	var err error
	DB, err = gorm.Open(
		"mysql",
		"root:zdf112233.@(sh-cynosdbmysql-grp-1vg8w4ba.sql.tencentcdb.com:20182)/blog_db?parseTime=true")
	if err != nil {
		logrus.Error("MySQL连接失败")
		panic(err)
	}
	logrus.Info("MySQL连接成功")
	DB.SingularTable(true)
	//空闲
	DB.DB().SetMaxIdleConns(50)
	//打开
	DB.DB().SetMaxOpenConns(100)
	//超时
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	DB.SetLogger(&GormLogger{})
	DB.LogMode(true)
}
