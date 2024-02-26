package br.edu.ufcg.middleware.dto.order;

import br.edu.ufcg.middleware.proto.order.ListUserOrdersResponse;
import br.edu.ufcg.middleware.proto.order.OrderSaveResponse;

import java.util.List;

public record OrderDto(String id, String date, List<OrderProductDto> products) {
    public OrderDto(ListUserOrdersResponse.Order o) {
        this(
                o.getId(),
                o.getDate(),
                o.getProductsList().stream().map(OrderProductDto::new).toList()
        );
    }
}
