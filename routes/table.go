package routes

import (
	"database/sql"
	"dummygpt/common"
	"fmt"
	_ "github.com/lib/pq"
)

func CheckTableExists(db *sql.DB, table common.Table) (bool, error) {
	query := fmt.Sprintf(`
    SELECT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE  table_schema = 'public'
        AND    table_name   = '%s'
    );`, table.Name)
	var exists bool
	err := db.QueryRow(query).Scan(&exists)
	return exists, err
}

func CreateTable(db *sql.DB, table common.Table) (sql.Result, error) {
	query := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (`, table.Name)
	for _, field := range table.Fields {
		query += fmt.Sprintf("%s %s,", field.Name, field.Datatype)
	}
	query = query[:len(query)-1]
	query += `);`
	res, err := db.Exec(query)
	return res, err
}

func AlterTable(db *sql.DB, table common.Table) (sql.Result, error) {
	query := fmt.Sprintf(`
    ALTER TABLE %s (`, table.Name)
	for _, field := range table.Fields {
		query += fmt.Sprintf("%s %s,", field.Name, field.Datatype)
	}
	query = query[:len(query)-1]
	query += `);`
	fmt.Println(query)
	res, err := db.Exec(query)
	return res, err
}
