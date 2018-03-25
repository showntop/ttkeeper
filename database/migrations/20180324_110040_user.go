package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20180324_110040 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20180324_110040{}
	m.Created = "20180324_110040"

	migration.Register("User_20180324_110040", m)
}

// Run the migrations
func (m *User_20180324_110040) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(`id` int(11) NOT NULL AUTO_INCREMENT,`username` varchar(128) NOT NULL,`password` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20180324_110040) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
