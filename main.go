package main

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"
	"net/http" // Gerencia URLs e Servidor Web

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

//Struct utilizada para exibir dados no template

type Items struct {
	Id     int
	Item   string
	Amount int
}

// Função dbConn, abre a conexão com o banco de dados
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := ""
	dbPass := ""
	dbName := ""

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	// Exibe mensagem que o servidor foi iniciado
	log.Println("Server started on: http://localhost:9000")

	// Gerencia as URLs
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	// Ações
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	// Inicia o servidor na porta 9000
	http.ListenAndServe(":9000", nil)
}
