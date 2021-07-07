package model

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/panda8z/shorturl/utils"
)

var db *gorm.DB
var err error

func getDB() *gorm.DB {
	return db
}

// InitDb init
func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(utils.Db,
		dsn)

	if err != nil {
		log.Println("database connect err:", err.Error())
	}
	log.Println("mysql at ", dsn)
	db.SingularTable(true)
	db.LogMode(true)

	db.AutoMigrate(&Url{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)
	// 调试单个操作，显示此操作的详细日志
	// db.Debug().Where("username = ?", "panda").First(&User{})

}

// DBClose close
func DBClose() {
	db.Close()
}
