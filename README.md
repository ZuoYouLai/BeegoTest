## BeeGo 学习

### ORM 操作
+ 参考文章 : https://www.jianshu.com/p/4441b8b765ac
+ 官方文档 : https://beego.me/docs/mvc/model/overview.md 

```text

    1. 建库建表的操作:
    CREATE DATABASE IF NOT EXISTS go_db DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

    2. 创建 user 表:
    DROP TABLE IF EXISTS `go_user`;
    CREATE TABLE `go_user` (
          `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
          `title` varchar(50) DEFAULT NULL COMMENT '标题',
          `name` varchar(50) DEFAULT NULL COMMENT '姓名',
          `age` int(3) DEFAULT NULL COMMENT '年龄',
          `content` mediumtext COMMENT '个人的内容简介',
          `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
          `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
          PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='go 项目的 user 表内容';



```

+ 引入 mysql 的 driver
```text
go get github.com/go-sql-driver/mysql
```

### 日志打印
+ 例子:
```cassandraql
	logs.Warn("this is samlai")
	logs.Error("samlai error")
	logs.Info("info msg")
```