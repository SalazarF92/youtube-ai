package db

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectToDB(
	db_host string,
	db_port string,
	db_user string,
	db_password string,
	db_name string,
) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_password, db_name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println(psqlInfo)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao estabelecer uma conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db
}
