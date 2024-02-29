package db

import (
	"fmt"
	"strings"
	"time"

	"cufoon.litkeep.service/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *conf.Conf) (*gorm.DB, error) {
	cgm := conf.Mysql
	cgg := conf.Gorm

	fmt.Println(cgm.DSN())
	db, err := gorm.Open(mysql.Open(cgm.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if strings.ToLower(cgg.Debug) == "true" {
		db = db.Debug()
	}

	sql, _ := db.DB()
	err = sql.Ping()
	if err != nil {
		return nil, err
	}

	sql.SetMaxIdleConns(cgg.MaxIdleConns)
	sql.SetMaxOpenConns(cgg.MaxOpenConns)
	sql.SetConnMaxLifetime(time.Duration(cgg.MaxLifetime) * time.Second)

	return db, nil
}
