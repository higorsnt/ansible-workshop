package br.edu.ufcg.middleware.dto.order;

import br.edu.ufcg.middleware.dto.product.ProductOrderSaveDto;

import java.util.List;

public record OrderSaveDto(Long userId, Long companyId, List<ProductOrderSaveDto> products) {
}
