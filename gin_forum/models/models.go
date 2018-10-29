package models

import (
    _"github.com/go-sql-driver/mysql";
 	"github.com/jinzhu/gorm";
    _"github.com/jinzhu/gorm/dialects/mysql";
 	"log"
)

const (
	DB_DBMS = "mysql"
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_HOST = "127.0.0.1"
	DB_PORT = "3306"
	DB_DATABASE = "gin_forum"
)

type User struct {
    UserID uint `gorm:"primary_key"`
    Username string `gorm:"size:255;unique;not null"`
    Password string `gorm:"size:255;unique;not null"`
    Email string `gorm:"size:255;unique;not null"`
    Posts []Post `gorm:"foreignkey:UserID"`
    Deleted bool
}

type Post struct {
    PostID uint `gorm:"primary_key"`
    Content string `gorm:"size:510;unique;not null"`
    Title string `gorm:"size:255;unique;not null"`
    UserID uint
    Deleted bool
    Likes uint
}

func Make_connection( ) *gorm.DB {
    db, err := gorm.Open(DB_DBMS, DB_USERNAME+":@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE)

    if err != nil {
        log.Panic(err)
    }
    log.Println("Connection to the database made")
    return db
}

func Migrate( ) {
    db := Make_connection( )
    defer db.Close( )
    db.Debug( ).DropTableIfExists(&User{})
    db.Debug( ).AutoMigrate(&User{})
    db.Debug( ).DropTableIfExists(&Post{})
    db.Debug( ).AutoMigrate(&Post{})
}

