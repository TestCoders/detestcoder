package com.training;

import java.util.*;

public class InMemoryCustomerRepository implements CustomerRepository {
    private final Map<Integer, Customer> customers = new HashMap<>();

    public Optional<Customer> getCustomer(int id) {
        return Optional.ofNullable(this.customers.get(id));
    }

    public List<Customer> getCustomers() {
        return new ArrayList<>(customers.values());
    }

    public void save(Customer customer) throws CustomerAlreadyExistsException {
        if (customers.containsKey(customer.getId())) {
            throw new CustomerAlreadyExistsException();
        }

        this.customers.put(customer.getId(), customer);
    }
}
