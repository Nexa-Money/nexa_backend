package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	DB_URI := os.Getenv("DB_URI")

	config, err := pgxpool.ParseConfig(DB_URI)
	if err != nil {
		log.Fatalf("Erro ao configurar conexão com o banco: %v", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %v", err)
	}

	fmt.Println("Conectado ao banco com sucesso.")
	return pool
}
