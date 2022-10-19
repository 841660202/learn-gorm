// package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
//   gorm.Model
//   Languages []Language `gorm:"many2many:user_languages;"`
// }

// type Language struct {
//   gorm.Model
//   Name string
// }

// func main() {
//   // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//   dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


//   if err != nil {
//     panic("failed to connect database")
//   }

//   // db.AutoMigrate(&User{}, &Language{})
//   db.AutoMigrate(&User{})

// }