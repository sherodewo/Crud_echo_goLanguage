package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCustomersTable_20191211_160042 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCustomersTable_20191211_160042{}
	m.Created = "20191211_160042"

	_ = migration.Register("CreateCustomersTable_20191211_160042", m)
}

// Run the migrations
func (m *CreateCustomersTable_20191211_160042) Up() {
	m.SQL("CREATE TABLE customers (" +
		"id VARCHAR(255) NOT NULL PRIMARY KEY, " +
		"name VARCHAR(100) NOT NULL," +
		"email VARCHAR(50) NOT NULL UNIQUE, " +
		"phone VARCHAR(13) NOT NULL UNIQUE," +
		"address VARCHAR(255), " +
		"created_at TIMESTAMP, " +
		"updated_at TIMESTAMP) ")

}

// Reverse the migrations
func (m *CreateCustomersTable_20191211_160042) Down() {
	m.SQL("DROP TABLE customers")

}
