package endpoints

import "fmt"

type Column struct {
	Name     string
	Datatype string
}

type Schema struct {
	Name    string
	columns []Column
}

type Endpoint struct {
	id     int
	Url    string
	Name   string
	Schema Schema
}

func (e *Endpoint) GetUrl() string {
	return e.Url
}

func (e *Endpoint) CreateDB() {
	query := fmt.Sprintf("CREATE DATABASE %s", e.Schema.Name)
	fmt.Println(query)
}

func (e *Endpoint) CreateTable() {
	for _, column := range e.Schema.columns {
		query := fmt.Sprintf("CREATE TABLE %s (%s %s)", e.Schema.Name, column.Name, column.Datatype)
		fmt.Println(query)
	}
}


