package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Message struct {
	ID    int    `json:"id"`
	Topic string `json:"topic"`
	Data  string `json:"data"`
	QoS   byte   `json:"qos"`
}


func handleMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Falha ao decodificar JSON", http.StatusBadRequest)
		return
	}

	err = saveMessage(message)
	if err != nil {
		http.Error(w, "Falha ao salvar mensagem no banco de dados", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Mensagem recebida na API - Tópico: %s, Dados: %s, QoS: %d\n", message.Topic, message.Data, message.QoS)

	w.WriteHeader(http.StatusOK)
}



// Inicializar banco de dados SQLite
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "../kil-db/kil-sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	// Criação da tabela para armazenar as mensagens
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		topic TEXT,
		data TEXT,
		qos BYTE
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

// Função para salvar mensagem no banco de dados
func saveMessage(message Message) error {
	stmt, err := db.Prepare("INSERT INTO messages(topic, data, qos) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.Topic, message.Data, message.QoS)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	initDB()

	http.HandleFunc("/messages", handleMessage)

	fmt.Println("API HTTP está rodando. Porta 8080. Pressione CTRL+C para sair.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
