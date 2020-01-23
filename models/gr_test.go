package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	"testing"
)

var err error
var engin *gorose.Engin

func init() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	//mysql Dsn示例 "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true"
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "root:samlai123@tcp(localhost:3307)/go_db?charset=utf8&parseTime=true"})
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}

//读取数据内容
func Test_find_first(t *testing.T) {
	res, err := DB().Table("go_user").First()
	if err == nil {
		fmt.Println(res)
	}

}
