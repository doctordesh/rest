package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func newDb() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	fmt.Printf("Creating new db\n")
	db, err := gorm.Open("mysql", "root:@/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.DropTableIfExists(&Invite{})
	db.DropTableIfExists(&InviteMember{})
	db.CreateTable(&Invite{})
	db.CreateTable(&InviteMember{})

	//db.Create(&Invite{Code: "toto", Token: "})

	return &db, nil
}
