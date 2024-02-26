package br.edu.ufcg.middleware.dto.company;

import br.edu.ufcg.middleware.dto.address.AddressDto;

public record CompanyDto(Long id, String name, String email, AddressDto address) {
}
