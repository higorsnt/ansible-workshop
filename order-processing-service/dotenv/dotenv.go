package dotenv

import (
    "order-processing-service/exception"
    "os"
    "strconv"
)

func PostgresHost() string {
    return os.Getenv("POSTGRES_HOST")
}

func PostgresPort() int {
    port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
    if err != nil {
        exception.ErrorPanic(err)
    }

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
