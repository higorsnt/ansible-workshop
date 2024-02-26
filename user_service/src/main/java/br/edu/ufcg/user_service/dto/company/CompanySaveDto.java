package br.edu.ufcg.user_service.dto.company;

import br.edu.ufcg.user_service.dto.address.AddressDto;
import br.edu.ufcg.user_service.models.Company;

public record CompanySaveDto(String name, String email, AddressDto addressDto) {

    public Company toEntity() {
        return new Company(name, email, addressDto.toEntity());
    }

}
