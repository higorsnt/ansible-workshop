package br.edu.ufcg.user_service.service.company;

import br.edu.ufcg.user_service.dto.company.CompanyDto;
import br.edu.ufcg.user_service.dto.company.CompanySaveDto;
import br.edu.ufcg.user_service.models.Company;

import java.util.List;

public interface CompanyService {

    CompanyDto findById(Long id);

    CompanyDto save(CompanySaveDto companyDto);

    List<CompanyDto> findAll();

}
