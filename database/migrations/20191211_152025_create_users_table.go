package main

import (
	"github.com/astaxie/beego/migration"
)

//CreateUsersTable_20191211_152025 :
type CreateUsersTable_20191211_152025 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20191211_152025{}
	m.Created = "20191211_152025"

	_ = migration.Register("CreateUsersTable_20191211_152025", m)
}

// Up :
func (m *CreateUsersTable_20191211_152025) Up() {
	m.SQL("CREATE TABLE users(" +
		"id VARCHAR(255) NOT NULL PRIMARY KEY, " +
		"username VARCHAR(100) NOT NULL," +
		"email VARCHAR(50) NOT NULL UNIQUE, " +
		"password VARCHAR(255) NOT NULL, " +
		"created_at TIMESTAMP, " +
		"updated_at TIMESTAMP) ")

}

// Down :
func (m *CreateUsersTable_20191211_152025) Down() {
	m.SQL("DROP TABLE users")

}
