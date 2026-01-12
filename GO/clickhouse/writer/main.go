package main

import (
	"context"
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

	// Можно переносить создание таблицы отдельно, если не нужно каждый раз
	if err := conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS example (
            id UInt64,
            name String,
            created_at DateTime
        ) Engine = Memory
    `); err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example")
	if err != nil {
		log.Fatalf("Ошибка подготовки батча: %v", err)
	}
	defer batch.Close()

	err = batch.Append(uint64(1), "writer test", time.Now())
	if err != nil {
		log.Fatalf("Ошибка добавления данных: %v", err)
	}

	if err := batch.Send(); err != nil {
		log.Fatalf("Ошибка отправки батча: %v", err)
	}

	log.Println("Данные успешно записаны в ClickHouse!")
}
