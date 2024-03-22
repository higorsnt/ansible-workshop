package dotenv

import (
    "os"
    "strconv"
)

func PostgresHost() string {
    return os.Getenv("POSTGRES_HOST")
}

func PostgresPort() int {
    port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
    return port
}

func PostgresUser() string {
    return os.Getenv("POSTGRES_USER")
}

func PostgresPassword() string {
    return os.Getenv("POSTGRES_PASSWORD")
}

func PostgresDatabase() string {
    return os.Getenv("POSTGRES_DATABASE")
}

func ServerPort() string {
    return os.Getenv("SERVER_PORT")
}

func KafkaServer() string {
    return os.Getenv("KAFKA_SERVER")
}
