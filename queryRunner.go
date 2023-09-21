package main

import "database/sql"

func selectQuery(connection *sql.DB, query string, params []interface{}) (map[string]interface{}, error) {
	statement, err := connection.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	return nil, nil
}
