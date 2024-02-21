package service

import (
    "context"
    "net/http"
    "order-processing-service/pkg/kafka"
    "order-processing-service/pkg/model"
    opb "order-processing-service/pkg/protobuf/order"
    "order-processing-service/pkg/repository"
    "time"
)

type OrderService struct {
    OrderRepository   repository.OrderRepository
    ProductRepository repository.ProductRepository
}

func (os *OrderService) Save(ctx context.Context, req *opb.OrderSaveRequest) (*opb.OrderSaveResponse, error) {
    order := model.Order{
        CompanyID: req.Company.Id,
        UserID:    req.User.Id,
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

    var p []*opb.OrderSaveResponse_Data_Product
    for _, orderProduct := range op {
        p = append(p, &opb.OrderSaveResponse_Data_Product{
            Name:     orderProduct.Product.Name,
            Price:    orderProduct.Product.Price,
            Quantity: orderProduct.Quantity,
        })
    }

    err = notifyUser(order, op, req.User, req.Company)
    if err != nil {
        return nil, err
    }

    return &opb.OrderSaveResponse{
        Status: http.StatusCreated,
        Data: &opb.OrderSaveResponse_Data{
            Id:        order.ID,
            UserId:    order.UserID,
            CompanyId: order.CompanyID,
            Products:  p,
        },
    }, nil
}

func notifyUser(o model.Order, op []*model.OrderProduct, u *opb.OrderSaveRequest_User, c *opb.OrderSaveRequest_Company) error {
    producer := kafka.NewKafkaProducer()

    var products []kafka.Product
    for _, p := range op {
        products = append(products, kafka.Product{
            Name:     p.Product.Name,
            Price:    p.Product.Price,
            Quantity: p.Quantity,
        })
    }

    order := kafka.OrderConfirmation{
        Id:       o.ID,
        Products: products,
        User: kafka.User{
            Name:  u.Name,
            Email: u.Email,
            Address: kafka.Address{
                City:   u.Address.City,
                State:  u.Address.State,
                Street: u.Address.Street,
                Number: u.Address.Number,
            },
        },
        Company: kafka.Company{
            Name:  c.Name,
            Email: c.Email,
            Address: kafka.Address{
                City:   c.Address.City,
                State:  c.Address.State,
                Street: c.Address.Street,
                Number: c.Address.Number,
            },
        },
    }
    err := producer.SendOrderNotification(order)

    if err != nil {
        return err
    }

    return nil
}

func (os *OrderService) ListByUserId(ctx context.Context, req *opb.ListUserOrdersRequest) (*opb.ListUserOrdersResponse, error) {
    ordersProducts, err := os.OrderRepository.ListByUserId(req.UserId)
    if err != nil {
        return nil, err
    }

    var orders []*opb.ListUserOrdersResponse_Order
    for _, op := range ordersProducts {
        orders = append(orders, &opb.ListUserOrdersResponse_Order{
            Id:       op.OrderId,
            Date:     op.Order.Date.Format("02/01/2006"),
            Products: getOrderProducts(op.OrderId, ordersProducts),
        })
    }

    return &opb.ListUserOrdersResponse{Orders: orders}, nil
}

func getOrderProducts(orderId string, orderProducts []model.OrderProduct) []*opb.ListUserOrdersResponse_Product {
    var products []*opb.ListUserOrdersResponse_Product
    for _, op := range orderProducts {
        if op.OrderId == orderId {
            products = append(products, &opb.ListUserOrdersResponse_Product{
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
