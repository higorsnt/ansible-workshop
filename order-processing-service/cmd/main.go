package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "google.golang.org/grpc"
    "log"
    "net"
    "order-processing-service/pkg/config"
    "order-processing-service/pkg/dotenv"
    pbc "order-processing-service/pkg/protobuf/cart"
    pbo "order-processing-service/pkg/protobuf/order"
    pbp "order-processing-service/pkg/protobuf/product"
    r "order-processing-service/pkg/repository"
    "order-processing-service/pkg/service"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error loading .env file")
    }

    db := config.DatabaseConnection()

    lis, err := net.Listen("tcp", dotenv.ServerPort())
    if err != nil {
        log.Fatalf("failed to listing: %s\n", err)
    }

    fmt.Printf("Server on. PORT: %s\n", dotenv.ServerPort())

    orderRepository := r.NewOrderRepository(db)
    productRepository := r.NewProductRepository(db)
    cartRepository := r.NewCartRepository(db)

    orderService := service.OrderService{
        ProductRepository: productRepository,
        OrderRepository:   orderRepository,
    }
    productService := service.ProductService{
        ProductRepository: productRepository,
    }
    cartService := service.CartService{
        CartRepository:    cartRepository,
        ProductRepository: productRepository,
    }

    grpcServer := grpc.NewServer()
    pbo.RegisterOrderServiceServer(grpcServer, &orderService)
    pbp.RegisterProductServiceServer(grpcServer, &productService)
    pbc.RegisterCartServiceServer(grpcServer, &cartService)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s\n", err)
    }
}
