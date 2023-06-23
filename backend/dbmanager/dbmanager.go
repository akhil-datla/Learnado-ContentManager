/*
 * File: dbmanager.go
 * File Created: Sunday, 11th June 2023 9:57:15 pm
 * Last Modified: Friday, 23rd June 2023 12:43:18 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package dbmanager

import (
	"github.com/asdine/storm"
)

var db *storm.DB

// Open opens the database with the given name.
func Open(name string) error {
	var err error
	db, err = storm.Open(name)
	return err
}

// AutoCreateStruct creates a table for the given struct.
func AutoCreateStruct(data interface{}) error {
	err := db.Init(data)
	return err
}

// Save saves the data (struct) to the database.
func Save(data interface{}) error {
	err := db.Save(data)
	return err
}

// Query executes a query and retrieves a single record based on the specified field and value.
func Query(fieldName string, value, to interface{}) error {
	err := db.One(fieldName, value, to)
	return err
}

// GroupQuery executes a query and retrieves multiple records based on the specified field and value.
func GroupQuery(fieldName string, value, to interface{}) error {
	err := db.Find(fieldName, value, to)
	return err
}

// QueryAll retrieves all records of the specified struct from the database.
func QueryAll(to interface{}) error {
	err := db.All(to)
	return err
}

// Update updates the data (struct) in the database.
func Update(data interface{}) error {
	err := db.Update(data)
	return err
}

// Delete deletes the data (struct) from the database.
func Delete(data interface{}) error {
	err := db.DeleteStruct(data)
	return err
}

// Close closes the database connection.
func Close() error {
	return db.Close()
}
