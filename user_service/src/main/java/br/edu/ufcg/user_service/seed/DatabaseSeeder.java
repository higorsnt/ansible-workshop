package br.edu.ufcg.user_service.seed;

import br.edu.ufcg.user_service.models.Address;
import br.edu.ufcg.user_service.models.Company;
import br.edu.ufcg.user_service.models.User;
import br.edu.ufcg.user_service.repository.CompanyRepository;
import br.edu.ufcg.user_service.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;

@Component
public class DatabaseSeeder implements CommandLineRunner {

    private final UserRepository userRepository;

    private final CompanyRepository companyRepository;

    @Autowired
    public DatabaseSeeder(UserRepository userRepository, CompanyRepository companyRepository) {
        this.userRepository = userRepository;
        this.companyRepository = companyRepository;
    }

    @Override
    public void run(String... args) throws Exception {
        loadCompanyData();
        loadUserData();
    }

    private void loadCompanyData() {
        if (companyRepository.count() > 0) return;

        List<Company> companyList = new ArrayList<>() {
            {
                add(new Company("Acme Corporation", "contact@acme.com", new Address("123 Main Street", "New York", "NY", "10001")));
                add(new Company("Tech Solutions Inc.", "info@techsolutions.com", new Address("456 Elm Street", "Los Angeles", "CA", "90001")));
                add(new Company("Globex Corporation", "support@globex.com", new Address("789 Oak Avenue", "Chicago", "IL", "60601")));
                add(new Company("Widget World", "info@widgetworld.com", new Address("1011 Pine Road", "Houston", "TX", "77001")));
                add(new Company("Innovative Industries", "info@innovative.com", new Address("1213 Maple Lane", "Miami", "FL", "33101")));
            }
        };

        companyRepository.saveAll(companyList);
    }

    private void loadUserData() {
        if (userRepository.count() > 0) return;

        List<User> userList = new ArrayList<>() {
            {
                new User("John Doe", "john.doe@example.com", "12345678901", new Address("1415 Cedar Boulevard", "San Francisco", "CA", "94101"));
                new User("Jane Smith", "jane.smith@example.com", "23456789012", new Address("1617 Willow Drive", "Seattle", "WA", "98101"));
                new User("Michael Johnson", "michael.johnson@example.com", "34567890123", new Address("1819 Birch Lane", "Boston", "MA", "02101"));
                new User("Emily Brown", "emily.brown@example.com", "45678901234", new Address("2021 Pinecrest Avenue", "Atlanta", "GA", "30301"));
                new User("David Wilson", "david.wilson@example.com", "56789012345", new Address("2223 Oakwood Street", "Denver", "CO", "80201"));
                new User("Jennifer Lee", "jennifer.lee@example.com", "67890123456", new Address("2425 Elmwood Drive", "Philadelphia", "PA", "19101"));
                new User("Christopher Martin", "christopher.martin@example.com", "78901234567", new Address("2627 Maplewood Lane", "Phoenix", "AZ", "85001"));
                new User("Jessica Taylor", "jessica.taylor@example.com", "89012345678", new Address("2829 Willow Avenue", "Dallas", "TX", "75201"));
                new User("Daniel Anderson", "daniel.anderson@example.com", "90123456789", new Address("3031 Birchwood Lane", "Detroit", "MI", "48201"));
                new User("Sarah Garcia", "sarah.garcia@example.com", "01234567890", new Address("3233 Pinecrest Drive", "Las Vegas", "NV", "89101"));
            }
        };

        userRepository.saveAll(userList);
    }
}
