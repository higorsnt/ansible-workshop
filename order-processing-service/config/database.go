package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "order-processing-service/dotenv"
    "order-processing-service/exception"
)

func DatabaseConnection() *gorm.DB {
    sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dotenv.PostgresHost(),
        dotenv.PostgresPort(), dotenv.PostgresUser(), dotenv.PostgresPassword(), dotenv.PostgresDatabase())
    db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

    exception.ErrorPanic(err)

    return db
}
