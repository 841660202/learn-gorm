package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 有多张 CreditCard，UserID 是外键
//
//	type User struct {
//		ID           uint
//		Name         string
//		Email        *string
//		Age          uint8
//		Birthday     *time.Time
//		MemberNumber sql.NullString
//		ActivatedAt  sql.NullTime
//		CreatedAt    time.Time
//		UpdatedAt    time.Time
//	}
type User struct {
	ID           uint
	Name         string `gorm:"default:galeone"`
	Age          int32  `gorm:"default:18"`
	Email        *string
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Active       sql.NullBool `gorm:"default:true"`
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// var user User
	// result := db.Debug().First(&user)
	// fmt.Println(result)
	// fmt.Println("")
	// fmt.Println(user)
	// fmt.Println("")
	// fmt.Println(&user)
	// fmt.Println("")
	// fmt.Println(user.Name) //  Jinzhu
	// fmt.Println(&user.Name) // 0xc0000a00a8
	// [2.781ms] [rows:1] SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// &{0xc000104120 <nil> 1 0xc000112380 0}

	// {1 Jinzhu 18 <nil> 2022-10-19 10:46:59.668 +0800 CST { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-10-19 10:46:59.668 +0800 CST 2022-10-19 10:46:59.668 +0800 CST {true true}}

	// &{1 Jinzhu 18 <nil> 2022-10-19 10:46:59.668 +0800 CST { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-10-19 10:46:59.668 +0800 CST 2022-10-19 10:46:59.668 +0800 CST {true true}}

	// Jinzhu
	// 0xc0000a00a8

	// var user User
	// db.Limit(1).Find(&user)
	// fmt.Print(user)

	// var users []User

	// db.Find(&users, []int{1, 2, 3})

	// fmt.Print(users)

	// [
	// {1 Jinzhu 18 <nil> 2022-10-19 10:46:59.668 +0800 CST { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-10-19 10:46:59.668 +0800 CST 2022-10-19 10:46:59.668 +0800 CST {true true}}
	// {2 Jinzhu-1 18 <nil> <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-10-19 10:51:25.764 +0800 CST 2022-10-19 10:51:25.764 +0800 CST {true true}}
	// {3 Jinzhu-2 18 <nil> <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-10-19 10:52:54.86 +0800 CST 2022-10-19 10:52:54.86 +0800 CST {true true}}
	// ]% // 这里怎么有个%呢？

	// var user = User{ID: 10}
	// db.Debug().First(&user)
	// // [5.284ms] [rows:1] SELECT * FROM `users` WHERE `users`.`id` = 10 ORDER BY `users`.`id` LIMIT 1

	// var result User
	// db.Debug().Model(User{ID: 10}).First(&result)
	// // [0.914ms] [rows:1] SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// var result User
	// res := db.Debug().Model(User{ID: 10}).First(&result) // 这么用又问题，查的是第一个，文档有误
	// fmt.Print(res)
	// fmt.Print(result)

	// // [0.670ms] [rows:1] SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	// // &{0xc000202120 <nil> 1 0xc000212380 0}{1 Jinzhu 18 <nil> 2022-10-19 10:46:59.668 +0800 CST { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-10-19 10:46:59.668 +0800 CST 2022-10-19 10:46:59.668 +0800 CST {true true}}%

	//  查全部
	// var users []User
	// result := db.Debug().Find(&users)
	// fmt.Print(result.Error)

	//	[14.915ms] [rows:21] SELECT * FROM `users`
	// <nil>%

	// 条件查询
	// var user User
	// var users []User
	// Get first matched record
	// db.Debug().Where("name = ?", "jinzhu").First(&user)
	// 2022/10/19 12:18:34 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:110
	// [9.451ms] [rows:1] SELECT * FROM `users` WHERE name = 'jinzhu' ORDER BY `users`.`id` LIMIT 1

	// Get all matched records
	// db.Debug().Where("name <> ?", "jinzhu").Find(&users)
	// 2022/10/19 12:18:34 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:114
	// [3.397ms] [rows:19] SELECT * FROM `users` WHERE name <> 'jinzhu'

	// IN
	// db.Debug().Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// [3.971ms] [rows:1] SELECT * FROM `users` WHERE name IN ('jinzhu','jinzhu 2')

	// LIKE
	// db.Debug().Where("name LIKE ?", "%jin%").Find(&users)
	// [0.865ms] [rows:6] SELECT * FROM `users` WHERE name LIKE '%jin%'

	// AND
	// db.Debug().Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// [1.120ms] [rows:0] SELECT * FROM `users` WHERE name = 'jinzhu' AND age >= '22'

	// Time

	// db.Debug().Where("updated_at > ?", time.Date(2022, 10, 19, 11, 05, 43, 938, time.Local)).Find(&users)
	// [1.067ms] [rows:8] SELECT * FROM `users` WHERE updated_at > '2022-10-19 11:05:43'

	// BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// 3小时后
	// db.Debug().Where("updated_at > ?", time.Now().Add(time.Hour*3)).Find(&users)

	// 2022/10/19 13:39:23 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:134
	// [14.216ms] [rows:0] SELECT * FROM `users` WHERE updated_at > '2022-10-19 16:39:23.391'

	// 3小时前
	// db.Debug().Where("updated_at > ?", time.Now().Add(-time.Hour*3)).Find(&users)

	// 👑 ~/haotian/github/go/learn-gorm git:(main) ✗ $ go run ./crud/r.go
	// 2022/10/19 13:40:51 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:134
	// [4.806ms] [rows:20] SELECT * FROM `users` WHERE updated_at > '2022-10-19 10:40:51.291'
	// 👑 ~/haotian/github/go/learn-gorm git:(main) ✗ $

	// db.Debug().Where("updated_at between ? and ?", time.Now().Add(-time.Hour*3), time.Now().Add(-time.Hour/2*5)).Find(&users)
	// 这里第二个时间不能写成 time.Now().Add(-time.Hour*2.5)
	// [1.998ms] [rows:17] SELECT * FROM `users` WHERE updated_at between '2022-10-19 10:44:10.321' and '2022-10-19 11:14:10.321'

	// Struct
	// var user User
	// db.Debug().Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	// [7.996ms] [rows:0] SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu' AND `users`.`age` = 20 ORDER BY `users`.`id` LIMIT 1

	// Map
	// db.Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// [1.461ms] [rows:0] SELECT * FROM `users` WHERE `age` = 20 AND `name` = 'jinzhu'

	// Slice of primary keys
	// db.Debug().Where([]int64{20, 21, 22}).Find(&users)
	// [1.270ms] [rows:2] SELECT * FROM `users` WHERE `users`.`id` IN (20,21,22)

	// 当使用struct查询时，GORM只对非零字段进行查询
	// 0，''，false或其他零值，它将不会被用来建立查询条件
	// db.Debug().Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// [1.239ms] [rows:1] SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu'

	// 如果想要包含零值查询条件，你可以使用 map，其会包含所有 key-value 的查询条件
	// db.Debug().Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	// db.Debug().Where("Age").Find(&users)
	// [3.155ms] [rows:7] SELECT * FROM `users` WHERE Age // 这个就有意思了，看似sql写了一半，其实，这里查Age非零

	// 当使用 struct 进行查询时，你可以通过向 Where() 传入 struct 来指定查询条件的字段、值、表名
	// db.Debug().Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
	// [3.016ms] [rows:0] SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu' AND `users`.`age` = 0

	// db.Debug().Where(&User{Name: "jinzhu"}, "Age").Find(&users)
	// [0.456ms] [rows:13] SELECT * FROM `users` WHERE `users`.`age` = 0

	//  上面这两句中间差了一个“name”，结构体赋值了，后面不跟，就没有，这应该是gorm中在组织sql时候进行了struct丢弃
	// db.Debug().Where(&User{Age: 18}, "CreatedAt").Find(&users) // 这里并没有将 created_at 做 空处理
	// 验证猜测
	// [5.075ms] [rows:0] SELECT * FROM `users` WHERE `users`.`created_at` = '0000-00-00 00:00:00'

	// 如果做空处理，sql组装时候不拼接，那么查询结果应该是如下
	// mysql> select * from users where created_at;
	// +----+-----------------------------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
	// | id | name                              | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              | active |
	// +----+-----------------------------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
	// |  1 | Jinzhu                            | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |      1 |
	// |  2 | Jinzhu-1                          | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |      1 |
	// |  3 | Jinzhu-2                          | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |      1 |
	// |  5 | jinzhu1                           | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
	// |  6 | jinzhu2                           | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
	// |  7 | jinzhu3                           | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
	// |  8 | Batches_1                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |      1 |
	// |  9 | Batches_2                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |      1 |
	// | 10 | Batches_3                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
	// | 11 | Batches_4                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
	// | 12 | Batches_4                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
	// | 13 | Batches_1-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |      1 |
	// | 14 | Batches_2-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |      1 |
	// | 15 | Batches_3-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
	// | 16 | Batches_4-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
	// | 17 | Batches_4-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
	// | 19 | galeone                           | NULL  |   18 | 2022-10-19 11:21:56.089 | NULL          | NULL         | 2022-10-19 11:21:56.090 | 2022-10-19 11:21:56.090 |      1 |
	// | 20 | galeone                           | NULL  |   18 | 2022-10-19 11:26:50.011 | NULL          | NULL         | 2022-10-19 11:26:50.012 | 2022-10-19 11:26:50.012 |      1 |
	// | 21 | 我要存false，不要默认true         | NULL  |   18 | 2022-10-19 11:28:52.852 | NULL          | NULL         | 2022-10-19 11:28:52.852 | 2022-10-19 11:28:52.852 |      0 |
	// +----+-----------------------------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
	// 19 rows in set (0.00 sec)

	// 下面是sql执行，可以看到没有结果，有20个 警告
	// mysql> select * from users where name;
	//
	//	Empty set, 20 warnings (0.00 sec)

	// var user User
	// // Get by primary key if it were a non-integer type
	// db.Debug().First(&user, "id = ?", "string_primary_key")
	// // 找不到有报错
	// // 2022/10/19 14:07:27 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:230 record not found
	// // [3.930ms] [rows:0] SELECT * FROM `users` WHERE id = 'string_primary_key' ORDER BY `users`.`id` LIMIT 1

	// // Plain SQL
	// db.Debug().Find(&user, "name = ?", "jinzhu")
	// // 2022/10/19 14:07:27 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:234
	// // [1.960ms] [rows:1] SELECT * FROM `users` WHERE name = 'jinzhu'

	// db.Debug().Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// // 2022/10/19 14:07:27 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:237
	// // [2.164ms] [rows:0] SELECT * FROM `users` WHERE name <> 'jinzhu' AND age > 20

	// // Struct
	// db.Debug().Find(&users, User{Age: 20})
	// // 2022/10/19 14:07:27 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:245
	// // [0.688ms] [rows:0] SELECT * FROM `users` WHERE `age` = 20

	// // [3.930ms] [rows:0] SELECT * FROM `users` WHERE id = 'string_primary_key' ORDER BY `users`.`id` LIMIT 1

	// // Map
	// db.Debug().Find(&users, map[string]interface{}{"age": 20})
	// // SELECT * FROM users WHERE age = 20;

	// Not只是条件不执行
	// var user User
	// db.Debug().Not("name = ?", "jinzhu").First(&user)
	// // 2022/10/19 14:12:11 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:257
	// // [2.305ms] [rows:1] SELECT * FROM `users` WHERE NOT name = 'jinzhu' ORDER BY `users`.`id` LIMIT 1
	// // mysql> SELECT * FROM `users` WHERE NOT name = 'jinzhu' ORDER BY `users`.`id` LIMIT 1;
	// // +----+----------+-------+------+----------+---------------+--------------+-------------------------+-------------------------+--------+
	// // | id | name     | email | age  | birthday | member_number | activated_at | created_at              | updated_at              | active |
	// // +----+----------+-------+------+----------+---------------+--------------+-------------------------+-------------------------+--------+
	// // |  2 | Jinzhu-1 | NULL  |   18 | NULL     | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |      1 |
	// // +----+----------+-------+------+----------+---------------+--------------+-------------------------+-------------------------+--------+
	// // 1 row in set (0.00 sec)

	// // Not In
	// // map[string]interface{} 每次见到这个都觉得奇怪 键为string类型，值 interface{} 任意类型
	// https://www.cnblogs.com/ricklz/p/9494661.html
	// db.Debug().Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// // 2022/10/19 14:12:11 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:262
	// // [1.174ms] [rows:19] SELECT * FROM `users` WHERE `name` NOT IN ('jinzhu','jinzhu 2')

	// // Struct
	// db.Debug().Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// // 2022/10/19 14:12:11 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:266 record not found
	// // [0.712ms] [rows:0] SELECT * FROM `users` WHERE (`users`.`name` <> 'jinzhu' AND `users`.`age` <> 18) AND `users`.`id` = 2 ORDER BY `users`.`id` LIMIT 1

	// // Not In slice of primary keys
	// db.Debug().Not([]int64{1, 2, 3}).First(&user) // 哈哈，这里把我上面查到的值带过来了
	// // 2022/10/19 14:12:11 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:270 record not found
	// // [0.663ms] [rows:0] SELECT * FROM `users` WHERE `users`.`id` NOT IN (1,2,3) AND `users`.`id` = 2 ORDER BY `users`.`id` LIMIT 1

	// or条件

	// db.Debug().Where("name = ?", "Jinzhu-1").Or("name = ?", " Batches_1").Find(&users)
	// [6.507ms] [rows:1] SELECT * FROM `users` WHERE name = 'Jinzhu-1' OR name = ' Batches_1'

	// Struct
	// db.Debug().Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// [2.043ms] [rows:1] SELECT * FROM `users` WHERE name = 'jinzhu' OR (`users`.`name` = 'jinzhu 2' AND `users`.`age` = 18)

	// Map
	// db.Debug().Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// [3.578ms] [rows:1] SELECT * FROM `users` WHERE name = 'jinzhu' OR (`age` = 18 AND `name` = 'jinzhu 2')

	// 选择特定字段

	// db.Debug().Select("name", "age").Find(&users)
	// [4.937ms] [rows:21] SELECT `name`,`age` FROM `users`

	// db.Debug().Select([]string{"name", "age"}).Find(&users)
	// [2.717ms] [rows:21] SELECT `name`,`age` FROM `users`

	// db.Debug().Table("users").Select("COALESCE(age,?)", 42).Rows()
	// // SELECT COALESCE(age,'42') FROM users;

	//   mysql> select * from users;
	// +----+-----------------------------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
	// | id | name                              | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              | active |
	// +----+-----------------------------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
	// |  1 | Jinzhu                            | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |      1 |
	// |  2 | Jinzhu-1                          | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |      1 |
	// |  3 | Jinzhu-2                          | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |      1 |
	// |  4 | NULL                              | NULL  | NULL | 2022-10-19 10:54:25.773 | NULL          | NULL         | NULL                    | 2022-10-19 10:54:25.774 |      1 |
	// |  5 | jinzhu1                           | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
	// |  6 | jinzhu2                           | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
	// |  7 | jinzhu3                           | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
	// |  8 | Batches_1                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |      1 |
	// |  9 | Batches_2                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |      1 |
	// | 10 | Batches_3                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
	// | 11 | Batches_4                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
	// | 12 | Batches_4                         | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
	// | 13 | Batches_1-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |      1 |
	// | 14 | Batches_2-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |      1 |
	// | 15 | Batches_3-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
	// | 16 | Batches_4-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
	// | 17 | Batches_4-1                       | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
	// | 18 | 使用map                           | NULL  |   18 | NULL                    | NULL          | NULL         | NULL                    | NULL                    |      1 |
	// | 19 | galeone                           | NULL  |   18 | 2022-10-19 11:21:56.089 | NULL          | NULL         | 2022-10-19 11:21:56.090 | 2022-10-19 11:21:56.090 |      1 |
	// | 20 | galeone                           | NULL  |   18 | 2022-10-19 11:26:50.011 | NULL          | NULL         | 2022-10-19 11:26:50.012 | 2022-10-19 11:26:50.012 |      1 |
	// | 21 | 我要存false，不要默认true         | NULL  |   18 | 2022-10-19 11:28:52.852 | NULL          | NULL         | 2022-10-19 11:28:52.852 | 2022-10-19 11:28:52.852 |      0 |
	// +----+-----------------------------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
	// 21 rows in set (0.01 sec)

	// mysql> SELECT COALESCE(age,42) FROM `users`;
	// +------------------+
	// | COALESCE(age,42) |
	// +------------------+
	// |               18 |
	// |               18 |
	// |               18 |
	// |               42 | // 这里null被做42返回了，可以理解为查询的42是默认值，返回null的时候改为42
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |                0 |
	// |               18 |
	// |               18 |
	// |               18 |
	// |               18 |
	// +------------------+
	// 21 rows in set (0.00 sec)

	// 排序
	// var users []User
	// db.Debug().Order("age desc, name").Find(&users)
	// [2.374ms] [rows:21] SELECT * FROM `users` ORDER BY age desc, name

	// Multiple orders
	// db.Debug().Order("age desc").Order("name").Find(&users)
	// [2.804ms] [rows:21] SELECT * FROM `users` ORDER BY age desc,name

	// Limit & Offset
	// Limit指定要检索的最大记录数。 Offset指定在开始返回记录前要跳过的记录数。
	// var users []User
	// var users1 []User
	// var users2 []User

	// db.Debug().Limit(3).Find(&users)
	// // [2.278ms] [rows:3] SELECT * FROM `users` LIMIT 3

	// // Cancel limit condition with -1
	// db.Debug().Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// // 2022/10/19 14:35:38 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:383
	// // [0.373ms] [rows:10] SELECT * FROM `users` LIMIT 10

	// // 2022/10/19 14:35:38 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:383
	// // [0.357ms] [rows:21] SELECT * FROM `users`

	// db.Debug().Offset(3).Find(&users) // 报错，没有limit

	// db.Debug().Limit(10).Offset(5).Find(&users)
	// [1.108ms] [rows:10] SELECT * FROM `users` LIMIT 10 OFFSET 5
	// db.Debug().Offset(10).Find(&users1).Offset(-1).Find(&users2) // 报错，没有limit

	// var user User
	//
	// db.Model(&User{}).Select("name,id, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&user)
	// id与group是冲突的

	// 2022/10/19 14:45:08 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:399 Error 1055: Expression #1 of ORDER BY clause is not in GROUP BY clause and contains nonaggregated column 'learn-gorm.users.id' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by
	// [9.352ms] [rows:0] SELECT name, sum(age) as total FROM `users` WHERE name LIKE 'group%' GROUP BY `name` ORDER BY `users`.`id` LIMIT 1

	// db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&user)
	// 2022/10/19 14:49:37 /Users/haotian/haotian/github/go/learn-gorm/crud/r.go:399 record not found
	// [4.833ms] [rows:0] SELECT name, sum(age) as total FROM `users` WHERE name LIKE 'group%' GROUP BY `name` LIMIT 1
	// 👑 ~/haotian/github/go/learn-gorm git:(main) ✗ $

	// var users []User
	// db.Debug().Distinct("name", "age").Order("name, age desc").Find(&users)
	// [2.627ms] [rows:18] SELECT DISTINCT `name`,`age` FROM `users` ORDER BY name, age desc

	// db.Debug().Distinct("name").Distinct("age").Order("name, age desc").Find(&users) // 错误语句

	// Scan

	type Result struct {
		Name string
		Age  int
	}

	var result Result
	// db.Debug().Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)
	// [1.733ms] [rows:0] SELECT name,age FROM `users` WHERE name = 'Antonio'
	// Raw SQL
	db.Debug().Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
	// [2.076ms] [rows:0] SELECT name, age FROM users WHERE name = 'Antonio'
}
