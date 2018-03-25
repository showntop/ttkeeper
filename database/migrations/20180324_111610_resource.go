package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Resource_20180324_111610 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Resource_20180324_111610{}
	m.Created = "20180324_111610"

	migration.Register("Resource_20180324_111610", m)
}

// Run the migrations
func (m *Resource_20180324_111610) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE resource(`id` int(11) NOT NULL AUTO_INCREMENT,`rtype` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Resource_20180324_111610) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `resource`")
}
