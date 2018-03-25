package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Permission_20180324_111512 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Permission_20180324_111512{}
	m.Created = "20180324_111512"

	migration.Register("Permission_20180324_111512", m)
}

// Run the migrations
func (m *Permission_20180324_111512) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE permission(`id` int(11) NOT NULL AUTO_INCREMENT,`action` int(11) DEFAULT NULL,`resource_id` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Permission_20180324_111512) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `permission`")
}
