package br.edu.ufcg.user_service.service.user;

import br.edu.ufcg.user_service.dto.user.UserDto;
import br.edu.ufcg.user_service.dto.user.UserSaveDto;

import java.util.List;

public interface UserService {

    UserDto findById(Long id);

    UserDto save(UserSaveDto userSaveDto);

    List<UserDto> findAll();

}
