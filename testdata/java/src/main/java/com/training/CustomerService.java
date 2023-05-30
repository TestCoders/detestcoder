package com.training;

import java.util.Optional;

public class CustomerService {
    private CustomerRepository customerRepository;

    public CustomerService(CustomerRepository customerRepository) {
        this.customerRepository = customerRepository;
    }

    public Optional<Customer> getCustomer(int customerId) {
        return customerRepository.getCustomer(customerId);
    }

    public boolean save(Customer newCustomer) {
        try {
            this.customerRepository.save(newCustomer);

            return true;
        } catch (CustomerAlreadyExistsException e) {
            return false;
        }
    }
}
