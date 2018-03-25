package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Role_20180324_110549 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Role_20180324_110549{}
	m.Created = "20180324_110549"

	migration.Register("Role_20180324_110549", m)
}

// Run the migrations
func (m *Role_20180324_110549) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE role(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Role_20180324_110549) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `role`")
}
