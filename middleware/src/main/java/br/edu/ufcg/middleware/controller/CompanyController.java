package br.edu.ufcg.middleware.controller;

import br.edu.ufcg.middleware.dto.company.CompanyDto;
import br.edu.ufcg.middleware.dto.company.CompanySaveDto;
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
@RequestMapping("/company")
public class CompanyController {

    private static final String RESOURCE = "/company";

    private final RestTemplate restTemplate;

    public CompanyController() {
        String host = System.getenv("USER_SERVICE_HOST");
        this.restTemplate = new RestTemplate();
        this.restTemplate.setUriTemplateHandler(new DefaultUriBuilderFactory(host));
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CompanyDto save(@RequestBody CompanySaveDto companyDto) {
        return this.restTemplate
                .postForObject(RESOURCE, companyDto, CompanyDto.class);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public CompanyDto findById(@PathVariable Long id) {
        return this.restTemplate.getForObject(RESOURCE, CompanyDto.class, Map.of("id", id));
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<CompanyDto> findAll() {
        return Arrays.stream(
                Objects.requireNonNull(
                        this.restTemplate.getForObject(RESOURCE, CompanyDto[].class)))
                .toList();
    }

}
