package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// belong to
type Dog struct {
  gorm.Model
  Name string
  GirlGodId uint // 狗链子
}

type GirlGod struct {
  gorm.Model
  Name string
	Dog Dog
}

func main() {
  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
  dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&GirlGod{}, &Dog{})
// 创建舔狗
	d:= Dog{
		// Model:     gorm.Model{
		// 	ID: 1,
		// },
		Name:      "舔狗",
	}
  // 创建女神
  g:= GirlGod{
  	// Model: gorm.Model{
    //   ID: 1,
    // },
  	Name:  "女神",
		Dog: d,
  }

	// 女神可以单独创建，如果没有给女神Dog
	db.Create(&g)
  


}