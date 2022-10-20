package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID string
	Company   Company `gorm:"references:CompanyID"` // 使用 Company.CompanyID 作为引用
}

type Company struct {
	CompanyID int
	Code      string
	Name      string
}

// type Dog struct {
// 	ID   int
// 	Name string
// 	Toys []Toy `gorm:"polymorphic:Owner;"`
// }

// type Cat struct {
// 	ID   int
// 	Name string
// 	Toys []Toy `gorm:"polymorphic:Owner;"`
// }

// type Toy struct {
// 	ID        int
// 	Name      string
// 	OwnerID   int
// 	OwnerType string
// }

func main() {
	dsn := "root:123456@tcp(localhost:3306)/learn-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(&Dog{}, &Cat{}, &Toy{}) // === db.AutoMigrate(&Dog{}, &Cat{}, Toy{})

	// 	mysql> desc toys;
	// +------------+------------+------+-----+---------+----------------+
	// | Field      | Type       | Null | Key | Default | Extra          |
	// +------------+------------+------+-----+---------+----------------+
	// | id         | bigint(20) | NO   | PRI | NULL    | auto_increment |
	// | name       | longtext   | YES  |     | NULL    |                |
	// | owner_id   | bigint(20) | YES  |     | NULL    |                |
	// | owner_type | longtext   | YES  |     | NULL    |                |
	// +------------+------------+------+-----+---------+----------------+
	// 4 rows in set (0.01 sec)

	// mysql> drop table toys;
	// Query OK, 0 rows affected (0.00 sec)

	// mysql> desc toys;
	// +------------+------------+------+-----+---------+----------------+
	// | Field      | Type       | Null | Key | Default | Extra          |
	// +------------+------------+------+-----+---------+----------------+
	// | id         | bigint(20) | NO   | PRI | NULL    | auto_increment |
	// | name       | longtext   | YES  |     | NULL    |                |
	// | owner_id   | bigint(20) | YES  |     | NULL    |                |
	// | owner_type | longtext   | YES  |     | NULL    |                |
	// +------------+------------+------+-----+---------+----------------+
	// 4 rows in set (0.00 sec)

	db.AutoMigrate(&User{}, &Company{}) // === db.AutoMigrate(&Dog{}, &Cat{}, Toy{})

}
