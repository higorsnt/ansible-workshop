package br.edu.ufcg.user_service.controller;

import br.edu.ufcg.user_service.dto.company.CompanyDto;
import br.edu.ufcg.user_service.dto.company.CompanySaveDto;
import br.edu.ufcg.user_service.service.company.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping
public class CompanyController {

    private final CompanyService companyService;

    @Autowired
    public CompanyController(CompanyService companyService) {
        this.companyService = companyService;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CompanyDto save(@RequestBody CompanySaveDto companyDto) {
        return companyService.save(companyDto);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public CompanyDto findById(@PathVariable Long id) {
        return companyService.findById(id);
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<CompanyDto> findAll() {
        return companyService.findAll();
    }
}
