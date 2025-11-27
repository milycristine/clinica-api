package utils

import (
	sql "clinica-api/database"
	_ "embed"
	"log"
	"math/rand"
	"time"
)

var ConnectionDb *sql.SQLStr

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SetSQLConn(l *sql.SQLStr) {
	if l == nil {
		log.Fatal("A conexão com o banco de dados não pode ser nula.")
	}
	ConnectionDb = l
}

func GerarStringAleatoria(tamanho int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	b := make([]byte, tamanho)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
