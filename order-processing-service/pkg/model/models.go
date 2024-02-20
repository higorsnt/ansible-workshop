package model

import (
    "gorm.io/gorm"
    "time"
)

type Product struct {
    gorm.Model
    Name        string  `gorm:"type:string;not null;default:null"`
    Description string  `gorm:"type:string"`
    Price       float64 `gorm:"type:float;not null;default:null"`
    Stock       int64   `gorm:"type:int;not null;default:null"`
}

type Order struct {
    ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid();"`
    UserID    int64     `gorm:"type:int;not null;default:null"`
    CompanyID int64     `gorm:"type:int;not null;default:null"`
    Date      time.Time `gorm:"type:date;not null;default:null"`
}

type OrderProduct struct {
    ProductId uint    `gorm:"primaryKey"`
    OrderId   string  `gorm:"primaryKey"`
    Product   Product `gorm:"foreignKey:ProductId"`
    Order     Order   `gorm:"foreignKey:OrderId"`
    Quantity  int64   `gorm:"type:int;not null;default:null"`
}

type Cart struct {
    UserId    int64   `gorm:"type:int;primaryKey"`
    ProductID uint    `gorm:"primaryKey"`
    Quantity  int64   `gorm:"type:int;not null;default:null"`
    Product   Product `gorm:"foreignKey:ProductID"`
}
