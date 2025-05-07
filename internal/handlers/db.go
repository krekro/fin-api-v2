package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"fin-api-go/internal/service"
)

type expense_data struct {
	Payment_id       uuid.UUID `db:"payment_id"`
	User_name        string    `db:"user_name"`
	Category         string    `db:"category"`
	Amount           float32   `db:"amount"`
	Create_date      time.Time `db:"create_date"`
	Transaction_desc string    `db:"transaction_desc"`
}

func connectPG(ctx context.Context, connString string) *pgx.Conn {

	conn, err := pgx.Connect(ctx, os.Getenv(connString))
	if err != nil {
		log.Printf("Unable to connect to DB : %s", err)
	}
	return conn
}

func sqlQuery(ctx context.Context, conn *pgx.Conn, query string) []expense_data {

	rows, _ := conn.Query(ctx, query)
	result, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[expense_data])
	if err != nil {
		log.Printf("Error processing rows : %s", err)
	}

	return result
}

//exported func

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	user_name := r.URL.Query().Get("user_name")
	session_id := r.URL.Query().Get("session_id")
	conn := connectPG(ctx, "RetoolDB_URL")

	//qeury section
	query := "select * from expense_data where user_name='" + user_name + "'"
	result := sqlQuery(ctx, conn, query)

	if service.ValidateSesssion(session_id) {
		//response section
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{"status": "success", "data": result}

		//log.Printf("%s", w)
		json.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{"status": "error", "message": "Invalid Session"}
		json.NewEncoder(w).Encode(response)
	}

}

func GetExpenseSummary(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	user_name := r.URL.Query().Get("user_name")
	session_id := r.URL.Query().Get("session_id")
	conn := connectPG(ctx, "RetoolDB_URL")

	//validate Session ID
	if service.ValidateSesssion(session_id) {
		//qeury section
		query := "SELECT category, SUM(amount) as amount, user_name FROM expense_data WHERE EXTRACT(MONTH from create_date)=EXTRACT(MONTH from CURRENT_DATE) AND user_name='" + user_name + "' GROUP BY category, user_name"
		result := sqlQuery(ctx, conn, query)

		//response section
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{"status": "success", "data": result}

		//log.Printf("%s", w)
		json.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{"status": "error", "message": "Invalid Session"}
		json.NewEncoder(w).Encode(response)
	}

}
