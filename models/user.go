package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type GoUser struct {
	Id        int
	Title     string
	Name      string
	Age       int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var O orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(GoUser))

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	//填写db 的 username password
	orm.RegisterDataBase("default", "mysql", "root:samlai123@tcp(127.0.0.1:3307)/go_db?charset=utf8", maxIdle, maxConn)
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC

	O = orm.NewOrm()
	orm.Debug = true

}

func main() {

}
