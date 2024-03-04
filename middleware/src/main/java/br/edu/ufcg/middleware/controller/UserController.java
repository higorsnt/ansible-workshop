package br.edu.ufcg.middleware.controller;

import br.edu.ufcg.middleware.dto.user.UserDto;
import br.edu.ufcg.middleware.dto.user.UserSaveDto;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.util.DefaultUriBuilderFactory;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@RestController
@RequestMapping("/user")
public class UserController {

    private static final String RESOURCE = "/user";

    private final RestTemplate restTemplate;

    public UserController(@Value("${user_service.host}") String host) {
        this.restTemplate = new RestTemplate();
        this.restTemplate.setUriTemplateHandler(new DefaultUriBuilderFactory(host));
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto save(@RequestBody UserSaveDto userDto) {
        return this.restTemplate.postForObject(RESOURCE, userDto, UserDto.class);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDto findById(@PathVariable Long id) {
        return this.restTemplate.getForObject(RESOURCE, UserDto.class, Map.of("id", id));
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<UserDto> findAll() {
        return Arrays.stream(
                        Objects.requireNonNull(this.restTemplate.getForObject(RESOURCE, UserDto[].class)))
                .toList();
    }

}
