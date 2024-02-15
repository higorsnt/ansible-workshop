package order_processing_service

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "time"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    router := gin.Default()
    //routes := router.NewRouter(tagsController)

    router.GET("", func(context *gin.Context) {
        context.JSON(http.StatusOK, "welcome home")
    })

    server := &http.Server{
        Addr:           ":8888",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    err = server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
