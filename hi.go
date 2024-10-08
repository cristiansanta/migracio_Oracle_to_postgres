package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
)

func main() {
	// Conectar a Oracle
	oracleDB, err := sql.Open("godror", "system/oracle_password@localhost:1521/XE")
	if err != nil {
		log.Fatal(err)
	}
	defer oracleDB.Close()

	// Conectar a PostgreSQL
	postgresDB, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres_password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	// Consulta para seleccionar datos de Oracle
	rows, err := oracleDB.Query("SELECT id, nombre, edad, salario FROM empleados")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Preparar la inserción en PostgreSQL
	stmt, err := postgresDB.Prepare("INSERT INTO empleados (id, nombre, edad, salario) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Iterar sobre los resultados de Oracle e insertar en PostgreSQL
	for rows.Next() {
		var id int
		var nombre string
		var edad int
		var salario float64
		err := rows.Scan(&id, &nombre, &edad, &salario)
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(id, nombre, edad, salario)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Migración completada con éxito")
}
