package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCategoriesTable_20191211_155958 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCategoriesTable_20191211_155958{}
	m.Created = "20191211_155958"

	_ = migration.Register("CreateCategoriesTable_20191211_155958", m)
}

// Run the migrations
func (m *CreateCategoriesTable_20191211_155958) Up() {
	m.SQL("CREATE TABLE categories (" +
		"id VARCHAR(255) NOT NULL PRIMARY KEY, " +
		"code VARCHAR(20) NOT NULL, " +
		"name VARCHAR(50) NOT NULL, " +
		"created_at TIMESTAMP, " +
		"updated_at TIMESTAMP)")

}

// Reverse the migrations
func (m *CreateCategoriesTable_20191211_155958) Down() {
	m.SQL("DROP TABLE categories")

}
