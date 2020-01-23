package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"testing"
	"time"
)

//批量插入
func Test_insert(t *testing.T) {
	O.Using("default") // 默认使用 default，你可以指定为其他数据库
	user := GoUser{Age: 28, Content: "赖豪达个人简介", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		user.Name = "samlai  - " + num
		user.Title = "title - " + num
		fmt.Println(user)
		id, err := O.Insert(&user)
		if err == nil {
			fmt.Printf("insert data id  =  %d ", id)
			user.Id = 0
			//fmt.Println("\n")
		}

	}
}

//读取数据内容
func Test_read(t *testing.T) {
	user := GoUser{Id: 35}
	err := O.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("查询不到主键内容")
	} else {
		fmt.Println(user)
	}
}

//批量的操作
func Test_InsertMulti(t *testing.T) {

	users := []GoUser{
		{Age: 28, Content: "我的28岁内容", Title: "28.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Age: 27, Content: "赖豪达的27岁", Title: "27.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Age: 26, Content: "啦啦啦，出来工作一年啦啦...", Title: "26.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	successNum, err := O.InsertMulti(len(users), users)
	if err == nil {
		fmt.Printf("success num  :  %d", successNum)
	}
}

//更新数据
func Test_updateOneData(t *testing.T) {

	user := GoUser{Id: 35}
	if O.Read(&user) == nil {
		user.Title = "id 为35的数据内容"
		if num, err := O.Update(&user); err == nil {
			fmt.Println(num)
		}
	}

	//如果指定某个字段的话则:
	// 只更新 Name
	//o.Update(&user, "Name")
	// 指定多个字段
	// o.Update(&user, "Field1", "Field2", ...)

}
