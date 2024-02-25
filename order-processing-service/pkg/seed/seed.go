package seed

import (
    "gorm.io/gorm"
    "order-processing-service/pkg/model"
)

var products = []model.Product{
    {Name: "Notebook", Description: "Notebook de alta performance com armazenamento SSD", Price: 1200.00, Stock: 50},
    {Name: "Smartphone", Description: "Modelo mais recente com câmera de alta resolução", Price: 800.00, Stock: 100},
    {Name: "Tablet", Description: "Tablet fino e leve para uso em movimento", Price: 400.00, Stock: 75},
    {Name: "Smartwatch", Description: "Smartwatch para rastreamento de fitness com monitor de frequência cardíaca", Price: 250.00, Stock: 30},
    {Name: "Caixa de Som Bluetooth", Description: "Caixa de som portátil com longa duração de bateria", Price: 100.00, Stock: 80},
    {Name: "Fones de Ouvido Sem Fio", Description: "Fones de ouvido com cancelamento de ruído e conectividade Bluetooth", Price: 150.00, Stock: 60},
    {Name: "Console de Jogos", Description: "Console de jogos poderoso com suporte a gráficos 4K", Price: 500.00, Stock: 20},
    {Name: "Câmera Digital", Description: "Câmera DSLR profissional com lentes intercambiáveis", Price: 1200.00, Stock: 15},
    {Name: "HD Externo", Description: "Solução de armazenamento de alta capacidade para backups e transferência de dados", Price: 150.00, Stock: 50},
    {Name: "Monitor de Atividade Física", Description: "Dispositivo vestível para monitoramento de níveis de atividade e qualidade do sono", Price: 80.00, Stock: 100},
    {Name: "Cafeteira", Description: "Cafeteira automática com configurações programáveis", Price: 50.00, Stock: 40},
    {Name: "Liquidificador", Description: "Liquidificador de alta velocidade para fazer smoothies e sopas", Price: 70.00, Stock: 30},
    {Name: "Torradeira", Description: "Torradeira de duas fatias com configurações de douramento ajustáveis", Price: 30.00, Stock: 50},
    {Name: "Chaleira Elétrica", Description: "Chaleira elétrica de fervura rápida para chá e café", Price: 40.00, Stock: 60},
    {Name: "Processador de Alimentos", Description: "Aparelho de cozinha versátil para picar, fatiar e triturar", Price: 90.00, Stock: 25},
    {Name: "Panela de Arroz", Description: "Panela de arroz automática com panela interna antiaderente", Price: 60.00, Stock: 35},
    {Name: "Ferro a Vapor", Description: "Ferro a vapor potente para roupas sem rugas", Price: 40.00, Stock: 70},
    {Name: "Aspirador de Pó", Description: "Aspirador de pó sem saco com filtro HEPA", Price: 200.00, Stock: 45},
    {Name: "Secador de Cabelo", Description: "Secador de cabelo profissional com múltiplas configurações de calor e velocidade", Price: 80.00, Stock: 55},
    {Name: "Escova de Dentes Elétrica", Description: "Escova de dentes elétrica recarregável com temporizador", Price: 50.00, Stock: 85},
}

func Load(db *gorm.DB) {
    db.Create(&products)
}
