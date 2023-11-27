package _case

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

// 数据源
var dsn = "root:ranzhou@tcp(192.168.239.100:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, //默认字符串大小，默认太大

	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info), //日志级别
		PrepareStmt: true,                                //预编译, 不支持嵌套事务
	})

	if err != nil {
		log.Fatal(err.Error())
	}

}

// 设置连接池
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err.Error())
		return
	}

	sqlDB.SetMaxIdleConns(5)            //最大空闲连接数
	sqlDB.SetMaxOpenConns(10)           //最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //连接最长存活时间

}
