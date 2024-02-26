package br.edu.ufcg.user_service.service.user;

import br.edu.ufcg.user_service.dto.user.UserDto;
import br.edu.ufcg.user_service.dto.user.UserSaveDto;
import br.edu.ufcg.user_service.models.User;
import br.edu.ufcg.user_service.repository.UserRepository;
import jakarta.persistence.EntityNotFoundException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
class UserServiceImpl implements UserService {

    private final UserRepository userRepository;

    @Autowired
    public UserServiceImpl(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Override
    public UserDto findById(Long id) {
        return new UserDto(userRepository.findById(id)
                .orElseThrow(EntityNotFoundException::new));
    }

    @Override
    public UserDto save(UserSaveDto userSaveDto) {
        User user = userSaveDto.toEntity();
        return new UserDto(userRepository.save(user));
    }

    @Override
    public List<UserDto> findAll() {
        return userRepository.findAll()
                .stream()
                .map(UserDto::new)
                .toList();
    }
}
