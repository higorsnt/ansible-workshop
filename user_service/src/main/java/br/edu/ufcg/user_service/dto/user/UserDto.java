package br.edu.ufcg.user_service.dto.user;

import br.edu.ufcg.user_service.dto.address.AddressDto;
import br.edu.ufcg.user_service.models.User;

public record UserDto(Long id, String name, String email, AddressDto address) {
    public UserDto(User user) {
        this(
                user.getId(),
                user.getName(),
                user.getEmail(),
                new AddressDto(user.getAddress())
        );
    }
}
