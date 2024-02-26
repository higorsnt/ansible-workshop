package br.edu.ufcg.user_service.dto.user;

import br.edu.ufcg.user_service.dto.address.AddressDto;
import br.edu.ufcg.user_service.models.User;

public record UserSaveDto(String name, String email, String document, AddressDto address) {
    public User toEntity() {
        return new User(name, email, document, address.toEntity());
    }
}
