package main

import (
	"fmt"
	"pigeon/im-auth-route/db"
	"pigeon/im-auth-route/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user string = "debian-sys-maint"
var pwd string = "8B2e4SuzKFCpn7AD"
var dbname string = "im_auth_route"
var host string = "127.0.0.1"
var port int = 3306

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, dbname)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = model.Migrate(gormDB)
	if err != nil {
		panic(err)
	}
	err = gormDB.Create(&model.UserModel{
		Username:       "markity",
		PasswordSha256: db.ToSha256([]byte("mark2004")),
	}).Error
	if err != nil {
		panic(err)
	}
	err = gormDB.Create(&model.UserModel{
		Username:       "usertest2",
		PasswordSha256: db.ToSha256([]byte("mark2004")),
	}).Error
	if err != nil {
		panic(err)
	}
}
