// package main

// import (
// 	"database/sql"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // User 有多张 CreditCard，UserID 是外键
// //
// //	type User struct {
// //		ID           uint
// //		Name         string
// //		Email        *string
// //		Age          uint8
// //		Birthday     *time.Time
// //		MemberNumber sql.NullString
// //		ActivatedAt  sql.NullTime
// //		CreatedAt    time.Time
// //		UpdatedAt    time.Time
// //	}
// type User struct {
// 	ID           uint
// 	Name         string `gorm:"default:galeone"`
// 	Age          int32  `gorm:"default:18"`
// 	Email        *string
// 	Birthday     *time.Time
// 	MemberNumber sql.NullString
// 	ActivatedAt  sql.NullTime
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	Active       sql.NullBool `gorm:"default:true"`
// }

// func main() {
// 	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
// 	dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// db.AutoMigrate(&User{})
// 	// db.AutoMigrate(&User{}) // 只会创建users表

// 	// 新增
// 	// user := User{Name: "Jinzhu", Age: 18, Birthday: &time.Now() } // 这里不能缩写报红
// 	// result := db.Create(&user) // 通过数据的指针来创建
// 	// // user.ID             // 返回插入数据的主键
// 	// // result.Error        // 返回 error
// 	// // result.RowsAffected // 返回插入记录的条数
// 	// // &{0xc0000cc630 <nil> 1 0xc00024c1c0 0}%

// 	// +----+--------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// | id | name   | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              |
// 	// +----+--------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// |  1 | Jinzhu | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |
// 	// +----+--------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+

// 	// 指定创建
// 	// t:=time.Now()
// 	// user := User{Name: "Jinzhu-1", Age: 18, Birthday: &t }

// 	// result:=db.Select("Name", "Age", "CreatedAt").Create(&user)

// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// | id | name     | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              |
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// |  1 | Jinzhu   | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |
// 	// |  2 | Jinzhu-1 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+

// 	// 指定创建
// 	// t:=time.Now()
// 	// user := User{Name: "Jinzhu-2", Age: 18, Birthday: &t }
// 	// result:=db.Select("Name", "Age").Create(&user) // created_at即使使用了select，也会创建

// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// | id | name     | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              |
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// |  1 | Jinzhu   | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |
// 	// |  2 | Jinzhu-1 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |
// 	// |  3 | Jinzhu-2 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+

// 	// 剔除字段
// 	// t:=time.Now()
// 	// user := User{Name: "Jinzhu-3", Age: 18, Birthday: &t }
// 	// db.Omit("Name", "Age", "CreatedAt").Create(&user) //剔除，可以剔除掉 created_at
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// | id | name     | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              |
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// |  1 | Jinzhu   | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |
// 	// |  2 | Jinzhu-1 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |
// 	// |  3 | Jinzhu-2 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |
// 	// |  4 | NULL     | NULL  | NULL | 2022-10-19 10:54:25.773 | NULL          | NULL         | NULL                    | 2022-10-19 10:54:25.774 |
// 	// +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// fmt.Print(result)

// 	// 批量插入
// 	// var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
// 	// db.Create(&users)

// 	// for _, user := range users {
// 	// 	fmt.Printf("Id = %v, ", user.ID) // 1,2,3
// 	// }
// 	// // +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// // | id | name     | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              |
// 	// // +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// // |  1 | Jinzhu   | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |
// 	// // |  2 | Jinzhu-1 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |
// 	// // |  3 | Jinzhu-2 | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |
// 	// // |  4 | NULL     | NULL  | NULL | 2022-10-19 10:54:25.773 | NULL          | NULL         | NULL                    | 2022-10-19 10:54:25.774 |
// 	// // |  5 | jinzhu1  | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |
// 	// // |  6 | jinzhu2  | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |
// 	// // |  7 | jinzhu3  | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |
// 	// // +----+----------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+

// 	// var users = []User{{Name: "Batches_1"}, {Name: "Batches_2"}, {Name: "Batches_3"}, {Name: "Batches_4"}, {Name: "Batches_4"}}

// 	// // 数量为 100
// 	// result := db.CreateInBatches(users, 2)
// 	// fmt.Print(result)

// 	// var users = []User{{Name: "Batches_1-1"}, {Name: "Batches_2-1"}, {Name: "Batches_3-1"}, {Name: "Batches_4-1"}, {Name: "Batches_4-1"}}

// 	// // 数量为 2,  5条 分了三批处理
// 	// result := db.Debug().CreateInBatches(users, 2)

// 	// 2022/10/19 11:05:43 /Users/haotian/haotian/github/go/learn-gorm/crud/c.go:119
// 	// [0.680ms] [rows:2] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Batches_1-1',NULL,0,NULL,NULL,NULL,'2022-10-19 11:05:43.938','2022-10-19 11:05:43.938'),('Batches_2-1',NULL,0,NULL,NULL,NULL,'2022-10-19 11:05:43.938','2022-10-19 11:05:43.938')

// 	// 2022/10/19 11:05:43 /Users/haotian/haotian/github/go/learn-gorm/crud/c.go:119
// 	// [0.345ms] [rows:2] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Batches_3-1',NULL,0,NULL,NULL,NULL,'2022-10-19 11:05:43.939','2022-10-19 11:05:43.939'),('Batches_4-1',NULL,0,NULL,NULL,NULL,'2022-10-19 11:05:43.939','2022-10-19 11:05:43.939')

// 	// 2022/10/19 11:05:43 /Users/haotian/haotian/github/go/learn-gorm/crud/c.go:119
// 	// [0.263ms] [rows:1] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Batches_4-1',NULL,0,NULL,NULL,NULL,'2022-10-19 11:05:43.939','2022-10-19 11:05:43.939')
// 	// &{0xc000207050 <nil> 5 0xc00023a000 0}%

// 	// fmt.Print(result)

// 	// 根据 map 创建记录时，association 不会被调用，且主键也不会自动填充
// 	// 主键会自增
// 	// db.Model(&User{}).Create(map[string]interface{}{
// 	// 	"Name": "使用map", "Age": 18,
// 	// })

// 	// | 18 | 使用map     | NULL  |   18 | NULL                    | NULL          | NULL         | NULL                    | NULL                    |
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+

// 	/*尝试更新表*/

// 	// db.Debug().AutoMigrate(&User{}) // 更新表结构，但不更新表数据

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:255
// 	// [0.408ms] [rows:-] SELECT DATABASE()

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:258
// 	// [8.334ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'learn-gorm%' ORDER BY SCHEMA_NAME='learn-gorm' DESC,SCHEMA_NAME limit 1

// 	// 2022/10/19 11:13:06 /Users/haotian/haotian/github/go/learn-gorm/crud/c.go:154
// 	// [0.593ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'learn-gorm' AND table_name = 'users' AND table_type = 'BASE TABLE'

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:255
// 	// [0.185ms] [rows:-] SELECT DATABASE()

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:258
// 	// [1.384ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'learn-gorm%' ORDER BY SCHEMA_NAME='learn-gorm' DESC,SCHEMA_NAME limit 1

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:168
// 	// [2.507ms] [rows:-] SELECT * FROM `users` LIMIT 1

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:186
// 	// [3.184ms] [rows:-] SELECT column_name, column_default, is_nullable = 'YES', data_type, character_maximum_length, column_type, column_key, extra, column_comment, numeric_precision, numeric_scale , datetime_precision FROM information_schema.columns WHERE table_schema = 'learn-gorm' AND table_name = 'users' ORDER BY ORDINAL_POSITION

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:52
// 	// [65.065ms] [rows:18] ALTER TABLE `users` MODIFY COLUMN `name` varchar(191) DEFAULT 'galeone'

// 	// 2022/10/19 11:13:06 /Users/haotian/go/pkg/mod/gorm.io/driver/mysql@v1.4.3/migrator.go:52
// 	// [51.806ms] [rows:18] ALTER TABLE `users` MODIFY COLUMN `age` int DEFAULT 18

// 	// mysql> desc users;
// 	// +---------------+---------------------+------+-----+---------+----------------+
// 	// | Field         | Type                | Null | Key | Default | Extra          |
// 	// +---------------+---------------------+------+-----+---------+----------------+
// 	// | id            | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
// 	// | name          | varchar(191)        | YES  |     | galeone |                |
// 	// | email         | longtext            | YES  |     | NULL    |                |
// 	// | age           | int(11)             | YES  |     | 18      |                |
// 	// | birthday      | datetime(3)         | YES  |     | NULL    |                |
// 	// | member_number | longtext            | YES  |     | NULL    |                |
// 	// | activated_at  | datetime(3)         | YES  |     | NULL    |                |
// 	// | created_at    | datetime(3)         | YES  |     | NULL    |                |
// 	// | updated_at    | datetime(3)         | YES  |     | NULL    |                |
// 	// +---------------+---------------------+------+-----+---------+----------------+
// 	// 9 rows in set (0.01 sec)

// 	// mysql> select * from users;
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// | id | name        | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              |
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// |  1 | Jinzhu      | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |
// 	// |  2 | Jinzhu-1    | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |
// 	// |  3 | Jinzhu-2    | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |
// 	// |  4 | NULL        | NULL  | NULL | 2022-10-19 10:54:25.773 | NULL          | NULL         | NULL                    | 2022-10-19 10:54:25.774 |
// 	// |  5 | jinzhu1     | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |
// 	// |  6 | jinzhu2     | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |
// 	// |  7 | jinzhu3     | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |
// 	// |  8 | Batches_1   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |
// 	// |  9 | Batches_2   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |
// 	// | 10 | Batches_3   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |
// 	// | 11 | Batches_4   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |
// 	// | 12 | Batches_4   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |
// 	// | 13 | Batches_1-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |
// 	// | 14 | Batches_2-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |
// 	// | 15 | Batches_3-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |
// 	// | 16 | Batches_4-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |
// 	// | 17 | Batches_4-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |
// 	// | 18 | 使用map     | NULL  |   18 | NULL                    | NULL          | NULL         | NULL                    | NULL                    |
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+
// 	// 18 rows in set (0.00 sec)

// 	// 增加  Active sql.NullBool `gorm:"default:true"`
// 	// 这里新增的column字段，自动会生成初始值
// 	// db.Debug().AutoMigrate(&User{})

// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
// 	// | id | name        | email | age  | birthday                | member_number | activated_at | created_at              | updated_at              | active |
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
// 	// |  1 | Jinzhu      | NULL  |   18 | 2022-10-19 10:46:59.668 | NULL          | NULL         | 2022-10-19 10:46:59.668 | 2022-10-19 10:46:59.668 |      1 |
// 	// |  2 | Jinzhu-1    | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:51:25.764 | 2022-10-19 10:51:25.764 |      1 |
// 	// |  3 | Jinzhu-2    | NULL  |   18 | NULL                    | NULL          | NULL         | 2022-10-19 10:52:54.860 | 2022-10-19 10:52:54.860 |      1 |
// 	// |  4 | NULL        | NULL  | NULL | 2022-10-19 10:54:25.773 | NULL          | NULL         | NULL                    | 2022-10-19 10:54:25.774 |      1 |
// 	// |  5 | jinzhu1     | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
// 	// |  6 | jinzhu2     | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
// 	// |  7 | jinzhu3     | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 10:58:15.865 | 2022-10-19 10:58:15.865 |      1 |
// 	// |  8 | Batches_1   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |      1 |
// 	// |  9 | Batches_2   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.017 | 2022-10-19 11:04:14.017 |      1 |
// 	// | 10 | Batches_3   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
// 	// | 11 | Batches_4   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
// 	// | 12 | Batches_4   | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:04:14.018 | 2022-10-19 11:04:14.018 |      1 |
// 	// | 13 | Batches_1-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |      1 |
// 	// | 14 | Batches_2-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.938 | 2022-10-19 11:05:43.938 |      1 |
// 	// | 15 | Batches_3-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
// 	// | 16 | Batches_4-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
// 	// | 17 | Batches_4-1 | NULL  |    0 | NULL                    | NULL          | NULL         | 2022-10-19 11:05:43.939 | 2022-10-19 11:05:43.939 |      1 |
// 	// | 18 | 使用map     | NULL  |   18 | NULL                    | NULL          | NULL         | NULL                    | NULL                    |      1 |
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+

// 	// 能够默认的就用默认字段，不能用默认的不管啦
// 	// t := time.Now()
// 	// user := User{Birthday: &t}
// 	// db.Debug().Create(&user) // created_at即使使用了select，也会创建

// 	// [5.752ms] [rows:1] INSERT INTO `users` (`name`,`age`,`email`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`active`)
// 	// VALUES ('galeone',18,NULL,'2022-10-19 11:21:56.088',NULL,NULL,'2022-10-19 11:21:56.09','2022-10-19 11:21:56.09',true)

// 	// | 19 | galeone     | NULL  |   18 | 2022-10-19 11:21:56.089 | NULL          | NULL         | 2022-10-19 11:21:56.090 | 2022-10-19 11:21:56.090 |      1 |
// 	// +----+-------------+-------+------+-------------------------+---------------+--------------+-------------------------+-------------------------+--------+
// 	// 0、''、false 等零值是不会保存到数据库 需要使用指针类型或 Scanner/Valuer 来避免这个问题

// 	// t := time.Now()
// 	// // Active       sql.NullBool `gorm:"default:true"`
// 	// user := User{Birthday: &t, Active: sql.NullBool{Bool: false}} // 这里不能直接写false, 下面这个注释也有意思，直接给我自动对其了
// 	// db.Debug().Create(&user)                                      // created_at即使使用了select，也会创建
// 	// 结果存入数据库的是true,依然是默认值
// 	// [3.668ms] [rows:1] INSERT INTO `users` (`name`,`age`,`email`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`active`)
// 	// VALUES ('galeone',18,NULL,'2022-10-19 11:26:50.011',NULL,NULL,'2022-10-19 11:26:50.012','2022-10-19 11:26:50.012',true)

// 	// t := time.Now()
// 	// // Active       sql.NullBool `gorm:"default:true"`
// 	// user := User{Name: "我要存false，不要默认true", Birthday: &t, Active: sql.NullBool{Bool: false, Valid: true}} // 这里不能直接写false, 下面这个注释也有意思，直接给我自动对其了
// 	// db.Debug().Create(&user)

// 	// [2.857ms] [rows:1] INSERT INTO `users` (`name`,`age`,`email`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`active`)
// 	// VALUES ('我要存false，不要默认true',18,NULL,'2022-10-19 11:28:52.851',NULL,NULL,'2022-10-19 11:28:52.852','2022-10-19 11:28:52.852',false)
// }
