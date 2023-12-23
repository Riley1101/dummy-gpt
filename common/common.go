package common

import (
	"encoding/json"
	"fmt"
)

type Field struct {
	Name     string
	Datatype string
}

type Table struct {
	Name   string
	Fields []Field
}

type SchemaForm struct {
	Message string `json:"message"`
}

type DbSchema struct {
	Fields []Field
}

func (dbSchema *DbSchema) ParseSchema(str string) {
	err := json.Unmarshal([]byte(str), &dbSchema)
	fmt.Println(dbSchema)
	if err != nil {
		fmt.Println(err)
	}
}

func (dbSchema *DbSchema) PrintSchema(name string) string {
	var str string
	str += fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (`, name)
	for _, field := range dbSchema.Fields {
		str += fmt.Sprintf("%s %s,", field.Name, field.Datatype)
	}
	str = str[:len(str)-1]
	str += `);`
	fmt.Println(str)
	return str
}

func ParseDBSchema(str string) DbSchema {
	var jsonSchema DbSchema
	err := json.Unmarshal([]byte(str), &jsonSchema)
	fmt.Println(jsonSchema)
	if err != nil {
		fmt.Println(err)
	}
	return jsonSchema
}
