package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductsTable_20191211_160620 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductsTable_20191211_160620{}
	m.Created = "20191211_160620"

	migration.Register("CreateProductsTable_20191211_160620", m)
}

// Run the migrations
func (m *CreateProductsTable_20191211_160620) Up() {
	m.SQL("CREATE TABLE products (" +
		"id VARCHAR(255) NOT NULL, " +
		"name VARCHAR(100) NOT NULL, " +
		"description VARCHAR(255), " +
		"category_id VARCHAR(255), " +
		"price VARCHAR(255) NOT NULL," +
		"created_at TIMESTAMP," +
		"updated_at TIMESTAMP, " +
		"PRIMARY KEY (id), " +
		"FOREIGN KEY(category_id) REFERENCES categories(id))")
}

// Reverse the migrations
func (m *CreateProductsTable_20191211_160620) Down() {
	m.SQL("DROP TABLE products")

}
