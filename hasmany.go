// package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // User 有多张 CreditCard，UserID 是外键
// type User struct {
//   gorm.Model
//   CreditCards []CreditCard
// }

// type CreditCard struct {
//   gorm.Model
//   Number string
//   UserID uint
// }
// func main() {
//   // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//   dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


//   if err != nil {
//     panic("failed to connect database")
//   }

//   db.AutoMigrate(&User{}, &CreditCard{})
//   // db.AutoMigrate(&User{}) // 只会创建users表

// }