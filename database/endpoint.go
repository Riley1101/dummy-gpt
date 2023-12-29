package database

import (
	"database/sql"
)

type Endpoint struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	OwnerId  int    `json:"owner_id"`
	IsPublic bool   `json:"is_public"`
}

type EndpointQuery struct {
	CreateEndpoint      string
	CreateEndpointTable string
	GetEndpointsByOwner string
}

var EndpointQuerySql = EndpointQuery{
	CreateEndpoint: `INSERT INTO endpoints (name, owner_id, is_public) VALUES ($1, $2, $3) RETURNING id`,
	CreateEndpointTable: `CREATE TABLE IF NOT EXISTS endpoints (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL,
		owner_id INT NOT NULL,
		is_public BOOLEAN NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (owner_id) REFERENCES users(id)
	)`,
	GetEndpointsByOwner: `SELECT id, name, is_public FROM endpoints WHERE owner_id=$1`,
}

type EndpointDB struct {
	DB *sql.DB
}

func (e *EndpointDB) CreateEndpoint(endpoint Endpoint) (int, error) {
	var id int
	err := e.DB.QueryRow(EndpointQuerySql.CreateEndpoint, endpoint.Name, endpoint.OwnerId, endpoint.IsPublic).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (e *EndpointDB) GetEndpointsByOwner(owner_id int) ([]Endpoint, error) {
	endpoints := []Endpoint{}
	rows, err := e.DB.Query(EndpointQuerySql.GetEndpointsByOwner, owner_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var endpoint Endpoint
		err := rows.Scan(&endpoint.ID, &endpoint.Name, &endpoint.IsPublic)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}
	defer rows.Close()
	return endpoints, nil
}

func (e *EndpointDB) CreateEndpointTable() error {
	_, err := e.DB.Exec(EndpointQuerySql.CreateEndpointTable)
	return err
}
