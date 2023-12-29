package common

import (
	"encoding/json"
	"fmt"
	"log"
)

type Field struct {
	Name        string
	Datatype    string
	Constraints []string
}

type Table struct {
	Name   string
	Fields []Field
}

type SchemaForm struct {
	Schema string `json:"schema"`
}

type DbSchema struct {
	Name   string
	Fields []Field
}

func (dbSchema *DbSchema) ParseSchema(str string) {
	err := json.Unmarshal([]byte(str), &dbSchema)
	if err != nil {
		log.Fatalf("Error parsing schema: %v", err)
	}
}

func (dbSchema *DbSchema) DescribeSchema() string {
	var desc string
	for _, field := range dbSchema.Fields {
		desc += fmt.Sprintf(" %s %s %s \n", field.Name, field.Datatype, field.Constraints)
	}
	return desc
}

func (DbSchema *DbSchema) GenerateQuery() string {
	var query string
	query = fmt.Sprintf("CREATE TABLE %s (", DbSchema.Name)
	for _, field := range DbSchema.Fields {
		column := fmt.Sprintf("%s %s", field.Name, field.Datatype)
		for _, constraint := range field.Constraints {
			column += fmt.Sprintf(" %s", constraint)
		}
		query += fmt.Sprintf("%s,", column)
	}
	query = query[:len(query)-1]
	query += ");"
	log.Println(query)
	return query
}
