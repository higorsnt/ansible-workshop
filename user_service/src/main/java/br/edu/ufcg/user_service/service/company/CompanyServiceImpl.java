package br.edu.ufcg.user_service.service.company;

import br.edu.ufcg.user_service.dto.company.CompanyDto;
import br.edu.ufcg.user_service.dto.company.CompanySaveDto;
import br.edu.ufcg.user_service.models.Company;
import br.edu.ufcg.user_service.repository.CompanyRepository;
import jakarta.persistence.EntityNotFoundException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
class CompanyServiceImpl implements CompanyService {

    private final CompanyRepository companyRepository;

    @Autowired
    public CompanyServiceImpl(CompanyRepository companyRepository) {
        this.companyRepository = companyRepository;
    }

    @Override
    public CompanyDto findById(Long id) {
        return new CompanyDto(companyRepository.findById(id)
                .orElseThrow(EntityNotFoundException::new));
    }

    @Override
    public CompanyDto save(CompanySaveDto companyDto) {
        Company company = companyDto.toEntity();
        return new CompanyDto(companyRepository.save(company));
    }

    @Override
    public List<CompanyDto> findAll() {
        return companyRepository.findAll()
                .stream()
                .map(CompanyDto::new)
                .toList();
    }
}
