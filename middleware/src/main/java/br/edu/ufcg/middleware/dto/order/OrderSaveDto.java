package br.edu.ufcg.middleware.dto.order;

import br.edu.ufcg.middleware.dto.company.CompanyDto;
import br.edu.ufcg.middleware.dto.product.ProductOrderSaveDto;
import br.edu.ufcg.middleware.dto.user.UserDto;

import java.util.List;

public record OrderSaveDto(UserDto user, CompanyDto company, List<ProductOrderSaveDto> products) {
}
