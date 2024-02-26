package br.edu.ufcg.middleware.controller;

import br.edu.ufcg.middleware.dto.company.CompanyDto;
import br.edu.ufcg.middleware.dto.order.OrderDto;
import br.edu.ufcg.middleware.dto.order.OrderSaveDto;
import br.edu.ufcg.middleware.dto.order.OrderSaveResponseDto;
import br.edu.ufcg.middleware.dto.product.ProductOrderSaveDto;
import br.edu.ufcg.middleware.dto.user.UserDto;
import br.edu.ufcg.middleware.proto.order.ListUserOrdersRequest;
import br.edu.ufcg.middleware.proto.order.ListUserOrdersResponse;
import br.edu.ufcg.middleware.proto.order.OrderSaveRequest;
import br.edu.ufcg.middleware.proto.order.OrderSaveResponse;
import br.edu.ufcg.middleware.proto.order.OrderServiceGrpc.OrderServiceBlockingStub;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/order")
public class OrderController {

    @GrpcClient("order-service")
    private OrderServiceBlockingStub orderServiceStub;

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public OrderSaveResponseDto create(@RequestBody OrderSaveDto orderSaveDto) {
        OrderSaveRequest.Builder request = OrderSaveRequest.newBuilder()
                .setUser(makeUserRequest(orderSaveDto.user()))
                .setCompany(makeCompanyRequest(orderSaveDto.company()));

        for (int i = 0; i < orderSaveDto.products().size(); i++) {
            ProductOrderSaveDto product = orderSaveDto.products().get(i);
            request.setProducts(i, OrderSaveRequest.Product.newBuilder()
                    .setId(product.id())
                    .setQuantity(product.quantity())
                    .build());
        }

        OrderSaveResponse.Data data = this.orderServiceStub.save(request.build()).getData();
        return new OrderSaveResponseDto(data);
    }

    private OrderSaveRequest.Company makeCompanyRequest(CompanyDto company) {
        OrderSaveRequest.Address address = OrderSaveRequest.Address.newBuilder()
                .setCity(company.address().city())
                .setNumber(company.address().number())
                .setState(company.address().state())
                .setStreet(company.address().street())
                .build();
        return OrderSaveRequest.Company.newBuilder()
                .setId(company.id())
                .setName(company.name())
                .setEmail(company.email())
                .setAddress(address)
                .build();
    }

    private OrderSaveRequest.User makeUserRequest(UserDto user) {
        OrderSaveRequest.Address address = OrderSaveRequest.Address.newBuilder()
                .setCity(user.address().city())
                .setNumber(user.address().number())
                .setState(user.address().state())
                .setStreet(user.address().street())
                .build();

        return OrderSaveRequest.User.newBuilder()
                .setId(user.id())
                .setName(user.name())
                .setEmail(user.email())
                .setAddress(address)
                .build();
    }

    @GetMapping("/{userId}")
    @ResponseStatus(HttpStatus.OK)
    public List<OrderDto> listByUserId(@PathVariable Integer userId) {
        ListUserOrdersRequest request = ListUserOrdersRequest.newBuilder()
                .setUserId(userId)
                .build();
        List<ListUserOrdersResponse.Order> order = this.orderServiceStub.listByUserId(request).getOrdersList();

        return order.stream().map(OrderDto::new).toList();
    }

}
