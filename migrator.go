// package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // type User struct {
// // 	gorm.Model
// // 	Name string `gorm:"size:255;index:idx_name,unique"`
// // }

// type User struct {
// 	gorm.Model
// 	Name           string
// 	FavoriteBookId string
// 	FavoriteBook   FavoriteBook `gorm:"foreignKey:FavoriteBookId;references:Title"`
// }

// type FavoriteBook struct {
// 	Title  string `gorm:"primaryKey;"`
// 	Author string
// }

// func main() {
// 	dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	// 方式1
// 	// db.Debug().AutoMigrate(&User{})

// 	// 方式2
// 	// db.Migrator().CreateIndex(&User{}, "Name") // 重复操作
// 	// 2022/10/21 09:52:36 /Users/haotian/haotian/github/go/learn-gorm/migrator.go:18 Error 1061: Duplicate key name 'idx_name'
// 	// [2.300ms] [rows:0] CREATE UNIQUE INDEX `idx_name` ON `users`(`name`)

// 	// 方式3
// 	// db.Migrator().CreateIndex(&User{}, "idx_name") // 重复操作
// 	// 2022/10/21 09:52:59 /Users/haotian/haotian/github/go/learn-gorm/migrator.go:19 Error 1061: Duplicate key name 'idx_name'
// 	// [1.194ms] [rows:0] CREATE UNIQUE INDEX `idx_name` ON `users`(`name`)

// 	// db.Migrator().DropIndex(&User{}, "Name")

// 	//   mysql> show create table  users \G;
// 	// *************************** 1. row ***************************
// 	//        Table: users
// 	// Create Table: CREATE TABLE `users` (
// 	//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	//   `created_at` datetime(3) DEFAULT NULL,
// 	//   `updated_at` datetime(3) DEFAULT NULL,
// 	//   `deleted_at` datetime(3) DEFAULT NULL,
// 	//   `name` varchar(255) DEFAULT NULL,
// 	//   PRIMARY KEY (`id`),
// 	//   UNIQUE KEY `idx_name` (`name`), // 之前这里有，删除就没了
// 	//   KEY `idx_users_deleted_at` (`deleted_at`)
// 	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
// 	// 1 row in set (0.00 sec)

// 	// ERROR:
// 	// No query specified

// 	// mysql> show create table  users \G;
// 	// *************************** 1. row ***************************
// 	//        Table: users
// 	// Create Table: CREATE TABLE `users` (
// 	//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	//   `created_at` datetime(3) DEFAULT NULL,
// 	//   `updated_at` datetime(3) DEFAULT NULL,
// 	//   `deleted_at` datetime(3) DEFAULT NULL,
// 	//   `name` varchar(255) DEFAULT NULL,
// 	//   PRIMARY KEY (`id`),
// 	//   KEY `idx_users_deleted_at` (`deleted_at`)
// 	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
// 	// 1 row in set (0.00 sec)

// 	// 再加上
// 	// db.Migrator().CreateIndex(&User{}, "Name") // 重复操作
// 	//   mysql> show create table  users \G; #hello
// 	// *************************** 1. row ***************************
// 	//        Table: users
// 	// Create Table: CREATE TABLE `users` (
// 	//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	//   `created_at` datetime(3) DEFAULT NULL,
// 	//   `updated_at` datetime(3) DEFAULT NULL,
// 	//   `deleted_at` datetime(3) DEFAULT NULL,
// 	//   `name` varchar(255) DEFAULT NULL,
// 	//   PRIMARY KEY (`id`),
// 	//   UNIQUE KEY `idx_name` (`name`),
// 	//   KEY `idx_users_deleted_at` (`deleted_at`)
// 	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
// 	// 1 row in set (0.00 sec)

// 	// ERROR:
// 	// No query specified

// 	// 换种方式删除
// 	// db.Migrator().DropIndex(&User{}, "idx_name")
// 	// 不存在时候再删除试试
// 	// 2022/10/21 10:01:25 /Users/haotian/haotian/github/go/learn-gorm/migrator.go:84 Error 1091: Can't DROP 'idx_name'; check that column/key exists
// 	// [1.977ms] [rows:0] DROP INDEX `idx_name` ON `users`

// 	db.Debug().AutoMigrate(&User{}, &FavoriteBook{})
// }
