package main

// // go get -u gorm.io/gorm
// // go get -u gorm.io/driver/sqlite
// // ♠ /Users/chenhailong/code/github/go/gorm $ go get -u gorm.io/gorm
// // go: go.mod file not found in current directory or any parent directory.
// //         'go get' is no longer supported outside a module.
// //         To build and install a command, use 'go install' with a version,
// //         like 'go install example.com/cmd@latest'
// //         For more information, see https://golang.org/doc/go-get-install-deprecation
// //         or run 'go help get' or 'go help install'.

// // ♠ /Users/chenhailong/code/github/go/learn-gorm $ go mod init learn-gorm
// // go: creating new go.mod: module learn-gorm
// // go: to add module requirements and sums:
// // 	go mod tidy
// // ♠ /Users/chenhailong/code/github/go/learn-gorm $ go mod tidy
// // go: finding module for package gorm.io/gorm
// // go: finding module for package gorm.io/driver/mysql
// // go: downloading gorm.io/gorm v1.24.0
// // go: downloading gorm.io/driver/mysql v1.4.3
// // go: found gorm.io/driver/mysql in gorm.io/driver/mysql v1.4.3
// // go: found gorm.io/gorm in gorm.io/gorm v1.24.0
// // go: downloading github.com/go-sql-driver/mysql v1.6.0
// // go: downloading github.com/jinzhu/now v1.1.5
// // go: downloading github.com/jinzhu/inflection v1.0.0
// // ♠ /Users/chenhailong/code/github/go/learn-gorm $
// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // belong to
// type Dog struct {
//   gorm.Model
//   Name string
//   GirlGodId uint
//   GirlGod GirlGod
// }

// type GirlGod struct {
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

//   db.AutoMigrate(&Dog{}, &GirlGod{})

//   // 创建女神
//   g:= GirlGod{
//   	Model: gorm.Model{
//       ID: 1,
//     },
//   	Name:  "女神",
//   }
//   // 创建舔狗
//   d:= Dog{
//   	Model:     gorm.Model{
//       ID: 1,
//     },
//   	Name:      "舔狗",
//   	GirlGodId: 0,
//   	GirlGod:   g,
//   }
//   // 创建1个舔狗
//   // db.Create(d)
// //   panic: reflect.Value.Addr of unaddressable value

// // goroutine 1 [running]:
// // reflect.Value.Addr({0x136fb00?, 0xc0001da3f0?, 0x1383fc0?})
// //         /usr/local/go/src/reflect/value.go:271 +0x65
// // gorm.io/gorm/callbacks.SaveBeforeAssociations.func1(0xc000293860)
// //         /Users/chenhailong/go/pkg/mod/gorm.io/gorm@v1.24.0/callbacks/associations.go:82 +0x24f
// // gorm.io/gorm.(*processor).Execute(0xc0001e8190, 0xc0001e0990?)
// //         /Users/chenhailong/go/pkg/mod/gorm.io/gorm@v1.24.0/callbacks.go:130 +0x436
// // gorm.io/gorm.(*DB).Create(0x1383fc0?, {0x1383fc0?, 0xc0001da380})
// //         /Users/chenhailong/go/pkg/mod/gorm.io/gorm@v1.24.0/finisher_api.go:24 +0xa5
// // main.main()
// //         /Users/chenhailong/code/github/go/learn-gorm/one2one.go:74 +0x2fb
// // exit status 2
//   db.Create(&d)

// }

// CREATE TABLE `dogs` (
//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   `name` longtext,
//   `girl_god_id` bigint(20) unsigned DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `idx_dogs_deleted_at` (`deleted_at`),
//   KEY `fk_dogs_girl_god` (`girl_god_id`),
//   CONSTRAINT `fk_dogs_girl_god` FOREIGN KEY (`girl_god_id`) REFERENCES `girl_gods` (`id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

// CREATE TABLE `girl_gods` (
//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   `name` longtext,
//   PRIMARY KEY (`id`),
//   KEY `idx_girl_gods_deleted_at` (`deleted_at`)
// ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;