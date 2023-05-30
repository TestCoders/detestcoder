package com.training;

import java.util.List;
import java.util.Optional;

public interface CustomerRepository {

    Optional<Customer> getCustomer(int id);
    List<Customer> getCustomers();
    void save(Customer customer) throws CustomerAlreadyExistsException;
}
