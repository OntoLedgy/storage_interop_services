//go:build sqlite3
// +build sqlite3

// Package database_i_o_service/sqlite_driver.go contains only the Driver for the sqlite3
// database_i_o_service. It will get only included in the build if the tag `sqlite3` is
// specified.
//
// Default build of tables-to-go does NOT include sqlite3 support.
//
// Support for sqlite3 can be enabled by specifiying the tag while
// building tables-to-go:
//
//		go {install/build} -mod=vendor -tags sqlite3 .
//
// Alternative the Makefile can be used which is an alias for the go command
// above:
//
//		make sqlite3
//
package sqlite

import (
	// sqlite3 database_i_o_service Driver
	_ "github.com/mattn/go-sqlite3"
)
