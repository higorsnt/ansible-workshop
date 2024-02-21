package kafka

import (
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
    "order-processing-service/pkg/dotenv"
)

type Address struct {
    Street string
    City   string
    State  string
    Number int64
}

type User struct {
    Email   string
    Name    string
    Address Address
}

type Product struct {
    Name     string
    Price    float64
    Quantity int64
}

type Company struct {
    Name    string
    Email   string
    Address Address
}

type OrderConfirmation struct {
    Id       string
    User     User
    Products []Product
    Company  Company
}

type Producer interface {
    SendOrderNotification(confirmation OrderConfirmation) error
}

type producerImpl struct {
}

func NewKafkaProducer() Producer {
    return &producerImpl{}
}

func (p producerImpl) SendOrderNotification(confirmation OrderConfirmation) error {
    const topic = "order-confirmation-email"

    conn, err := kafka.DialLeader(context.Background(), "tcp", dotenv.KafkaServer(), topic, 0)
    if err != nil {
        return err
    }

    marshal, err := json.Marshal(confirmation)
    if err != nil {
        return err
    }

    _, err = conn.WriteMessages(kafka.Message{Value: marshal})
    if err != nil {
        return err
    }

    return nil
}
