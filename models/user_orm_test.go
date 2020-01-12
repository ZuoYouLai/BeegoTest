package main
//
//import (
//	"fmt"
//	"github.com/astaxie/beego/orm"
//	"time"
//)
//
//func init() {
//	orm.RegisterDriver("mysql", orm.DRMySQL)
//
//	// 参数4(可选)  设置最大空闲连接
//	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
//	maxIdle := 30
//	maxConn := 30
//	//填写db 的 username password
//	orm.RegisterDataBase("default", "mysql", "root:samlai123@tcp(127.0.0.1:3307)//go_db?charset=utf8", maxIdle, maxConn)
//	// 设置为 UTC 时间
//	orm.DefaultTimeLoc = time.UTC
//}
//
//func main() {
//	o := orm.NewOrm()
//	o.Using("default") // 默认使用 default，你可以指定为其他数据库
//
//	user := new(User)
//	user.Age = 28
//	user.Title = "samlai-title"
//	user.Name = "samlai"
//	user.Content = "赖豪达个人简介"
//	user.CreatedAt = time.Now()
//	user.UpdatedAt = time.Now()
//
//	fmt.Println(o.Insert(user))
//
//}
