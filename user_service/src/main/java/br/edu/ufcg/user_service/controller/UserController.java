package br.edu.ufcg.user_service.controller;

import br.edu.ufcg.user_service.dto.user.UserDto;
import br.edu.ufcg.user_service.dto.user.UserSaveDto;
import br.edu.ufcg.user_service.service.user.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/user")
public class UserController {

    private final UserService userService;

    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto save(@RequestBody UserSaveDto userDto) {
        return userService.save(userDto);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto findById(@PathVariable Long id) {
        return userService.findById(id);
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<UserDto> findAll() {
        return userService.findAll();
    }
}
