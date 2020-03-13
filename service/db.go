package service

import (
	"fmt"
	"log"

	"gamenews.niracler.com/monitor/model"
	"gamenews.niracler.com/monitor/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open(setting.DBType, fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		setting.DBHost,
		setting.DBUser,
		setting.DBName,
		setting.DBPassword,
	))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DBTablePrefix + defaultTableName
	}

	// db表迁移:
	db.SingularTable(true)
	if err = db.AutoMigrate(&model.Game{}, &model.UserOperation{}, &model.VisitorCount{}).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// db.LogMode(model.Conf.ShowSQL)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}
