package postgres

import (
	"database/sql"
	"fmt"
	"github.com/marcopollivier/dio-expert-session-pre-class/model/transaction"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "diodb"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var db, _ = sql.Open("postgres", psqlInfo)
	return db
}

func Create(transaction transaction.Transaction) int {
	var db = connect()
	defer db.Close()

	var sqlStatement = `INSERT INTO  transactions (title, amount, type, installment, created_at)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING id;`

	var id int
	_ = db.QueryRow(sqlStatement,
					transaction.Title,
					transaction.Amount,
					transaction.Type,
					transaction.Installment,
					transaction.CreatedAt).Scan(&id)
	fmt.Println("New record ID is:", id)

	return id
}


func FetchAll() transaction.Transactions {
	var db = connect()
	defer db.Close()

	rows, _ := db.Query("SELECT title, amount, type, installment, created_at FROM transactions")
	defer rows.Close()

	var transactionSlice []transaction.Transaction
	for rows.Next() {
		var transaction transaction.Transaction
		_ = rows.Scan(&transaction.Title,
					  &transaction.Amount,
					  &transaction.Type,
					  &transaction.Installment,
					  &transaction.CreatedAt)

		transactionSlice = append(transactionSlice, transaction)
	}

	return transactionSlice
}