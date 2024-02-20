package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "order-processing-service/pkg/dotenv"
    "order-processing-service/pkg/model"
)

func DatabaseConnection() *gorm.DB {
    sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dotenv.PostgresHost(),
        dotenv.PostgresPort(), dotenv.PostgresUser(), dotenv.PostgresPassword(), dotenv.PostgresDatabase())

    db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
    if err != nil {
        log.Fatalf("error on opening database connection: %s", err)
    }

    err = db.AutoMigrate(&model.Order{})
    if err != nil {
        log.Fatalf("error on migrate order model: %s", err)
    }

    err = db.AutoMigrate(&model.Product{})
    if err != nil {
        log.Fatalf("error on migrate product model: %s", err)
    }

    err = db.AutoMigrate(&model.OrderProduct{})
    if err != nil {
        log.Fatalf("error on migrate order product model: %s", err)
    }

    err = db.AutoMigrate(&model.Cart{})
    if err != nil {
        log.Fatalf("error on migrate cart model: %s", err)
    }

    return db
}
