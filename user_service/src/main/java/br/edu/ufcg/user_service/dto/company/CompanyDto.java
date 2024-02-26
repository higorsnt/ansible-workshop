package br.edu.ufcg.user_service.dto.company;

import br.edu.ufcg.user_service.dto.address.AddressDto;
import br.edu.ufcg.user_service.models.Company;

public record CompanyDto(Long id, String name, String email, AddressDto addressDto) {
    public CompanyDto(Company company) {
        this(
                company.getId(),
                company.getName(),
                company.getEmail(),
                new AddressDto(company.getAddress())
        );
    }
}
