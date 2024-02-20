package service

import (
    "context"
    "net/http"
    "order-processing-service/pkg/model"
    orderpb "order-processing-service/pkg/protobuf/order"
    "order-processing-service/pkg/repository"
    "time"
)

type OrderService struct {
    OrderRepository   repository.OrderRepository
    ProductRepository repository.ProductRepository
}

func (os *OrderService) Save(ctx context.Context, req *orderpb.OrderSaveRequest) (*orderpb.OrderSaveResponse, error) {
    order := model.Order{
        CompanyID: req.CompanyId,
        UserID:    req.UserId,
        Date:      time.Now(),
    }

    var op []*model.OrderProduct

    err := os.OrderRepository.Save(&order)
    if err != nil {
        return nil, err
    }

    for _, product := range req.Products {
        p, err := os.ProductRepository.FindById(product.Id)

        if err != nil {
            return nil, err
        }

        op = append(op, &model.OrderProduct{
            Product:  p,
            Order:    order,
            Quantity: product.Quantity,
        })
    }

    err = os.OrderRepository.SaveOrderProduct(op)
    if err != nil {
        return nil, err
    }

    var p []*orderpb.OrderSaveResponse_Data_Product
    for _, orderProduct := range op {
        p = append(p, &orderpb.OrderSaveResponse_Data_Product{
            Name:     orderProduct.Product.Name,
            Price:    orderProduct.Product.Price,
            Quantity: orderProduct.Quantity,
        })
    }

    return &orderpb.OrderSaveResponse{
        Status: http.StatusCreated,
        Data: &orderpb.OrderSaveResponse_Data{
            Id:        order.ID,
            UserId:    order.UserID,
            CompanyId: order.CompanyID,
            Products:  p,
        },
    }, nil
}

func (os *OrderService) ListByUserId(ctx context.Context, req *orderpb.ListUserOrdersRequest) (*orderpb.ListUserOrdersResponse, error) {
    ordersProducts, err := os.OrderRepository.ListByUserId(req.UserId)
    if err != nil {
        return nil, err
    }

    var orders []*orderpb.ListUserOrdersResponse_Order
    for _, op := range ordersProducts {
        orders = append(orders, &orderpb.ListUserOrdersResponse_Order{
            Id:       op.OrderId,
            Date:     op.Order.Date.Format("02/01/2006"),
            Products: getOrderProducts(op.OrderId, ordersProducts),
        })
    }

    return &orderpb.ListUserOrdersResponse{Orders: orders}, nil
}

func getOrderProducts(orderId string, orderProducts []model.OrderProduct) []*orderpb.ListUserOrdersResponse_Product {
    var products []*orderpb.ListUserOrdersResponse_Product
    for _, op := range orderProducts {
        if op.OrderId == orderId {
            products = append(products, &orderpb.ListUserOrdersResponse_Product{
                Id:          int32(op.ProductId),
                Name:        op.Product.Name,
                Description: op.Product.Description,
                Price:       op.Product.Price,
                Quantity:    op.Quantity,
            })
        }
    }

    return products
}
