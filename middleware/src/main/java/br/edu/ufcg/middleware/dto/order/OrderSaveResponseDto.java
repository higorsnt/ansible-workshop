package br.edu.ufcg.middleware.dto.order;

import br.edu.ufcg.middleware.dto.product.ProductOrderResponseDto;
import br.edu.ufcg.middleware.proto.order.OrderSaveResponse;

import java.util.List;

public record OrderSaveResponseDto(String id, Long userId, Long companyId, List<ProductOrderResponseDto> products) {
    public OrderSaveResponseDto(OrderSaveResponse.Data data) {
        this(
                data.getId(),
                data.getUserId(),
                data.getCompanyId(),
                data.getProductsList().stream().map(ProductOrderResponseDto::new).toList()
        );
    }
}
