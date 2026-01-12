package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func main() {
	ctx := context.Background()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}

	rows, err := conn.Query(ctx, "SELECT id, name, created_at FROM example")
	if err != nil {
		log.Fatalf("Ошибка выборки: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        uint64
			name      string
			createdAt time.Time
		)
		if err := rows.Scan(&id, &name, &createdAt); err != nil {
			log.Fatalf("Ошибка сканирования строки: %v", err)
		}
		fmt.Printf("ID=%d Name=%s CreatedAt=%s\n", id, name, createdAt)
	}
}
