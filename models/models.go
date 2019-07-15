package models

import (
    "log"
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "MumiChat/pkg/setting"
)

var db *gorm.DB

type Model struct {
    ID int `gorm:"primary_key" json:"id"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
}

func init() {
    var (
        err error
        dbType, dbName, dbPath string
    )

    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal(2, "Fail to get section 'database': %v", err)
    }

    dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	dbPath = sec.Key("PATH").String()
    // tablePrefix = sec.Key("TABLE_PREFIX").String()

	// sqlite3
    db, err = gorm.Open(dbType, fmt.Sprintf("%s%s.db", 
        dbPath, 
        dbName))

    if err != nil {
        log.Println(err)
    }

    // gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
    //     return tablePrefix + defaultTableName;
    // }

    db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}