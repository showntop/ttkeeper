package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Orgunit_20180324_111906 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Orgunit_20180324_111906{}
	m.Created = "20180324_111906"

	migration.Register("Orgunit_20180324_111906", m)
}

// Run the migrations
func (m *Orgunit_20180324_111906) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE orgunit(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`upunit` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Orgunit_20180324_111906) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `orgunit`")
}
