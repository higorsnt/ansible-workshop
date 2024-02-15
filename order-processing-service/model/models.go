package model

type Address struct {
    Id     int
    Street string
    City   string
    State  string
    Number int
}

type User struct {
    Id        int
    Name      string
    Email     string
    AddressId int
    Address   Address
}

type Company struct {
    Id        int
    Name      string
    Email     string
    AddressId int
    Address   Address
}

type Product struct {
    Id       int
    Name     string
    Price    float32
    Quantity int
}

type Order struct {
    Id       string `gorm:"type:string,primary_key"`
    User     User
    Products Product
    Company  Company
}
