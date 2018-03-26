package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbc *gorm.DB

func init() {
	var err error
	dbc, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/ttkeeper")
	if err != nil {
		panic("failed to connect database")
	}
	dbc.LogMode(true)
	// Migrate the schema
	// dbc.AutoMigrate(&User{})
	// dbc.AutoMigrate(&Role{})
	// dbc.AutoMigrate(&UserRole{})
	// dbc.AutoMigrate(&Resource{})
	// dbc.AutoMigrate(&Permission{})
	// dbc.AutoMigrate(&RolePermission{})

	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	// var product Product
	// db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)
}
