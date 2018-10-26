package main

import (
	"testing"
	"fmt"
)

var dsn = "root:admin@tcp(127.0.0.1:3306)/test"

func TestNewMysql(t *testing.T) {
	m, err := NewMysql(dsn)
	if err != nil {
		t.Error(err.Error())
	}
	defer m.Close()
}

func TestMysql_ShowTables(t *testing.T) {
	m, err := NewMysql(dsn)
	if err != nil {
		t.Error(err.Error())
	}
	defer m.Close()

	tables, err := m.ShowTables()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(tables)
}

func TestMysql_DescTable(t *testing.T) {
	m, err := NewMysql(dsn)
	if err != nil {
		t.Error(err.Error())
	}
	defer m.Close()

	tables, err := m.DescTable("user_info")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(tables)
}

func TestMysql_TableIsExists(t *testing.T) {
	m, err := NewMysql(dsn)
	if err != nil {
		t.Error(err.Error())
	}
	defer m.Close()

	ok, err := m.TableIsExists("user_info")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(ok)
}

func TestMysql_ShowCreateTable(t *testing.T) {
	m, err := NewMysql(dsn)
	if err != nil {
		t.Error(err.Error())
	}
	defer m.Close()

	createTable, err := m.ShowCreateTable("user_info")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(createTable)
}